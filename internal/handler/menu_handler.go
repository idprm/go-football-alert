package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
)

func (h *IncomingHandler) Callback(c *fiber.Ctx) error {
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
			if req.IsLiveMatch() {
				data := h.LiveMatch()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsSchedule() {
				data := h.Schedule()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsLineup() {
				data := h.Lineup()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsMatchStats() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsDisplayLiveMatch() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsFlashNews() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsCreditGoal() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampResults() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampStandings() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampTeam() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampCreditScore() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampCreditGoal() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampSMSAlerte() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampSMSAlerteEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsPrediction() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsKitFoot() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsEurope() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAfrique() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsSMSAlerteEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsFootInternational() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlerteChampMaliEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlertePremierLeagueEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlertePremierLeagueEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlerteLaLigaEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlerteLigue1Equipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlerteSerieAEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsAlerteBundesligueEquipe() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampionLeague() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsPremierLeague() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsLaLiga() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsLigue1() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsSerieA() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsBundesligua() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsChampPortugal() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			if req.IsSaudiLeague() {
				data := h.MatchStats()
				for i, s := range data {
					idx := strconv.Itoa(i + 1)
					var row string
					if strings.Contains(s.Name, "No") {
						row = s.Name
					} else {
						row = idx + ". " + s.Name
					}
					subData = append(subData, row)
				}
			}

			row := "0. Suiv"
			subData = append(subData, row)
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
		justString := strings.Join(subData, "\n")
		return c.Status(fiber.StatusOK).SendString(justString)
	}
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
