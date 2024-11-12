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
	"github.com/idprm/go-football-alert/internal/providers/africatopsports"
	"github.com/idprm/go-football-alert/internal/providers/apifb"
	"github.com/idprm/go-football-alert/internal/providers/footmercato"
	"github.com/idprm/go-football-alert/internal/providers/madeinfoot"
	"github.com/idprm/go-football-alert/internal/providers/maxifoot"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

const (
	SRC_MAXIFOOT        string = "MAXIFOOT"
	SRC_MADEINFOOT      string = "MADEINFOOT"
	SRC_AFRICATOPSPORTS string = "AFRICATOPSPORTS"
	SRC_FOOTMERCATO     string = "FOOTMERCATO"
)

type ScraperHandler struct {
	rmq               rmqp.AMQP
	leagueService     services.ILeagueService
	teamService       services.ITeamService
	fixtureService    services.IFixtureService
	predictionService services.IPredictionService
	standingService   services.IStandingService
	lineupService     services.ILineupService
	newsService       services.INewsService
}

func NewScraperHandler(
	rmq rmqp.AMQP,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	fixtureService services.IFixtureService,
	predictionService services.IPredictionService,
	standingService services.IStandingService,
	lineupService services.ILineupService,
	newsService services.INewsService,
) *ScraperHandler {
	return &ScraperHandler{
		rmq:               rmq,
		leagueService:     leagueService,
		teamService:       teamService,
		fixtureService:    fixtureService,
		predictionService: predictionService,
		standingService:   standingService,
		lineupService:     lineupService,
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

	if len(leagues) > 0 {

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

					team, _ := h.teamService.GetByPrimaryId(el.Team.ID)
					h.teamService.SaveLeagueTeam(
						&entity.LeagueTeam{
							LeagueID: l.ID,
							TeamID:   team.GetId(),
						},
					)
				} else {
					h.teamService.UpdateByPrimaryId(
						&entity.Team{
							PrimaryID: int64(el.Team.ID),
							Slug:      slug.Make(el.Team.Name),
							Code:      el.Team.Code,
							Logo:      el.Team.Logo,
							Founded:   el.Team.Founded,
							Country:   el.Team.Country,
						},
					)
					team, _ := h.teamService.GetByPrimaryId(el.Team.ID)
					if !h.teamService.IsLeagueTeam(&entity.LeagueTeam{LeagueID: l.ID, TeamID: team.GetId()}) {
						h.teamService.SaveLeagueTeam(
							&entity.LeagueTeam{
								LeagueID: l.ID,
								TeamID:   team.GetId(),
							},
						)
					}
				}
				log.Println(string(f))
			}

		}
	}

}

func (h *ScraperHandler) Fixtures() {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		log.Println(err.Error())
	}

	if len(leagues) > 0 {
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
							Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
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
							Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
						},
					)
				}
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

	if len(fixtures) > 0 {
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
			} else {
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

					h.predictionService.UpdateByFixtureId(
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

			time.Sleep(45 * time.Second)
		}
	}
}

func (h *ScraperHandler) Standings() {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		log.Println(err.Error())
	}

	if len(leagues) > 0 {

		for _, l := range leagues {
			league, err := h.leagueService.GetByPrimaryId(int(l.PrimaryID))
			if err != nil {
				log.Println(err.Error())
			}

			f, err := fb.GetStandings(int(l.PrimaryID))
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(string(f))
			var resp model.StandingResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {
				for _, l := range el.League.Standing {
					for _, e := range l {
						team, err := h.teamService.GetByPrimaryId(e.Team.PrimaryID)
						if err != nil {
							log.Println(err.Error())
						}

						if !h.standingService.IsRank(int(league.GetId()), e.Rank) {
							updatedAt, _ := time.Parse(time.RFC3339, e.UpdateAt)
							h.standingService.Save(
								&entity.Standing{
									LeagueID:    league.GetId(),
									TeamID:      team.GetId(),
									Ranking:     e.Rank,
									TeamName:    e.Team.Name,
									Points:      e.Points,
									GoalsDiff:   e.GoalsDiff,
									Group:       e.Group,
									Form:        e.Form,
									Status:      e.Status,
									Description: e.Description,
									Played:      e.All.Played,
									Win:         e.All.Win,
									Draw:        e.All.Draw,
									Lose:        e.All.Lose,
									UpdateAt:    updatedAt,
								},
							)
						} else {
							updatedAt, _ := time.Parse(time.RFC3339, e.UpdateAt)
							h.standingService.UpdateByRank(
								&entity.Standing{
									LeagueID:    league.GetId(),
									TeamID:      team.GetId(),
									Ranking:     e.Rank,
									TeamName:    e.Team.Name,
									Points:      e.Points,
									GoalsDiff:   e.GoalsDiff,
									Group:       e.Group,
									Form:        e.Form,
									Status:      e.Status,
									Description: e.Description,
									Played:      e.All.Played,
									Win:         e.All.Win,
									Draw:        e.All.Draw,
									Lose:        e.All.Lose,
									UpdateAt:    updatedAt,
								},
							)
						}

					}
				}

			}
		}
	}

}

