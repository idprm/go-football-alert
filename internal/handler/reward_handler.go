package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/services"
)

type RewardHandler struct {
	rewardService services.IRewardService
}

func NewRewardHandler(
	rewardService services.IRewardService,
) *RewardHandler {
	return &RewardHandler{
		rewardService: rewardService,
	}
}

func (h *RewardHandler) USSD(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
