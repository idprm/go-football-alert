package cmd

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
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
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	historyRepo := repository.NewHistoryRepository(p.db)
	historyService := services.NewHistoryService(historyRepo)
	summaryRepo := repository.NewSummaryRepository(p.db)
	summaryService := services.NewSummaryService(summaryRepo)

	var req *model.MORequest
	json.Unmarshal([]byte(message), &req)

	h := handler.NewMOHandler(
		p.rmq,
		p.logger,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		historyService,
		summaryService,
		req,
	)

	if h.IsService() {
		// filter REG
		if req.IsREG() {
			if !h.IsActiveSub() {
				h.Firstpush()
			} else {
				// already reg
				log.Println("ALREADY_REG")
			}
		}
		if req.IsUNREG() {
			// active sub
			if h.IsActiveSub() {
				// unsub
				h.Unsub()
			}
		}
	}

	wg.Done()
}

func (p *Processor) Renewal(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	summaryRepo := repository.NewSummaryRepository(p.db)
	summaryService := services.NewSummaryService(summaryRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewRenewalHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		summaryService,
	)

	// Dailypush MT API
	h.Dailypush()

	wg.Done()
}

func (p *Processor) Retry(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	summaryRepo := repository.NewSummaryRepository(p.db)
	summaryService := services.NewSummaryService(summaryRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewRetryHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		summaryService,
	)
	if sub.IsFirstpush() {
		if sub.IsRetryAtToday() {
			h.Firstpush()
		} else {
			h.Dailypush()
		}
	} else {
		h.Dailypush()
	}

	wg.Done()
}

func (p *Processor) Prediction(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewBulkHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		newsService,
	)

	// Send the prediction
	h.Prediction()

	wg.Done()
}

func (p *Processor) CreditGoal(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	rewardRepo := repository.NewRewardRepository(p.db)
	rewardService := services.NewRewardService(rewardRepo)
	summaryRepo := repository.NewSummaryRepository(p.db)
	summaryService := services.NewSummaryService(summaryRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewCreditGoalHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		rewardService,
		summaryService,
	)

	// send credit goal
	h.CreditGoal()

	wg.Done()
}

func (p *Processor) News(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewBulkHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		newsService,
	)

	// Send the news
	h.News()

	wg.Done()
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