func (h *ScraperHandler) Lineups() {
	fb := apifb.NewApiFb()

	fixtures, err := h.fixtureService.GetAllByFixtureDate(time.Now())
	if err != nil {
		log.Println(err.Error())
	}

	if len(fixtures) > 0 {

		for _, l := range fixtures {

			f, err := fb.GetFixturesLineups(int(l.PrimaryID))
			if err != nil {
				log.Println(err.Error())
			}

			var resp model.LineupResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {
				log.Println(el)

				team, err := h.teamService.GetByPrimaryId(el.Team.PrimaryID)
				if err != nil {
					log.Println(err.Error())
				}

				// rate limit is 10 requests per minute.
				if !h.lineupService.IsLineupByFixtureId(int(l.ID)) {
					h.lineupService.Save(
						&entity.Lineup{
							LeagueID:    l.LeagueID,
							FixtureID:   l.ID,
							TeamID:      team.GetId(),
							TeamName:    el.Team.Name,
							FixtureDate: l.FixtureDate,
							Formation:   el.Formation,
						},
					)
				} else {
					h.lineupService.UpdateByFixtureId(
						&entity.Lineup{
							LeagueID:    l.LeagueID,
							FixtureID:   l.ID,
							TeamID:      team.GetId(),
							TeamName:    el.Team.Name,
							FixtureDate: l.FixtureDate,
							Formation:   el.Formation,
						},
					)
				}

			}
		}
	}

}

func (h *ScraperHandler) NewsMaxiFoot() {
	mf := maxifoot.NewMaxifoot()
	n, err := mf.GetNews()
	if err != nil {
		log.Println(err.Error())
	}

	var resp model.MaxfootRSSResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse(time.RFC1123Z, el.PubDate)

		if !h.newsService.IsNews(d, slug.Make(el.Title)) {
			news := &entity.News{
				Title:       el.Title,
				Slug:        slug.Make(el.Title),
				Description: el.Description,
				Source:      SRC_MAXIFOOT,
				PublishAt:   d,
			}
			h.newsService.Save(news)
			jsonData, err := json.Marshal(news)
			if err != nil {
				log.Println(err.Error())
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}

	}
}

func (h *ScraperHandler) NewsMadeInFoot() {
	mf := madeinfoot.NewMadeInFoot()
	n, err := mf.GetNews()
	if err != nil {
		log.Println(err.Error())
	}

	var resp model.MadeInFootRSSResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse(time.RFC1123Z, el.PubDate)

		if !h.newsService.IsNews(d, slug.Make(el.Title)) {
			news := &entity.News{
				Title:       el.Title,
				Slug:        slug.Make(el.Title),
				Description: el.Description,
				Source:      SRC_MADEINFOOT,
				PublishAt:   d,
			}

			h.newsService.Save(news)

			jsonData, err := json.Marshal(news)
			if err != nil {
				log.Println(err.Error())
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)

		}

	}
}

func (h *ScraperHandler) NewsAfricaTopSports() {
	m := africatopsports.NewAfricaTopSports()
	n, err := m.GetNews()
	if err != nil {
		log.Println(err.Error())
	}

	var resp model.AfricaTopSportsRSSResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse(time.RFC1123Z, el.PubDate)

		if !h.newsService.IsNews(d, slug.Make(el.Title)) {

			news := &entity.News{
				Title:       el.Title,
				Slug:        slug.Make(el.Title),
				Description: "-",
				Source:      SRC_AFRICATOPSPORTS,
				PublishAt:   d,
			}

			h.newsService.Save(news)

			jsonData, err := json.Marshal(news)
			if err != nil {
				log.Println(err.Error())
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}
}

func (h *ScraperHandler) NewsFootMercato() {
	m := footmercato.NewFootMercato()
	n, err := m.GetNews()
	if err != nil {
		log.Println(err.Error())
	}

	var resp model.FootMercatoSitemapResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Url.News {
		d, _ := time.Parse(time.RFC3339, el.PubDate)
		if !h.newsService.IsNews(d, slug.Make(el.Title)) {
			news := &entity.News{
				Title:       el.Title,
				Slug:        slug.Make(el.Title),
				Description: el.Keywords,
				Source:      SRC_FOOTMERCATO,
				PublishAt:   d,
			}

			h.newsService.Save(news)

			jsonData, err := json.Marshal(news)
			if err != nil {
				log.Println(err.Error())
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}
}
