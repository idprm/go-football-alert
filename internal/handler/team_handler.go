package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type TeamHandler struct {
	teamService services.ITeamService
}

func NewTeamHandler(
	teamService services.ITeamService,
) *TeamHandler {
	return &TeamHandler{
		teamService: teamService,
	}
}

func (h *TeamHandler) GetAllPaginate(c *fiber.Ctx) error {
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

	teams, err := h.teamService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(teams)
}

func (h *TeamHandler) GetBySlug(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *TeamHandler) Update(c *fiber.Ctx) error {
	req := new(model.TeamRequest)

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

	errors := ValidateStruct(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if !h.teamService.IsTeamByPrimaryId(int(req.PrimaryID)) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	h.teamService.UpdateByPrimaryId(
		&entity.Team{
			PrimaryID: req.PrimaryID,
			Keyword:   req.Keyword,
		},
	)

	return c.Status(fiber.StatusOK).JSON(
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "updated",
		},
	)
}

func (h *TeamHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
