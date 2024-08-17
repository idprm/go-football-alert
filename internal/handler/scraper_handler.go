package handler

import (
	"encoding/json"
	"log"

	"github.com/gosimple/slug"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/providers/apifb"
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

	for _, element := range resp.Response {
		log.Println(element.Fixtures.ID)

		if !h.fixtureService.IsFixtureByPrimaryId(element.Fixtures.ID) {

			if !h.homeService.IsHomeByPrimaryId(element.Teams.Home.ID) {
				if !h.teamService.IsTeam(slug.Make(element.Teams.Home.Name)) {
					h.teamService.Save(
						&entity.Team{
							Name: element.Teams.Home.Name,
							Slug: slug.Make(element.Teams.Home.Name),
							Logo: element.Teams.Home.Logo,
						},
					)
				}

				team, err := h.teamService.Get(slug.Make(element.Teams.Home.Name))
				if err != nil {
					log.Println(err.Error())
				}

				h.homeService.Save(
					&entity.Home{
						PrimaryID: int64(element.Teams.Home.ID),
						TeamID:    team.GetId(),
						Goal:      0,
						IsWinner:  element.Teams.Home.Winner,
					},
				)
			}

			if !h.awayService.IsAwayByPrimaryId(element.Teams.Away.ID) {
				if !h.teamService.IsTeam(slug.Make(element.Teams.Away.Name)) {
					h.teamService.Save(
						&entity.Team{
							Name: element.Teams.Away.Name,
							Slug: slug.Make(element.Teams.Away.Name),
							Logo: element.Teams.Away.Logo,
						},
					)
				}

				team, err := h.teamService.Get(slug.Make(element.Teams.Away.Name))
				if err != nil {
					log.Println(err.Error())
				}

				h.awayService.Save(
					&entity.Away{
						PrimaryID: int64(element.Teams.Away.ID),
						TeamID:    team.GetId(),
						Goal:      0,
						IsWinner:  element.Teams.Away.Winner,
					},
				)
			}
		}

		home, err := h.homeService.GetByPrimaryId(element.Teams.Home.ID)
		if err != nil {
			log.Println(err.Error())
		}

		away, err := h.awayService.GetByPrimaryId(element.Teams.Away.ID)
		if err != nil {
			log.Println(err.Error())
		}

		if !h.leagueService.IsLeagueByPrimaryId(element.League.ID) {
			h.leagueService.Save(
				&entity.League{
					PrimaryID: int64(element.League.ID),
					Name:      element.League.Name,
					Slug:      slug.Make(element.League.Name),
					Logo:      element.League.Logo,
					Country:   element.League.Country,
				},
			)
		}

		league, err := h.leagueService.GetByPrimaryId(element.League.ID)
		if err != nil {
			log.Println(err.Error())
		}

		h.fixtureService.Save(
			&entity.Fixture{
				PrimaryID: int64(element.Fixtures.ID),
				Timezone:  element.Fixtures.TimeZone,
				Date:      element.Fixtures.Date,
				TimeStamp: element.Fixtures.Timestamp,
				LeagueID:  league.ID,
				HomeID:    home.ID,
				AwayID:    away.ID,
			},
		)
	}
}
