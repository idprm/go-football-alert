package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/providers/apifb"
	"github.com/idprm/go-football-alert/internal/providers/rss"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

const (
	SRC_MAXIFOOT        string = "MAXIFOOT"
	SRC_MADEINFOOT      string = "MADEINFOOT"
	SRC_AFRICATOPSPORTS string = "AFRICATOPSPORTS"
	SRC_FOOTMERCATO     string = "FOOTMERCATO"
	SRC_RMCSPORT        string = "RMCSPORT"
	SRC_MOBIMIUMNEWS    string = "MOBIMIUMNEWS"
)

type ScraperHandler struct {
	rmq               rmqp.AMQP
	leagueService     services.ILeagueService
	teamService       services.ITeamService
	fixtureService    services.IFixtureService
	predictionService services.IPredictionService
	standingService   services.IStandingService
	lineupService     services.ILineupService
	livematchService  services.ILiveMatchService
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
	livematchService services.ILiveMatchService,
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
		livematchService:  livematchService,
		newsService:       newsService,
	}
}

func (h *ScraperHandler) Leagues() error {
	fb := apifb.NewApiFb()
	f, err := fb.GetLeagues(2024)
	if err != nil {
		return err
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
					Name:      el.League.GetNameWithoutAccents(),
					Slug:      slug.Make(el.League.GetNameWithoutAccents()),
					Logo:      el.League.Logo,
					Country:   el.Country.Name,
				},
			)
		} else {
			h.leagueService.UpdateByPrimaryId(
				&entity.League{
					PrimaryID: int64(el.League.ID),
					Name:      el.League.GetNameWithoutAccents(),
					Slug:      slug.Make(el.League.GetNameWithoutAccents()),
					Logo:      el.League.Logo,
					Country:   el.Country.Name,
				},
			)
		}
	}
	return nil
}

func (h *ScraperHandler) Teams() error {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		return err
	}

	if len(leagues) > 0 {

		for _, l := range leagues {

			f, err := fb.GetTeams(int(l.PrimaryID), l.GetSeason())
			if err != nil {
				return err
			}

			var resp model.TeamResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {
				if !h.teamService.IsTeamByPrimaryId(el.Team.ID) {
					h.teamService.Save(
						&entity.Team{
							PrimaryID: int64(el.Team.ID),
							Name:      el.Team.GetNameWithoutAccents(),
							Slug:      slug.Make(el.Team.GetNameWithoutAccents()),
							Code:      el.Team.Code,
							Logo:      el.Team.Logo,
							Founded:   el.Team.Founded,
							Country:   el.Team.Country,
						},
					)

					if !h.teamService.IsTeamByPrimaryId(el.Team.ID) {
						team, err := h.teamService.GetByPrimaryId(el.Team.ID)
						if err != nil {
							return err
						}
						h.teamService.SaveLeagueTeam(
							&entity.LeagueTeam{
								LeagueID: l.ID,
								TeamID:   team.GetId(),
								IsActive: team.IsActive,
							},
						)
					}

				} else {
					h.teamService.UpdateByPrimaryId(
						&entity.Team{
							PrimaryID: int64(el.Team.ID),
							Name:      el.Team.GetNameWithoutAccents(),
							Slug:      slug.Make(el.Team.GetNameWithoutAccents()),
							Code:      el.Team.Code,
							Logo:      el.Team.Logo,
							Founded:   el.Team.Founded,
							Country:   el.Team.Country,
						},
					)
					team, err := h.teamService.GetByPrimaryId(el.Team.ID)
					if err != nil {
						return err
					}
					if h.teamService.IsLeagueTeam(&entity.LeagueTeam{LeagueID: l.ID, TeamID: team.GetId()}) {
						h.teamService.UpdateLeagueTeam(
							&entity.LeagueTeam{
								LeagueID: l.ID,
								TeamID:   team.GetId(),
								IsActive: team.IsActive,
							},
						)
					}

				}
				log.Println(string(f))
			}

		}
	}

	return nil
}

