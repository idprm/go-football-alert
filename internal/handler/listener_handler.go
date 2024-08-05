package handler

import "github.com/gofiber/fiber/v2"

type ListenerHandler struct {
}

func (h *ListenerHandler) USSD(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
