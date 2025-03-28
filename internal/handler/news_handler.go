package handler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type NewsHandler struct {
	rmq                             rmqp.AMQP
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
	newsService                     services.INewsService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	news                            *entity.News
}

func NewNewsHandler(
	rmq rmqp.AMQP,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	newsService services.INewsService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	news *entity.News,
) *NewsHandler {
	return &NewsHandler{
		rmq:                             rmq,
		leagueService:                   leagueService,
		teamService:                     teamService,
		newsService:                     newsService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		news:                            news,
	}
}

func (h *NewsHandler) Filter() {
	if h.news.IsHeadTitle() {
		if h.news.IsMatch() {
			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetHomeTeam())) {
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

					h.SMSAlerteTeam(team.GetId())
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetAwayTeam())) {
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

					h.SMSAlerteTeam(team.GetId())
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetTitleWithoutAccents()) {
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

					h.SMSAlerteLeague(league.GetId())
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetTitleWithoutAccents()) {
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
					h.SMSAlerteTeam(team.GetId())

					// assign league by team
					// if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					// 	l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					// 	if err != nil {
					// 		log.Println(err.Error())
					// 	}

					// 	h.newsService.SaveNewsLeague(
					// 		&entity.NewsLeagues{
					// 			NewsID:   h.news.GetId(),
					// 			LeagueID: l.League.GetId(),
					// 		},
					// 	)
					// 	h.SMSAlerteLeague(l.League.GetId())
					// }
				}

			}

		} else {
			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetParseTitleLeft())) {
				// league
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
					h.SMSAlerteLeague(league.GetId())

				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetParseTitleRight())) {
				// league
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
					h.SMSAlerteLeague(league.GetId())
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetTitleWithoutAccents()) {
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
					h.SMSAlerteLeague(league.GetId())
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetParseTitleLeft())) {
				// team
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

					h.SMSAlerteTeam(team.GetId())

					// assign league by team
					// if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					// 	l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					// 	if err != nil {
					// 		log.Println(err.Error())
					// 	}

					// 	h.newsService.SaveNewsLeague(
					// 		&entity.NewsLeagues{
					// 			NewsID:   h.news.GetId(),
					// 			LeagueID: l.League.GetId(),
					// 		},
					// 	)
					// 	h.SMSAlerteLeague(l.League.GetId())
					// }
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetWithoutAccent(h.news.GetParseTitleRight())) {
				// team
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
					h.SMSAlerteTeam(team.GetId())

					// assign league by team
					// if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					// 	l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					// 	if err != nil {
					// 		log.Println(err.Error())
					// 	}

					// 	h.newsService.SaveNewsLeague(
					// 		&entity.NewsLeagues{
					// 			NewsID:   h.news.GetId(),
					// 			LeagueID: l.League.GetId(),
					// 		},
					// 	)
					// 	h.SMSAlerteLeague(l.League.GetId())
					// }
				}
			}

			// filtered
			if !h.news.IsFilteredKeyword(h.news.GetTitleWithoutAccents()) {
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
					h.SMSAlerteTeam(team.GetId())

					// assign league by team
					// if h.teamService.IsLeagueByTeam(int(team.GetId())) {
					// 	l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
					// 	if err != nil {
					// 		log.Println(err.Error())
					// 	}

					// 	h.newsService.SaveNewsLeague(
					// 		&entity.NewsLeagues{
					// 			NewsID:   h.news.GetId(),
					// 			LeagueID: l.League.GetId(),
					// 		},
					// 	)
					// 	h.SMSAlerteLeague(l.League.GetId())
					// }
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

		// filtered
		if !h.news.IsFilteredKeyword(title) {
			// league
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
				h.SMSAlerteLeague(league.GetId())
			}
		}

		// filtered
		if !h.news.IsFilteredKeyword(title) {
			// team
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
				h.SMSAlerteTeam(team.GetId())

				// assign league by team
				// if h.teamService.IsLeagueByTeam(int(team.GetId())) {
				// 	l, err := h.teamService.GetLeagueByTeam(int(team.GetId()))
				// 	if err != nil {
				// 		log.Println(err.Error())
				// 	}

				// 	h.newsService.SaveNewsLeague(
				// 		&entity.NewsLeagues{
				// 			NewsID:   h.news.GetId(),
				// 			LeagueID: l.League.GetId(),
				// 		},
				// 	)
				// 	h.SMSAlerteLeague(l.League.GetId())
				// }
			}
		}
	}
}

func (h *NewsHandler) SMSAlerteLeague(leagueId int64) {
	// valid in league
	subs := h.subscriptionFollowLeagueService.GetAllSubByLeague(leagueId)

	if len(*subs) > 0 {
		for _, s := range *subs {
			// is time
			if h.IsTime() {
				// counter
				h.subscriptionFollowLeagueService.Sent(&entity.SubscriptionFollowLeague{SubscriptionID: s.SubscriptionID, LeagueID: leagueId})
				// limit
				if h.subscriptionFollowLeagueService.IsLimit(s.SubscriptionID, leagueId) {
					jsonData, err := json.Marshal(&entity.SMSAlerte{SubscriptionID: s.SubscriptionID, NewsID: h.news.GetId()})
					if err != nil {
						log.Println(err.Error())
					}

					h.rmq.IntegratePublish(
						RMQ_SMS_ALERTE_EXCHANGE,
						RMQ_SMS_ALERTE_QUEUE,
						RMQ_DATA_TYPE, "", string(jsonData),
					)
				}
			}

		}
	}
}

func (h *NewsHandler) SMSAlerteTeam(teamId int64) {
	// valid in team
	subs := h.subscriptionFollowTeamService.GetAllSubByTeam(teamId)

	if len(*subs) > 0 {
		for _, s := range *subs {
			// is time
			if h.IsTime() {
				// counter
				h.subscriptionFollowTeamService.Sent(&entity.SubscriptionFollowTeam{SubscriptionID: s.SubscriptionID, TeamID: teamId})
				// limit
				if h.subscriptionFollowTeamService.IsLimit(s.SubscriptionID, teamId) {
					jsonData, err := json.Marshal(&entity.SMSAlerte{SubscriptionID: s.SubscriptionID, NewsID: h.news.GetId()})
					if err != nil {
						log.Println(err.Error())
					}

					h.rmq.IntegratePublish(
						RMQ_SMS_ALERTE_EXCHANGE,
						RMQ_SMS_ALERTE_QUEUE,
						RMQ_DATA_TYPE, "", string(jsonData),
					)
				}
			}
		}
	}
}

func (h *NewsHandler) IsTime() bool {
	// from 9am to 10pm
	start, _ := time.Parse("15:04", "08:00")
	end, _ := time.Parse("15:04", "22:30")

	t, err := time.Parse("15:04", time.Now().Format("15:04"))
	if err != nil {
		log.Fatal(err)
	}
	return start.Before(t) && end.After(t)
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