func (h *ScraperHandler) Fixtures() error {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		return err
	}

	if len(leagues) > 0 {
		for _, l := range leagues {
			f, err := fb.GetFixtures(int(l.PrimaryID), l.GetSeason())
			if err != nil {
				return err
			}
			log.Println(string(f))

			var resp model.FixtureResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {

				// validation
				if h.teamService.IsTeamByPrimaryId(el.Teams.Home.ID) && h.teamService.IsTeamByPrimaryId(el.Teams.Away.ID) {

					home, err := h.teamService.GetByPrimaryId(el.Teams.Home.ID)
					if err != nil {
						return err
					}

					away, err := h.teamService.GetByPrimaryId(el.Teams.Away.ID)
					if err != nil {
						return err
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
						if h.fixtureService.IsFixtureByPastTime() {
							h.fixtureService.UpdateByPrimaryId(
								&entity.Fixture{
									PrimaryID:   int64(el.Fixtures.ID),
									Timezone:    el.Fixtures.TimeZone,
									FixtureDate: fixtureDate,
									TimeStamp:   el.Fixtures.Timestamp,
									IsDone:      true,
									Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
								},
							)
						}
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

						if h.fixtureService.IsFixtureByPastTime() {
							h.fixtureService.UpdateByPrimaryId(
								&entity.Fixture{
									PrimaryID:   int64(el.Fixtures.ID),
									Timezone:    el.Fixtures.TimeZone,
									FixtureDate: fixtureDate,
									TimeStamp:   el.Fixtures.Timestamp,
									IsDone:      true,
									Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
								},
							)
						}
					}
				}
			}
		}
	}
	return nil
}

func (h *ScraperHandler) LiveMatches() error {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllByActive()
	if err != nil {
		return err
	}
	if len(leagues) > 0 {
		for _, l := range leagues {
			f, err := fb.GetLiveMatch(int(l.PrimaryID), l.GetSeason())
			if err != nil {
				return err
			}
			log.Println(string(f))

			var resp model.FixtureResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {

				fixtureDate, _ := time.Parse(time.RFC3339, el.Fixtures.Date)

				if h.fixtureService.IsFixtureByPrimaryId(el.Fixtures.ID) {

					fixture, err := h.fixtureService.GetByPrimaryId(el.Fixtures.ID)
					if err != nil {
						return err
					}

					h.fixtureService.Update(
						&entity.Fixture{
							ID:          fixture.ID,
							FixtureDate: fixtureDate,
							Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
							Elapsed:     el.Fixtures.Status.Elapsed,
						},
					)

					if h.fixtureService.IsFixtureByPastTime() {
						h.fixtureService.UpdateByPrimaryId(
							&entity.Fixture{
								PrimaryID:   int64(el.Fixtures.ID),
								IsDone:      true,
								FixtureDate: fixtureDate,
								Goal:        strconv.Itoa(el.Goals.Home) + "-" + strconv.Itoa(el.Goals.Away),
								Elapsed:     el.Fixtures.Status.Elapsed,
							},
						)
					}
				}

			}

		}
	}
	return nil
}

