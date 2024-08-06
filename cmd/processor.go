package cmd

import (
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/services"
	"gorm.io/gorm"
)

type Processor struct {
	db *gorm.DB
}

func NewProcessor(db *gorm.DB) *Processor {
	return &Processor{
		db: db,
	}
}

func (p *Processor) Scraping() {

	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)

	seasonRepo := repository.NewSeasonRepository(p.db)
	seasonService := services.NewSeasonService(seasonRepo)

	fixtureRepo := repository.NewFixtureRepository(p.db)
	fixtureService := services.NewFixtureService(fixtureRepo)

	homeRepo := repository.NewHomeRepository(p.db)
	homeService := services.NewHomeService(homeRepo)

	awayRepo := repository.NewAwayRepository(p.db)
	awayService := services.NewAwayService(awayRepo)

	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)

	liveScoreRepo := repository.NewLivescoreRepository(p.db)
	liveScoreService := services.NewLiveScoreService(liveScoreRepo)

	predictionRepo := repository.NewPredictionRepository(p.db)
	predictionService := services.NewPredictionService(predictionRepo)

	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		leagueService,
		seasonService,
		fixtureService,
		homeService,
		awayService,
		teamService,
		liveScoreService,
		predictionService,
		newsService,
	)
}
