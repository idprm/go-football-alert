package cmd

import (
	"encoding/json"
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

func (p *Processor) USSD(wg *sync.WaitGroup, message []byte) {

	menuRepo := repository.NewMenuRepository(p.db, p.rds)
	menuService := services.NewMenuService(menuRepo)
	ussdRepo := repository.NewUssdRepository(p.db, p.rds)
	ussdService := services.NewUssdService(ussdRepo)
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

	var req *model.UssdRequest
	json.Unmarshal([]byte(message), &req)

	h := handler.NewUssdHandler(
		p.rmq,
		p.logger,
		menuService,
		ussdService,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		historyService,
		summaryService,
		req,
	)

	h.Registration()

	wg.Done()
}

func (p *Processor) SMS(wg *sync.WaitGroup, message []byte) {

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
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)
	subscriptionCreditGoalRepo := repository.NewSubscriptionCreditGoalRepository(p.db)
	subscriptionCreditGoalService := services.NewSubscriptionCreditGoalService(subscriptionCreditGoalRepo)
	subscriptionPredictRepo := repository.NewSubscriptionPredictRepository(p.db)
	subscriptionPredictService := services.NewSubscriptionPredictService(subscriptionPredictRepo)
	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(p.db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)
	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(p.db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)
	verifyRepo := repository.NewVerifyRepository(p.rds)
	verifyService := services.NewVerifyService(verifyRepo)

	var req *model.MORequest
	json.Unmarshal([]byte(message), &req)

	h := handler.NewSMSHandler(
		p.rmq,
		p.rds,
		p.logger,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		historyService,
		summaryService,
		leagueService,
		teamService,
		subscriptionCreditGoalService,
		subscriptionPredictService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		verifyService,
		req,
	)

	h.Registration()

	wg.Done()
}

func (p *Processor) MT(wg *sync.WaitGroup, message []byte) {

	mtRepo := repository.NewMTRepository(p.db)
	mtService := services.NewMTService(mtRepo)

	var req *model.MTRequest
	json.Unmarshal([]byte(message), &req)

	h := handler.NewMTHandler(
		p.rmq,
		p.logger,
		mtService,
		req,
	)

	h.MessageTerminated()

	wg.Done()
}

func (p *Processor) News(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)
	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)

	// parsing json to string
	var news *entity.News
	json.Unmarshal(message, &news)

	h := handler.NewNewsHandler(
		leagueService,
		teamService,
		newsService,
		news,
	)

	h.Filter()

	wg.Done()
}

func (p *Processor) SMSAlerte(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)
	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(p.db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)
	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(p.db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)
	smsAlerteRepo := repository.NewSMSAlerteRespository(p.db)
	smsAlerteService := services.NewSMSAlerteService(smsAlerteRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewSMSAlerteHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		subscriptionService,
		newsService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		smsAlerteService,
	)

	// Send SMS Alerte
	h.SMSAlerte()

	wg.Done()
}

func (p *Processor) Pronostic(wg *sync.WaitGroup, message []byte) {
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
	predictionRepo := repository.NewPredictionRepository(p.db)
	predictionService := services.NewPredictionService(predictionRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewPredictionHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		predictionService,
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
	bettingRepo := repository.NewBettingRepository(p.db)
	bettingService := services.NewBettingService(bettingRepo)
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
		bettingService,
		summaryService,
	)

	// send credit goal
	h.CreditGoal()

	wg.Done()
}

func (p *Processor) PredictWin(wg *sync.WaitGroup, message []byte) {
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
	predictionRepo := repository.NewPredictionRepository(p.db)
	predictionService := services.NewPredictionService(predictionRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewPredictionHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		predictionService,
	)

	// Send the prediction
	h.Prediction()

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
