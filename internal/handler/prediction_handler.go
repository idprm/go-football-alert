package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type PredictionHandler struct {
	predictionService services.IPredictionService
}

func NewPredictionHandler(
	predictionService services.IPredictionService,
) *PredictionHandler {
	return &PredictionHandler{
		predictionService: predictionService,
	}
}

func (h *PredictionHandler) GetAllPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	predictions, err := h.predictionService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(predictions)
}

func (h *PredictionHandler) GetBySlug(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *PredictionHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *PredictionHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *PredictionHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}