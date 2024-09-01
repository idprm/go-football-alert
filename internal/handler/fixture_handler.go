package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/services"
)

type FixtureHandler struct {
	fixtureService services.IFixtureService
}

func NewFixtureHandler(
	fixtureService services.IFixtureService,
) *FixtureHandler {
	return &FixtureHandler{
		fixtureService: fixtureService,
	}
}

func (h *FixtureHandler) USSD(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
