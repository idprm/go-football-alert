package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/services"
)

type LeagueHandler struct {
	leagueService services.ILeagueService
}

func NewLeagueHandler(
	leagueService services.ILeagueService,
) *LeagueHandler {
	return &LeagueHandler{
		leagueService: leagueService,
	}
}

func (h *LeagueHandler) USSD(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
