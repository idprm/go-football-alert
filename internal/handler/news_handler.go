package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type NewsHandler struct {
	leagueService            services.ILeagueService
	teamService              services.ITeamService
	newsService              services.INewsService
	followCompetitionService services.IFollowCompetitionService
	followTeamService        services.IFollowTeamService
	news                     *entity.News
}

func NewNewsHandler(
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	newsService services.INewsService,
	followCompetitionService services.IFollowCompetitionService,
	followTeamService services.IFollowTeamService,
	news *entity.News,
) *NewsHandler {
	return &NewsHandler{
		leagueService:            leagueService,
		teamService:              teamService,
		newsService:              newsService,
		followCompetitionService: followCompetitionService,
		followTeamService:        followTeamService,
		news:                     news,
	}
}

func (h *NewsHandler) Mapping() {
	if h.news.IsParseTitle() {
		fmt.Println(h.news.GetParseTitle())
		if h.news.IsMatch() {
			// home
			if h.teamService.IsTeamByName(h.news.GetHomeTeam()) {
				team, err := h.teamService.GetByName(h.news.GetHomeTeam())
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(team)
			}
			// away
			if h.teamService.IsTeamByName(h.news.GetAwayTeam()) {
				team, err := h.teamService.GetByName(h.news.GetAwayTeam())
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(team)
			}
		} else {
			// income data by slug
			if h.leagueService.IsLeagueByName(h.news.GetParseTitle()) {
				league, err := h.leagueService.GetByName(h.news.GetParseTitle())
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(league)

			}

			if h.teamService.IsTeamByName(h.news.GetParseTitle()) {
				team, err := h.teamService.GetByName(h.news.GetParseTitle())
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(team)
			}
		}

	}

}

func (h *NewsHandler) GetAllPaginate(c *fiber.Ctx) error {
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

	news, err := h.newsService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(news)
}

func (h *NewsHandler) GetBySlug(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *NewsHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *NewsHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *NewsHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
