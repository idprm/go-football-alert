package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
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

func (h *FixtureHandler) GetAllPaginate(c *fiber.Ctx) error {
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

	fixtures, err := h.fixtureService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(fixtures)
}

func (h *FixtureHandler) GetAllCurrent(c *fiber.Ctx) error {
	fixtures, err := h.fixtureService.GetAllCurrent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	if len(fixtures) > 0 {
		return c.Status(fiber.StatusOK).JSON(fixtures)
	}

	return c.Status(fiber.StatusNotFound).JSON(
		&model.WebResponse{
			Error:      true,
			StatusCode: fiber.StatusNotFound,
			Message:    "not_found",
		},
	)
}

func (h *FixtureHandler) Get(c *fiber.Ctx) error {

	if !h.fixtureService.IsFixture(1) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	fixture, err := h.fixtureService.Get(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(fixture)
}

func (h *FixtureHandler) Save(c *fiber.Ctx) error {
	req := new(model.WebResponse)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    "bad_request",
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "OK"})
}

func (h *FixtureHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *FixtureHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
