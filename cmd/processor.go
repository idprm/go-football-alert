package cmd

import (
	"sync"

	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/gorm"
)

type Processor struct {
	db     *gorm.DB
	rds    *redis.Client
	rmq    rmqp.AMQP
	logger *logger.Logger
}

func NewProcessor(
	db *gorm.DB,
	rds *redis.Client,
	rmq rmqp.AMQP,
	logger *logger.Logger,
) *Processor {
	return &Processor{
		db:     db,
		rds:    rds,
		rmq:    rmq,
		logger: logger,
	}
}

func (p *Processor) MO(wg *sync.WaitGroup, message []byte) {
	/**
	 * -. Filter REG / UNREG
	 * -. Check Blacklist
	 * -. Check Active Sub
	 * -. MT API
	 * -. Save Sub
	 * -/ Save Transaction
	 */
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

	liveScoreRepo := repository.NewLiveScoreRepository(p.db)
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

	h.Fixtures()
	h.News()
}
