package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/sirupsen/logrus"
)

const (
	ACT_MENU    string = "MENU"
	ACT_NO_MENU string = "NO_MENU"
	ACT_CONFIRM string = "CONFIRM"
	ACT_REG     string = "REG"
)

var (
	USSD_TITLE    string = "Orange Football Club, votre choix:"
	USSD_MENU_404 string = "Menu not found"
)

func (h *IncomingHandler) Callback(c *fiber.Ctx) error {
	l := h.logger.Init("ussd", true)
	req := new(model.UssdRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	// valid service
	if h.IsService(req.GetServiceCode()) {
		// get cache
		cache, err := h.ussdService.Get(req.GetMsisdn())
		if err != nil {
			log.Println(err.Error())
		}

		// get service
		service, err := h.serviceService.Get(req.GetServiceCode())
		if err != nil {
			log.Println(err.Error())
		}

		// main menu
		if req.IsMain() {

			// init array string
			var mainData []string

			// menu level 1
			main, err := h.menuService.GetAll()
			if err != nil {
				mainData = append(mainData, err.Error())
			}

			// title ussd
			mainData = append(mainData, USSD_TITLE)
			// loop main menu
			for i, s := range main {
				idx := strconv.Itoa(i + 1)
				row := idx + ". " + s.Name
				mainData = append(mainData, row)
			}
			mainData = append(mainData, "0. Suiv")

			l.WithFields(logrus.Fields{"request": req}).Info("USSD")

			jsonData, _ := json.Marshal(req)
			h.rmq.IntegratePublish(
				RMQ_USSD_EXCHANGE,
				RMQ_USSD_QUEUE,
				RMQ_DATA_TYPE,
				"", string(jsonData),
			)
			justString := strings.Join(mainData, "\n")
			// set cache
			h.ussdService.Set(
				&entity.Ussd{
					Msisdn:    req.GetMsisdn(),
					KeyPress:  req.GetText(),
					Action:    ACT_MENU,
					Result:    justString,
					CreatedAt: time.Now(),
				},
			)
			return c.Status(fiber.StatusOK).SendString(justString)
		} else {
			// init cache
			ca := &entity.Ussd{
				Msisdn:    req.GetMsisdn(),
				KeyPress:  req.GetText(),
				CreatedAt: time.Now(),
			}

			// if keyPress not found in database
			if !h.menuService.IsKeyPress(req.GetText()) {
				var subData []string

				// is valid pressKey level 3
				if req.IsFilterLevel3() {

					// menu level 3
					if req.IsLiveMatch() {
						prevs := &[]entity.Menu{
							{ID: 0, Name: "Yes Yes", KeyPress: "0"},
							{ID: 98, Name: "Accueil", KeyPress: "98"},
						}
						ca.SetAction(ACT_MENU)
						for _, p := range *prevs {
							row := p.KeyPress + ". " + p.Name
							subData = append(subData, row)
						}
					} else {
						prevs := &[]entity.Menu{
							{ID: 0, Name: "Préc", KeyPress: "0"},
							{ID: 98, Name: "Accueil", KeyPress: "98"},
						}
						ca.SetAction(ACT_MENU)
						for _, p := range *prevs {
							row := p.KeyPress + ". " + p.Name
							subData = append(subData, row)
						}
					}
				} else {
					// menu not found
					notFound := &[]entity.Menu{
						{ID: 0, Name: USSD_MENU_404, KeyPress: "0"},
						{ID: 98, Name: "Accueil", KeyPress: "98"},
					}
					ca.SetAction(ACT_NO_MENU)
					for _, p := range *notFound {
						var row string
						if strings.Contains(p.KeyPress, "0") {
							row = p.Name
						} else {
							row = p.KeyPress + ". " + p.Name
						}
						subData = append(subData, row)
					}
				}

				justString := strings.Join(subData, "\n")
				// set cache
				ca.SetResult(justString)
				h.ussdService.Set(ca)
				return c.Status(fiber.StatusOK).SendString(justString)
			}

			// init array string
			var subData []string

			if cache != nil {
				if h.menuService.IsKeyPress(cache.GetKeyPress()) {
					menu, err := h.menuService.GetMenuByKeyPress(cache.GetKeyPress())
					if err != nil {
						subData = append(subData, err.Error())
					}

					submenus, err := h.menuService.GetMenuByParentId(menu.GetId())
					if err != nil {
						subData = append(subData, err.Error())
					}

					// is valid pressKey level 1
					if req.IsFilterLevel1() {
						if !h.IsActiveSub(req) {
							subData = h.confirmSub(service, subData)
							ca.SetAction(ACT_CONFIRM)
						}
					} else {
						ca.SetAction(ACT_MENU)
						subData = append(subData, menu.GetName())
						for i, s := range submenus {
							idx := strconv.Itoa(i + 1)
							row := idx + ". " + s.Name
							subData = append(subData, row)
						}
					}
				}

			}

			if req.IsLevel() {
				subData = h.convertToArrayString(req, subData)
				subData = append(subData, "0. Suiv")
				ca.SetAction(ACT_MENU)
			}

			l.WithFields(logrus.Fields{"request": req}).Info("USSD")
			jsonData, _ := json.Marshal(req)
			h.rmq.IntegratePublish(
				RMQ_USSD_EXCHANGE,
				RMQ_USSD_QUEUE,
				RMQ_DATA_TYPE,
				"", string(jsonData),
			)
			justString := strings.Join(subData, "\n")
			// set cache
			ca.SetResult(justString)
			h.ussdService.Set(ca)
			return c.Status(fiber.StatusOK).SendString(justString)
		}
	}

	var noServiceData []string
	noServiceData = append(noServiceData, "No service")
	justString := strings.Join(noServiceData, "\n")
	return c.Status(fiber.StatusOK).SendString(justString)
}

func (h *IncomingHandler) convertToArrayString(req *model.UssdRequest, subData []string) []string {
	var menus []*entity.Menu
	switch req.GetText() {
	case "1*1":
		menus = h.LiveMatch()
	case "1*2":
		menus = h.Schedule()
	case "1*3":
		menus = h.Lineup()
	case "1*4":
		menus = h.MatchStats()
	case "1*5":
		menus = h.DisplayLiveMatch()
	case "2":
		menus = h.FlashNews()
	case "3":
		menus = h.CreditGoal()
	case "4*1":
		menus = h.ChampResults()
	case "4*2":
		menus = h.ChampStandings()
	case "4*3":
		menus = h.ChampSchedule()
	case "4*4":
		menus = h.ChampTeam()
	case "4*5":
		menus = h.ChampCreditScore()
	case "4*6":
		menus = h.ChampCreditGoal()
	case "4*7":
		menus = h.ChampSMSAlerte()
	case "4*8":
		menus = h.ChampSMSAlerteEquipe()
	case "5":
		menus = h.Prediction()
	case "6*1":
		menus = h.KitFoot()
	case "6*2":
		menus = h.Europe()
	case "6*3":
		menus = h.Afrique()
	case "6*4":
		menus = h.SMSAlerteEquipe()
	case "6*5":
		menus = h.FootInternational()
	case "7*1":
		menus = h.AlerteChampMaliEquipe()
	case "7*2":
		menus = h.AlertePremierLeagueEquipe()
	case "7*3":
		menus = h.AlerteLaLigaEquipe()
	case "7*4":
		menus = h.AlerteLigue1Equipe()
	case "7*5":
		menus = h.AlerteSerieAEquipe()
	case "7*6":
		menus = h.AlerteBundesligueEquipe()
	case "8*1":
		menus = h.ChampionLeague()
	case "8*2":
		menus = h.PremierLeague()
	case "8*3":
		menus = h.LaLiga()
	case "8*4":
		menus = h.Ligue1()
	case "8*5":
		menus = h.LEuropa()
	case "8*6":
		menus = h.SerieA()
	case "8*7":
		menus = h.Bundesligua()
	case "8*8":
		menus = h.ChampPortugal()
	case "8*9":
		menus = h.SaudiLeague()
	default:
		menus = []*entity.Menu{{ID: 0, Name: USSD_MENU_404, KeyPress: "0"}}
	}

	for i, s := range menus {
		idx := strconv.Itoa(i + 1)
		var row string
		if strings.Contains(s.Name, "No") {
			row = s.Name
		} else {
			row = idx + ". " + s.Name
		}
		subData = append(subData, row)
	}
	return subData
}

func (h *IncomingHandler) confirmSub(service *entity.Service, subData []string) []string {
	price := strconv.FormatFloat(service.GetPrice(), 'f', 0, 64)
	menus := []*entity.Menu{
		{ID: 0, Name: "You are about to subscribe to the " + service.GetPackage() + " offer " + price, KeyPress: "0"},
		{ID: 1, Name: "Yes", KeyPress: "1"},
		{ID: 2, Name: "No", KeyPress: "2"},
	}
	for _, s := range menus {
		var row string
		if strings.Contains(s.KeyPress, "0") {
			row = s.Name
		} else {
			row = s.KeyPress + ". " + s.Name
		}
		subData = append(subData, row)
	}
	return subData
}

func (h *IncomingHandler) IsActiveSub(req *model.UssdRequest) bool {
	service, err := h.getService(req.GetServiceCode())
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsActiveSubscription(service.GetId(), req.GetMsisdn())
}

func (h *IncomingHandler) IsService(code string) bool {
	return h.serviceService.IsService(code)
}

func (h *IncomingHandler) getService(code string) (*entity.Service, error) {
	return h.serviceService.Get(code)
}

func (h *IncomingHandler) ChooseMenu(req *model.UssdRequest) {
	// if req.IsLineup() {
	// 	menus =  h.Lineup()
	// }
	// if req.IsMatchStats() {
	// 	menus =  h.MatchStats()
	// }
	// if req.IsDisplayLiveMatch() {
	// 	menus =  h.DsiplayMatchStats()
	// }
	// if req.IsFlashNews() {
	// 	menus =  h.FlashNews()
	// }
	// if req.IsCreditGoal() {
	// 	menus =  h.CreditGoal()
	// }
	// if req.IsChampResults() {
	// 	menus =  h.ChampResults()
	// }
	// if req.IsChampStandings() {
	// 	menus =  h.ChampStandings()
	// }
	// if req.IsChampTeam() {
	// 	menus =  h.ChampTeam()
	// }
	// if req.IsChampCreditScore() {
	// 	menus = h.ChampCreditScore()
	// }
	// if req.IsChampCreditGoal() {
	// 	menus =  h.ChampCreditGoal()
	// }
	// if req.IsChampSMSAlerte() {
	// 	menus = h.ChampSMSAlerte()
	// }

	// if req.IsChampSMSAlerteEquipe() {
	// 	menus = h.ChampSMSAlerteEquipe()

	// }

	// if req.IsPrediction() {
	// 	menus =  h.Prediction()

	// }

	// if req.IsKitFoot() {
	// 	menus =  h.KitFoot()

	// }

	// if req.IsEurope() {
	// 	menus =  h.Europe()

	// }

	// if req.IsAfrique() {
	// 	menus =  h.Afrique()

	// }

	// if req.IsSMSAlerteEquipe() {
	// 	menus = h.SMSAlerteEquipe()

	// }
	// if req.IsFootInternational() {
	// 	menus =  h.FootInternational()

	// }
	// if req.IsAlerteChampMaliEquipe() {
	// 	menus =  h.AlerteChampMaliEquipe()

	// }
	// if req.IsAlertePremierLeagueEquipe() {
	// 	menus = h.AlertePremierLeagueEquipe()

	// }

	// if req.IsAlerteLaLigaEquipe() {
	// 	menus = h.AlerteLaLigaEquipe()

	// if req.IsAlerteLigue1Equipe() {
	// 	menus =  h.AlerteLigue1Equipe()

	// }

	// if req.IsAlerteSerieAEquipe() {
	// 	menus =  h.AlerteSerieAEquipe()

	// }

	// if req.IsAlerteBundesligueEquipe() {
	// 	menus =  h.AlerteBundesligueEquipe()

	// }

	// if req.IsChampionLeague() {
	// 	menus = h.ChampionLeague()

	// }

	// if req.IsPremierLeague() {
	// 	menus =  h.PremierLeague()

	// }

	// if req.IsLaLiga() {
	// 	menus =  h.LaLiga()

	// }

	// if req.IsLigue1() {
	// 	menus =  h.Ligue1()

	// }

	// if req.IsSerieA() {
	// 	menus =  h.SerieA()

	// }

	// if req.IsBundesligua() {
	// 	menus = h.Bundesligua()

	// }

	// if req.IsChampPortugal() {
	// 	menus = h.ChampPortugal()

	// }
	// if req.IsSaudiLeague() {
	// 	menus =  h.SaudiLeague()
	// }
}

func (h *IncomingHandler) LiveMatch() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Live Match", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Schedule() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Schedule", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Lineup() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Lineup", KeyPress: "0"},
	}
}

