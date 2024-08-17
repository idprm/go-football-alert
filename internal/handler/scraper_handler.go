package handler

import (
	"encoding/json"
	"log"

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

	log.Println(resp.Results)
	log.Println(resp.Response.Fixtures)

	// log.Println(resp.FixturesResponse)
	// for index, element := range resp.Response.ResponseFixtures {
	// 	log.Println(index)
	// 	log.Panicln(element)
	// }

	// if !h.fixtureService.IsFixtureByPrimaryId(resp.ID) {

	// 	if !h.homeService.IsHomeByPrimaryId(resp.Teams.Home.ID) {
	// 		if !h.teamService.IsTeam(slug.Make(resp.Teams.Home.Name)) {
	// 			h.teamService.Save(
	// 				&entity.Team{
	// 					Name: resp.Teams.Home.Name,
	// 					Slug: slug.Make(resp.Teams.Home.Name),
	// 					Logo: resp.Teams.Home.Logo,
	// 				},
	// 			)
	// 		}

	// 		team, err := h.teamService.Get(slug.Make(resp.Teams.Home.Name))
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}

	// 		h.homeService.Save(
	// 			&entity.Home{
	// 				PrimaryID: int64(resp.Teams.Home.ID),
	// 				TeamID:    team.GetId(),
	// 				Goal:      0,
	// 				IsWinner:  resp.Teams.Home.Winner,
	// 			},
	// 		)
	// 	}

	// 	if !h.awayService.IsAwayByPrimaryId(resp.Teams.Away.ID) {
	// 		if !h.teamService.IsTeam(slug.Make(resp.Teams.Away.Name)) {
	// 			h.teamService.Save(
	// 				&entity.Team{
	// 					Name: resp.Teams.Away.Name,
	// 					Slug: slug.Make(resp.Teams.Away.Name),
	// 					Logo: resp.Teams.Away.Logo,
	// 				},
	// 			)
	// 		}

	// 		team, err := h.teamService.Get(slug.Make(resp.Teams.Away.Name))
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}

	// 		h.awayService.Save(
	// 			&entity.Away{
	// 				PrimaryID: int64(resp.Teams.Away.ID),
	// 				TeamID:    team.GetId(),
	// 				Goal:      0,
	// 				IsWinner:  resp.Teams.Away.Winner,
	// 			},
	// 		)
	// 	}

	// 	h.fixtureService.Save(
	// 		&entity.Fixture{
	// 			PrimaryID: int64(resp.ID),
	// 			Timezone:  resp.TimeZone,
	// 			Date:      resp.Date,
	// 			TimeStamp: resp.Timestamp,
	// 		},
	// 	)
	// }

}
