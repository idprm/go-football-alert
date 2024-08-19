package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"time"

	"github.com/gosimple/slug"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/providers/apifb"
	"github.com/idprm/go-football-alert/internal/providers/maxifoot"
	"github.com/idprm/go-football-alert/internal/services"
)

type ScraperHandler struct {
	leagueService     services.ILeagueService
	seasonService     services.ISeasonService
	fixtureService    services.IFixtureService
	homeService       services.IHomeService
	awayService       services.IAwayService
	teamService       services.ITeamService
	livescoreService  services.ILiveScoreService
	predictionService services.IPredictionService
	newsService       services.INewsService
}

func NewScraperHandler(
	leagueService services.ILeagueService,
	seasonService services.ISeasonService,
	fixtureService services.IFixtureService,
	homeService services.IHomeService,
	awayService services.IAwayService,
	teamService services.ITeamService,
	livescoreService services.ILiveScoreService,
	predictionService services.IPredictionService,
	newsService services.INewsService,
) *ScraperHandler {
	return &ScraperHandler{
		leagueService:     leagueService,
		seasonService:     seasonService,
		fixtureService:    fixtureService,
		homeService:       homeService,
		awayService:       awayService,
		teamService:       teamService,
		livescoreService:  livescoreService,
		predictionService: predictionService,
		newsService:       newsService,
	}
}

func (h *ScraperHandler) Leagues() {
}

func (h *ScraperHandler) Teams() {
	fb := apifb.NewApiFb()
	f, err := fb.GetFixtures()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(f)
}

func (h *ScraperHandler) Fixtures() {
	fb := apifb.NewApiFb()
	f, err := fb.GetFixtures()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(f))

	var resp model.FixtureResult
	json.Unmarshal(f, &resp)

	for _, el := range resp.Response {
		log.Println(el.Fixtures.ID)

		if !h.fixtureService.IsFixtureByPrimaryId(el.Fixtures.ID) {

			if !h.homeService.IsHomeByPrimaryId(el.Teams.Home.ID) {
				if !h.teamService.IsTeam(slug.Make(el.Teams.Home.Name)) {
					h.teamService.Save(
						&entity.Team{
							Name: el.Teams.Home.Name,
							Slug: slug.Make(el.Teams.Home.Name),
							Logo: el.Teams.Home.Logo,
						},
					)
				}

				team, err := h.teamService.Get(slug.Make(el.Teams.Home.Name))
				if err != nil {
					log.Println(err.Error())
				}

				h.homeService.Save(
					&entity.Home{
						PrimaryID: int64(el.Teams.Home.ID),
						TeamID:    team.GetId(),
						Goal:      0,
						IsWinner:  el.Teams.Home.Winner,
					},
				)
			}

			if !h.awayService.IsAwayByPrimaryId(el.Teams.Away.ID) {
				if !h.teamService.IsTeam(slug.Make(el.Teams.Away.Name)) {
					h.teamService.Save(
						&entity.Team{
							Name: el.Teams.Away.Name,
							Slug: slug.Make(el.Teams.Away.Name),
							Logo: el.Teams.Away.Logo,
						},
					)
				}

				team, err := h.teamService.Get(slug.Make(el.Teams.Away.Name))
				if err != nil {
					log.Println(err.Error())
				}

				h.awayService.Save(
					&entity.Away{
						PrimaryID: int64(el.Teams.Away.ID),
						TeamID:    team.GetId(),
						Goal:      0,
						IsWinner:  el.Teams.Away.Winner,
					},
				)
			}
		}

		home, err := h.homeService.GetByPrimaryId(el.Teams.Home.ID)
		if err != nil {
			log.Println(err.Error())
		}

		away, err := h.awayService.GetByPrimaryId(el.Teams.Away.ID)
		if err != nil {
			log.Println(err.Error())
		}

		if !h.leagueService.IsLeagueByPrimaryId(el.League.ID) {
			h.leagueService.Save(
				&entity.League{
					PrimaryID: int64(el.League.ID),
					Name:      el.League.Name,
					Slug:      slug.Make(el.League.Name),
					Logo:      el.League.Logo,
					Country:   el.League.Country,
				},
			)
		}

		league, err := h.leagueService.GetByPrimaryId(el.League.ID)
		if err != nil {
			log.Println(err.Error())
		}

		h.fixtureService.Save(
			&entity.Fixture{
				PrimaryID: int64(el.Fixtures.ID),
				Timezone:  el.Fixtures.TimeZone,
				Date:      el.Fixtures.Date,
				TimeStamp: el.Fixtures.Timestamp,
				LeagueID:  league.ID,
				HomeID:    home.ID,
				AwayID:    away.ID,
			},
		)
	}
}

func (h *ScraperHandler) News() {
	mf := maxifoot.NewMaxifoot()
	n, err := mf.GetNews()
	if err != nil {
		log.Println(err.Error())
	}
	var resp model.MaxfootRSSResponse
	xml.Unmarshal(n, &resp)
	log.Println()

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 +0200", el.PubDate)

		if !h.newsService.IsNews(slug.Make(el.Title), d.Format("2006-01-02")) {
			h.newsService.Save(
				&entity.News{
					Title:       el.Title,
					Slug:        slug.Make(el.Title),
					Description: el.Description,
					PublishAt:   d,
				},
			)
		}

	}
}
