package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
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

func (h *LeagueHandler) GetAllPaginate(c *fiber.Ctx) error {
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

	leagues, err := h.leagueService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(leagues)
}

func (h *LeagueHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if !h.leagueService.IsLeague(slug) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	league, err := h.leagueService.Get(slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(league)
}

func (h *LeagueHandler) Update(c *fiber.Ctx) error {
	req := new(model.LeagueRequest)

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

	if !h.leagueService.IsLeagueByPrimaryId(int(req.PrimaryID)) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	h.leagueService.Save(
		&entity.League{
			PrimaryID: req.PrimaryID,
			Keyword:   req.Keyword,
			IsActive:  req.IsActive,
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

func (h *LeagueHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
