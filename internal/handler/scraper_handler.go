package handler

import (
	"encoding/json"
	"log"

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

	var resp model.FixturesResponse
	json.Unmarshal(f, &resp)

	if !h.fixtureService.IsFixtureByPrimaryId(resp.ID) {

		if !h.homeService.IsHomeByPrimaryId(resp.Teams.Home.ID) {
			h.homeService.Save(
				&entity.Home{
					PrimaryID: int64(resp.Teams.Home.ID),
					TeamID:    1,
					Goal:      0,
					IsWinner:  resp.Teams.Home.Winner,
				})
		}

		h.fixtureService.Save(
			&entity.Fixture{
				PrimaryID: int64(resp.ID),
				Timezone:  resp.TimeZone,
				Date:      resp.Date,
				TimeStamp: resp.Timestamp,
			},
		)
	}

}
