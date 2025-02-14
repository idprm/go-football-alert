package handler

import (
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/wiliehidayat87/rmqp"
)

type MOHandler struct {
	rmq                             rmqp.AMQP
	rds                             *redis.Client
	logger                          *logger.Logger
	serviceService                  services.IServiceService
	contentService                  services.IContentService
	subscriptionService             services.ISubscriptionService
	transactionService              services.ITransactionService
	historyService                  services.IHistoryService
	summaryService                  services.ISummaryService
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
	subscriptionCreditGoalService   services.ISubscriptionCreditGoalService
	subscriptionPredictWinService   services.ISubscriptionPredictWinService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	verifyService                   services.IVerifyService
}

func NewMOHandler(
	rmq rmqp.AMQP,
	rds *redis.Client,
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	summaryService services.ISummaryService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	subscriptionCreditGoalService services.ISubscriptionCreditGoalService,
	subscriptionPredictWinService services.ISubscriptionPredictWinService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	verifyService services.IVerifyService,
) *MOHandler {
	return &MOHandler{
		rmq:                             rmq,
		rds:                             rds,
		logger:                          logger,
		serviceService:                  serviceService,
		contentService:                  contentService,
		subscriptionService:             subscriptionService,
		transactionService:              transactionService,
		historyService:                  historyService,
		summaryService:                  summaryService,
		leagueService:                   leagueService,
		teamService:                     teamService,
		subscriptionCreditGoalService:   subscriptionCreditGoalService,
		subscriptionPredictWinService:   subscriptionPredictWinService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		verifyService:                   verifyService,
	}
}

func (h *MOHandler) Firstpush() {

}

func (h *MOHandler) Unsub() {

}
