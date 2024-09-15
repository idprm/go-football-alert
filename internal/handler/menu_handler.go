package handler

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/sirupsen/logrus"
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

	if req.IsMain() {
		main, err := h.menuService.GetAll()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		var mainData []string
		for i, s := range main {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			mainData = append(mainData, row)
		}

		l.WithFields(logrus.Fields{"request": req}).Info("USSD")
		dataJson, _ := json.Marshal(req)
		h.rmq.IntegratePublish(
			RMQ_USSD_EXCHANGE,
			RMQ_USSD_QUEUE,
			RMQ_DATA_TYPE,
			"", string(dataJson),
		)

		justString := strings.Join(mainData, "\n")
		return c.Status(fiber.StatusOK).SendString(justString)

	} else {

		if !h.menuService.IsKeyPress(req.GetText()) {
			return c.Status(fiber.StatusNotFound).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusNotFound,
					Message:    "menu_not_found",
				},
			)
		}
		menu, err := h.menuService.GetMenuByKeyPress(req.GetText())
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		submenus, err := h.menuService.GetMenuByParentId(menu.GetId())
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		var subData []string
		subData = append(subData, menu.GetName())
		for i, s := range submenus {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			subData = append(subData, row)
		}
		if req.IsLevel() {
			/**
			 **	for loop data
			 **/

			// row := "0. Suiv"
			// subData = append(subData, row)
			subData = h.convertToArrayString(req, subData)
			// log.Println(subData)
			// subData = append(subData, row)
		} else {
			prevs := &[]entity.Menu{
				{ID: 0, Name: "Préc", KeyPress: "0"},
				{ID: 98, Name: "Accueil", KeyPress: "98"},
			}
			for _, p := range *prevs {
				row := p.KeyPress + ". " + p.Name
				subData = append(subData, row)
			}
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
		return c.Status(fiber.StatusOK).SendString(justString)
	}
}

func (h *IncomingHandler) convertToArrayString(req *model.UssdRequest, subData []string) []string {
	var menus []*entity.Menu

	switch req.GetText() {
	case "1*1":
		menus = h.LiveMatch()
	case "1*2":
		menus = h.Schedule()
	default:
		menus = []*entity.Menu{}
	}

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
