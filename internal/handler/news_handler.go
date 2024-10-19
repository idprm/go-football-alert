package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type NewsHandler struct {
	leagueService services.ILeagueService
	teamService   services.ITeamService
	newsService   services.INewsService
	news          *entity.News
}

func NewNewsHandler(
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	newsService services.INewsService,
	news *entity.News,
) *NewsHandler {
	return &NewsHandler{
		leagueService: leagueService,
		teamService:   teamService,
		newsService:   newsService,
		news:          news,
	}
}

func (h *NewsHandler) Filter() {
	if h.news.IsHeadTitle() {
		if h.news.IsMatch() {
			// home
			if h.teamService.IsTeamByName(h.news.GetWithoutAccent(h.news.GetHomeTeam())) {
				team, err := h.teamService.GetByName(h.news.GetWithoutAccent(h.news.GetHomeTeam()))
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)
				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}

			// away
			if h.teamService.IsTeamByName(h.news.GetWithoutAccent(h.news.GetAwayTeam())) {
				team, err := h.teamService.GetByName(h.news.GetWithoutAccent(h.news.GetAwayTeam()))
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)
				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}

			// league
			if h.leagueService.IsLeagueByName(h.news.GetTitleWithoutAccents()) {
				league, err := h.leagueService.GetByName(h.news.GetTitleWithoutAccents())
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsLeague(
					&entity.NewsLeagues{
						NewsID:   h.news.GetId(),
						LeagueID: league.GetId(),
					},
				)
				log.Println(league)
			}

			// team
			if h.teamService.IsTeamByName(h.news.GetTitleWithoutAccents()) {
				team, err := h.teamService.GetByName(h.news.GetTitleWithoutAccents())
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)
				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}

		} else {
			if h.leagueService.IsLeagueByName(h.news.GetWithoutAccent(h.news.GetParseTitleLeft())) {
				league, err := h.leagueService.GetByName(h.news.GetWithoutAccent(h.news.GetParseTitleLeft()))
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsLeague(
					&entity.NewsLeagues{
						NewsID:   h.news.GetId(),
						LeagueID: league.GetId(),
					},
				)
				log.Println(league)
			}

			if h.leagueService.IsLeagueByName(h.news.GetWithoutAccent(h.news.GetParseTitleRight())) {
				league, err := h.leagueService.GetByName(h.news.GetWithoutAccent(h.news.GetParseTitleRight()))
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsLeague(
					&entity.NewsLeagues{
						NewsID:   h.news.GetId(),
						LeagueID: league.GetId(),
					},
				)
				log.Println(league)
			}

			if h.leagueService.IsLeagueByName(h.news.GetTitleWithoutAccents()) {
				league, err := h.leagueService.GetByName(h.news.GetTitleWithoutAccents())
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsLeague(
					&entity.NewsLeagues{
						NewsID:   h.news.GetId(),
						LeagueID: league.GetId(),
					},
				)
				log.Println(league)
			}

			if h.teamService.IsTeamByName(h.news.GetWithoutAccent(h.news.GetParseTitleLeft())) {
				team, err := h.teamService.GetByName(h.news.GetWithoutAccent(h.news.GetParseTitleLeft()))
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)

				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}

			if h.teamService.IsTeamByName(h.news.GetWithoutAccent(h.news.GetParseTitleRight())) {
				team, err := h.teamService.GetByName(h.news.GetParseTitleRight())
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)
				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}

			if h.teamService.IsTeamByName(h.news.GetTitleWithoutAccents()) {
				team, err := h.teamService.GetByName(h.news.GetTitleWithoutAccents())
				if err != nil {
					log.Println(err.Error())
				}
				// save
				h.newsService.SaveNewsTeam(
					&entity.NewsTeams{
						NewsID: h.news.GetId(),
						TeamID: team.GetId(),
					},
				)
				log.Println(team)

				// assign league by team
				if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					if err != nil {
						log.Println(err.Error())
					}

					h.newsService.SaveNewsLeague(
						&entity.NewsLeagues{
							NewsID:   h.news.GetId(),
							LeagueID: l.League.GetId(),
						},
					)
					log.Println(l.League)
				}
			}
		}
	} else {

		var title string
		if h.news.IsFootMercato() {
			title = h.news.GetWithoutAccent(h.news.GetDescription())
		} else {
			title = h.news.GetWithoutAccent(h.news.GetTitle())
		}

		if h.leagueService.IsLeagueByName(title) {
			league, err := h.leagueService.GetByName(title)
			if err != nil {
				log.Println(err.Error())
			}
			// save
			h.newsService.SaveNewsLeague(
				&entity.NewsLeagues{
					NewsID:   h.news.GetId(),
					LeagueID: league.GetId(),
				},
			)
			log.Println(league)
		}

		if h.teamService.IsTeamByName(title) {
			team, err := h.teamService.GetByName(title)
			if err != nil {
				log.Println(err.Error())
			}
			// save
			h.newsService.SaveNewsTeam(
				&entity.NewsTeams{
					NewsID: h.news.GetId(),
					TeamID: team.GetId(),
				},
			)
			log.Println(team)

			// assign league by team
			if h.teamService.IsLeagueByTeam(int(team.GetId())) {
				l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
				if err != nil {
					log.Println(err.Error())
				}

				h.newsService.SaveNewsLeague(
					&entity.NewsLeagues{
						NewsID:   h.news.GetId(),
						LeagueID: l.League.GetId(),
					},
				)
				log.Println(l.League)
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

func (h *NewsHandler) GetById(c *fiber.Ctx) error {
	news, err := h.newsService.GetById(1)
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

func (h *NewsHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *NewsHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *NewsHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