func (h *ScraperHandler) Predictions() error {
	fb := apifb.NewApiFb()

	fixtures, err := h.fixtureService.GetAllByFixtureDate(time.Now())
	if err != nil {
		return err
	}

	if len(fixtures) > 0 {
		for _, l := range fixtures {
			// checking league and team
			if h.leagueService.IsLeagueActiveById(int(l.LeagueID)) {

				if h.teamService.IsTeamActiveById(int(l.HomeID)) || h.teamService.IsTeamActiveById(int(l.AwayID)) {
					// rate limit is 10 requests per minute.
					if !h.predictionService.IsPredictionByFixtureId(int(l.ID)) {
						f, err := fb.GetPredictions(int(l.PrimaryID))
						if err != nil {
							return err
						}

						var resp model.PredictionResult
						json.Unmarshal(f, &resp)
						log.Println(string(f))

						for _, el := range resp.Response {
							team, err := h.teamService.GetByPrimaryId(el.Prediction.Winner.PrimaryID)
							if err != nil {
								return err
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
							return err
						}

						var resp model.PredictionResult
						json.Unmarshal(f, &resp)
						log.Println(string(f))

						for _, el := range resp.Response {
							team, err := h.teamService.GetByPrimaryId(el.Prediction.Winner.PrimaryID)
							if err != nil {
								return err
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
					time.Sleep(1 * time.Second)
				}
			}
		}
	}
	return nil
}

func (h *ScraperHandler) Standings() error {
	fb := apifb.NewApiFb()

	leagues, err := h.leagueService.GetAllUSSDByActive()
	if err != nil {
		return err
	}

	if len(leagues) > 0 {

		for _, l := range leagues {
			league, err := h.leagueService.GetByPrimaryId(int(l.PrimaryID))
			if err != nil {
				return err
			}

			f, err := fb.GetStandings(int(l.PrimaryID), l.GetSeason())
			if err != nil {
				return err
			}
			log.Println(string(f))
			var resp model.StandingResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {
				for _, l := range el.League.Standing {
					for _, e := range l {
						team, err := h.teamService.GetByPrimaryId(e.Team.PrimaryID)
						if err != nil {
							return err
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
	return nil
}

func (h *ScraperHandler) Lineups() error {
	fb := apifb.NewApiFb()

	fixtures, err := h.fixtureService.GetAllByFixtureDate(time.Now())
	if err != nil {
		return err
	}

	if len(fixtures) > 0 {

		for _, l := range fixtures {

			f, err := fb.GetFixturesLineups(int(l.PrimaryID))
			if err != nil {
				return err
			}

			var resp model.LineupResult
			json.Unmarshal(f, &resp)

			for _, el := range resp.Response {
				log.Println(el)

				team, err := h.teamService.GetByPrimaryId(el.Team.PrimaryID)
				if err != nil {
					return err
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

	return nil
}

func (h *ScraperHandler) NewsMaxiFoot() error {
	mf := rss.NewNewsRSS()
	n, err := mf.GetNews(URL_MAXFOOT)
	if err != nil {
		return err
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

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}

	return nil
}

func (h *ScraperHandler) NewsMadeInFoot() error {
	mf := rss.NewNewsRSS()
	n, err := mf.GetNews(URL_MADEINFOOT)
	if err != nil {
		return err
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

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)

		}
	}

	return nil
}

func (h *ScraperHandler) NewsAfricaTopSports() error {
	m := rss.NewNewsRSS()
	n, err := m.GetNews(URL_AFRICATOPSPORTS)
	if err != nil {
		return err
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

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}

	return nil
}

func (h *ScraperHandler) NewsFootMercato() error {
	m := rss.NewNewsRSS()
	n, err := m.GetNews(URL_FOOTMERCATO)
	if err != nil {
		return err
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

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}

	return nil
}

func (h *ScraperHandler) NewsRmcSport() error {
	m := rss.NewNewsRSS()
	n, err := m.GetNews(URL_RMCSPORT)
	if err != nil {
		return err
	}

	var resp model.RmcSportRSSResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse(time.RFC1123, el.PubDate)

		replacer := strings.NewReplacer("<![CDATA[", "", "]]>", "")

		title := strings.Trim(replacer.Replace(el.Title), " ")

		if !h.newsService.IsNews(d, slug.Make(title)) {
			news := &entity.News{
				Title:       title,
				Slug:        slug.Make(title),
				Description: el.Description,
				Source:      SRC_RMCSPORT,
				PublishAt:   d,
			}
			h.newsService.Save(news)

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}

	}

	return nil

}

func (h *ScraperHandler) MobimiumNews() error {
	m := rss.NewNewsRSS()
	n, err := m.GetNews(URL_MOBIMIUMNEWS)
	if err != nil {
		return err
	}

	var resp model.MobimiumNewsRSSResponse
	xml.Unmarshal(n, &resp)

	for _, el := range resp.Channel.Item {

		d, _ := time.Parse(time.RFC1123Z, el.PubDate)

		if !h.newsService.IsNews(d, slug.Make(el.Title)) {

			news := &entity.News{
				Title:       el.Title,
				Slug:        slug.Make(el.Title),
				Description: "-",
				Source:      SRC_MOBIMIUMNEWS,
				PublishAt:   d,
			}

			h.newsService.Save(news)

			t, err := h.newsService.Get(d, slug.Make(el.Title))
			if err != nil {
				return err
			}

			jsonData, err := json.Marshal(t)
			if err != nil {
				return err
			}

			h.rmq.IntegratePublish(
				RMQ_NEWS_EXCHANGE,
				RMQ_NEWS_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}

	return nil
}
