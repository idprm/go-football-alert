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
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)
	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(p.db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)
	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(p.db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)

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
		leagueService,
		teamService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		req,
	)

	if req.IsREG() {
		h.Registration()
	}

	if req.IsSTOP() {
		h.UnRegistration()
	}

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
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)
	subscriptionCreditGoalRepo := repository.NewSubscriptionCreditGoalRepository(p.db)
	subscriptionCreditGoalService := services.NewSubscriptionCreditGoalService(subscriptionCreditGoalRepo)
	subscriptionPredictWinRepo := repository.NewSubscriptionPredictWinRepository(p.db)
	subscriptionPredictWinService := services.NewSubscriptionPredictWinService(subscriptionPredictWinRepo)
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
		leagueService,
		teamService,
		subscriptionCreditGoalService,
		subscriptionPredictWinService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		verifyService,
		req,
	)

	h.Registration()

	wg.Done()
}

func (p *Processor) MO(wg *sync.WaitGroup, message []byte) {

	moRepo := repository.NewMORepository(p.db)
	moService := services.NewMOService(moRepo)

	var req *entity.MO
	json.Unmarshal([]byte(message), &req)

	h := handler.NewMOHandler(
		p.logger,
		moService,
		req,
	)

	h.Insert()

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
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(p.db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)
	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(p.db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)

	// parsing json to string
	var news *entity.News
	json.Unmarshal(message, &news)

	h := handler.NewNewsHandler(
		p.rmq,
		leagueService,
		teamService,
		newsService,
		subscriptionService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
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
	var smsAlerte *entity.SMSAlerte
	json.Unmarshal(message, &smsAlerte)

	h := handler.NewSMSAlerteHandler(
		p.rmq,
		p.logger,
		serviceService,
		subscriptionService,
		newsService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		smsAlerteService,
		smsAlerte,
	)

	// Send SMS Alerte
	h.SMSAlerte()

	wg.Done()
}

func (p *Processor) SMSActu(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	newsRepo := repository.NewNewsRepository(p.db)
	newsService := services.NewNewsService(newsRepo)
	smsActuRepo := repository.NewSMSActuRespository(p.db)
	smsActuService := services.NewSMSActuService(smsActuRepo)

	// parsing json to string
	var smsActu *entity.SMSActu
	json.Unmarshal(message, &smsActu)

	h := handler.NewSMSActuHandler(
		p.rmq,
		p.logger,
		serviceService,
		subscriptionService,
		newsService,
		smsActuService,
		smsActu,
	)

	// Send SMS Actu
	h.SMSActu()

	wg.Done()
}

func (p *Processor) SMSProno(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	smsPronoRepo := repository.NewSMSPronoRespository(p.db)
	smsPronoService := services.NewSMSPronoService(smsPronoRepo)
	pronosticRepo := repository.NewPronosticRepository(p.db)
	pronosticService := services.NewPronosticService(pronosticRepo)

	// parsing json to string
	var smsProno *entity.SMSProno
	json.Unmarshal(message, &smsProno)

	h := handler.NewSMSPronoHandler(
		p.rmq,
		p.logger,
		serviceService,
		subscriptionService,
		pronosticService,
		smsPronoService,
		smsProno,
	)

	// Send the pronostic
	h.SMSProno()

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
	subscriptionCreditGoalRepo := repository.NewSubscriptionCreditGoalRepository(p.db)
	subscriptionCreditGoalService := services.NewSubscriptionCreditGoalService(subscriptionCreditGoalRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	bettingRepo := repository.NewBettingRepository(p.db)
	bettingService := services.NewBettingService(bettingRepo)

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
		subscriptionCreditGoalService,
		transactionService,
		bettingService,
	)

	// send credit goal
	h.CreditGoal()

	wg.Done()
}

func (p *Processor) CreditScore(wg *sync.WaitGroup, message []byte) {
	/**
	 * load repo
	 */
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)
	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionCreditScoreRepo := repository.NewSubscriptionCreditScoreRepository(p.db)
	subscriptionCreditScoreService := services.NewSubscriptionCreditScoreService(subscriptionCreditScoreRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	bettingRepo := repository.NewBettingRepository(p.db)
	bettingService := services.NewBettingService(bettingRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewCreditScoreHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		subscriptionCreditScoreService,
		transactionService,
		bettingService,
	)

	// send credit goal
	h.CreditScore()

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
	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(p.db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)
	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(p.db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)
	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)

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
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		transactionService,
		leagueService,
		teamService,
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
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)

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
		leagueService,
		teamService,
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

func (p *Processor) RetryUnderpayment(wg *sync.WaitGroup, message []byte) {
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
	leagueRepo := repository.NewLeagueRepository(p.db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(p.db)
	teamService := services.NewTeamService(teamRepo)

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
		leagueService,
		teamService,
	)

	h.Underpayment()

	wg.Done()
}

func (p *Processor) Reminder48HBeforeCharging(wg *sync.WaitGroup, message []byte) {

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

	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewReminderHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
	)

	h.Remind48HBeforeCharging()

	wg.Done()
}

func (p *Processor) ReminderAfterTrialEnds(wg *sync.WaitGroup, message []byte) {

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

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewReminderHandler(
		p.rmq,
		p.logger,
		sub,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
	)

	h.RemindAfterTrialEnds()

	wg.Done()
}

func (p *Processor) PostbackMO(wg *sync.WaitGroup, message []byte) {
	wg.Done()
}