func (h *IncomingHandler) MatchStats() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Match Stats", KeyPress: "0"},
	}
}

func (h *IncomingHandler) DisplayLiveMatch() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Display Live Match", KeyPress: "0"},
	}
}

func (h *IncomingHandler) FlashNews() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Flash News", KeyPress: "0"},
	}
}

func (h *IncomingHandler) CreditGoal() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Credit Goal", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampResults() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Results", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampStandings() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Standings", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSchedule() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Schedule", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampTeam() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Team", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampCreditScore() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Credit Score", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampCreditGoal() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Credit Goal", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSMSAlerte() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ SMS Alerte", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSMSAlerteEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ SMS Alerte Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Prediction() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Prediction", KeyPress: "0"},
	}
}

func (h *IncomingHandler) KitFoot() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No KitFoot", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Europe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Europe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Afrique() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Afrique", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SMSAlerteEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No SMS Alerte Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) FootInternational() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Foot International", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteChampMaliEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Champ Mali Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlertePremierLeagueEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Premier League Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteLaLigaEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte LaLiga Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteLigue1Equipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Ligue 1 Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteSerieAEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Serie A Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteBundesligueEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Bundesligue Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampionLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champion League", KeyPress: "0"},
	}
}

func (h *IncomingHandler) PremierLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Premier League", KeyPress: "0"},
	}
}

func (h *IncomingHandler) LaLiga() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No La Liga", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Ligue1() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Ligue 1", KeyPress: "0"},
	}
}

func (h *IncomingHandler) LEuropa() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No L Europa", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SerieA() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Serie A", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Bundesligua() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Bundesligua", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampPortugal() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Portugal", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SaudiLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Saudi League", KeyPress: "0"},
	}
}
