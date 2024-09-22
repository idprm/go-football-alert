package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"strconv"
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
	teamService       services.ITeamService
	fixtureService    services.IFixtureService
	predictionService services.IPredictionService
	standingService   services.IStandingService
	newsService       services.INewsService
}

func NewScraperHandler(
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	fixtureService services.IFixtureService,
	predictionService services.IPredictionService,
	standingService services.IStandingService,
	newsService services.INewsService,
) *ScraperHandler {
	return &ScraperHandler{
		leagueService:     leagueService,
		teamService:       teamService,
		fixtureService:    fixtureService,
		predictionService: predictionService,
		standingService:   standingService,
		newsService:       newsService,
	}
}

func (h *ScraperHandler) Leagues() {
	fb := apifb.NewApiFb()
	f, err := fb.GetLeagues()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(f))

	var resp model.LeagueResult
	json.Unmarshal(f, &resp)

	for _, el := range resp.Response {
		log.Println(el.Country)
		if !h.leagueService.IsLeagueByPrimaryId(el.League.ID) {
			h.leagueService.Save(
				&entity.League{
					PrimaryID: int64(el.League.ID),
					Name:      el.League.Name,
					Slug:      slug.Make(el.League.Name),
					Logo:      el.League.Logo,
					Country:   el.Country.Name,
				},
			)
		} else {
			h.leagueService.UpdateByPrimaryId(
				&entity.League{
					PrimaryID: int64(el.League.ID),
					Name:      el.League.Name,
					Slug:      slug.Make(el.League.Name),
					Logo:      el.League.Logo,
					Country:   el.Country.Name,
				},
			)
		}
	}
}

func (h *ScraperHandler) Teams() {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		log.Println(err.Error())
	}

	for _, l := range leagues {

		f, err := fb.GetTeams(int(l.PrimaryID))
		if err != nil {
			log.Println(err.Error())
		}

		var resp model.TeamResult
		json.Unmarshal(f, &resp)

		for _, el := range resp.Response {
			if !h.teamService.IsTeamByPrimaryId(el.Team.ID) {
				h.teamService.Save(
					&entity.Team{
						PrimaryID: int64(el.Team.ID),
						Name:      el.Team.Name,
						Slug:      slug.Make(el.Team.Name),
						Code:      el.Team.Code,
						Logo:      el.Team.Logo,
						Founded:   el.Team.Founded,
						Country:   el.Team.Country,
					},
				)
			} else {
				h.teamService.UpdateByPrimaryId(
					&entity.Team{
						PrimaryID: int64(el.Team.ID),
						Name:      el.Team.Name,
						Slug:      slug.Make(el.Team.Name),
						Code:      el.Team.Code,
						Logo:      el.Team.Logo,
						Founded:   el.Team.Founded,
						Country:   el.Team.Country,
					},
				)
			}
			log.Println(string(f))
		}

	}
}

func (h *ScraperHandler) Fixtures() {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		log.Println(err.Error())
	}

	for _, l := range leagues {
		f, err := fb.GetFixtures(int(l.PrimaryID))
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(f))

		var resp model.FixtureResult
		json.Unmarshal(f, &resp)

		for _, el := range resp.Response {

			home, err := h.teamService.GetByPrimaryId(el.Teams.Home.ID)
			if err != nil {
				log.Println(err.Error())
			}

			away, err := h.teamService.GetByPrimaryId(el.Teams.Away.ID)
			if err != nil {
				log.Println(err.Error())
			}

			fixtureDate, _ := time.Parse(time.RFC3339, el.Fixtures.Date)

			if !h.fixtureService.IsFixtureByPrimaryId(el.Fixtures.ID) {
				h.fixtureService.Save(
					&entity.Fixture{
						PrimaryID:   int64(el.Fixtures.ID),
						Timezone:    el.Fixtures.TimeZone,
						FixtureDate: fixtureDate,
						TimeStamp:   el.Fixtures.Timestamp,
						LeagueID:    l.ID,
						HomeID:      home.ID,
						AwayID:      away.ID,
						Goal:        strconv.Itoa(el.Goals.Home) + "|" + strconv.Itoa(el.Goals.Away),
					},
				)
			} else {
				h.fixtureService.UpdateByPrimaryId(
					&entity.Fixture{
						PrimaryID:   int64(el.Fixtures.ID),
						Timezone:    el.Fixtures.TimeZone,
						FixtureDate: fixtureDate,
						TimeStamp:   el.Fixtures.Timestamp,
						LeagueID:    l.ID,
						HomeID:      home.ID,
						AwayID:      away.ID,
						Goal:        strconv.Itoa(el.Goals.Home) + "|" + strconv.Itoa(el.Goals.Away),
					},
				)
			}
		}

	}
}

