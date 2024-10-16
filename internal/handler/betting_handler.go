package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/services"
)

type BettingHandler struct {
	rewardService services.IBettingService
}

func NewBettingHandler(
	rewardService services.IBettingService,
) *BettingHandler {
	return &BettingHandler{
		rewardService: rewardService,
	}
}

func (h *BettingHandler) GetAllPaginate(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *BettingHandler) GetBySlug(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *BettingHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *BettingHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *BettingHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