func (h *ScraperHandler) Predictions() {

	fb := apifb.NewApiFb()

	fixtures, err := h.fixtureService.GetAllByFixtureDate(time.Now())
	if err != nil {
		log.Println(err.Error())
	}

	for _, l := range fixtures {
		// rate limit is 10 requests per minute.
		if !h.predictionService.IsPredictionByFixtureId(int(l.ID)) {
			f, err := fb.GetPredictions(int(l.PrimaryID))
			if err != nil {
				log.Println(err.Error())
			}

			var resp model.PredictionResult
			json.Unmarshal(f, &resp)
			log.Println(string(f))

			for _, el := range resp.Response {
				team, err := h.teamService.GetByPrimaryId(el.Prediction.Winner.PrimaryID)
				if err != nil {
					log.Println(err.Error())
				}

				h.predictionService.Save(
					&entity.Prediction{
						FixtureID:     l.ID,
						FixtureDate:   l.FixtureDate,
						WinnerID:      team.ID,
						WinnerName:    el.Prediction.Winner.Name,
						WinnerComment: el.Prediction.Winner.Comment,
						Advice:        el.Prediction.Advice,
						PercentHome:   el.Prediction.Percent.Home,
						PercentDraw:   el.Prediction.Percent.Draw,
						PercentAway:   el.Prediction.Percent.Away,
					},
				)
			}
		}

	}

}

func (h *ScraperHandler) Standings() {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		log.Println(err.Error())
	}

	for _, l := range leagues {

		// league, err := h.leagueService.GetByPrimaryId(int(l.PrimaryID))
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		f, err := fb.GetStandings(int(l.PrimaryID))
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(f))
		var resp model.StandingResult
		json.Unmarshal(f, &resp)

		for _, el := range resp.Response {

			log.Println(el)
			// for _, ol := range el.League.Standing {

			// 	team, err := h.teamService.GetByPrimaryId(ol.Team.PrimaryID)
			// 	if err != nil {
			// 		log.Println(err.Error())
			// 	}

			// 	if !h.standingService.IsRank(ol.Rank) {
			// 		updatedAt, _ := time.Parse(time.RFC3339, ol.UpdateAt)
			// 		h.standingService.Save(
			// 			&entity.Standing{
			// 				Rank:        ol.Rank,
			// 				LeagueID:    league.GetId(),
			// 				TeamID:      team.GetId(),
			// 				TeamName:    ol.Team.Name,
			// 				Points:      ol.Team.Points,
			// 				GoalsDiff:   ol.Team.GoalsDiff,
			// 				Group:       ol.Team.Group,
			// 				Form:        ol.Team.Form,
			// 				Status:      ol.Team.Status,
			// 				Description: ol.Team.Description,
			// 				Played:      ol.All.Played,
			// 				Win:         ol.All.Win,
			// 				Draw:        ol.All.Draw,
			// 				Lose:        ol.All.Lose,
			// 				UpdateAt:    updatedAt,
			// 			},
			// 		)
			// 	} else {
			// 		updatedAt, _ := time.Parse(time.RFC3339, ol.UpdateAt)
			// 		h.standingService.UpdateByRank(
			// 			&entity.Standing{
			// 				Rank:        ol.Rank,
			// 				LeagueID:    league.GetId(),
			// 				TeamID:      team.GetId(),
			// 				TeamName:    ol.Team.Name,
			// 				Points:      ol.Team.Points,
			// 				GoalsDiff:   ol.Team.GoalsDiff,
			// 				Group:       ol.Team.Group,
			// 				Form:        ol.Team.Form,
			// 				Status:      ol.Team.Status,
			// 				Description: ol.Team.Description,
			// 				Played:      ol.All.Played,
			// 				Win:         ol.All.Win,
			// 				Draw:        ol.All.Draw,
			// 				Lose:        ol.All.Lose,
			// 				UpdateAt:    updatedAt,
			// 			},
			// 		)
			// 	}
			// }

		}

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

		d, _ := time.Parse(time.RFC1123Z, el.PubDate)

		if !h.newsService.IsNews(slug.Make(el.Title), d.Format("2006-01-02")) {
			h.newsService.Save(
				&entity.News{
					LeagueID:    1,
					TeamID:      1,
					Title:       el.Title,
					Slug:        slug.Make(el.Title),
					Description: el.Description,
					Source:      "MAXIFOOT",
					PublishAt:   d,
				},
			)
		}

	}
}
