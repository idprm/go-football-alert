package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type CreditScoreHandler struct {
	rmq                            rmqp.AMQP
	logger                         *logger.Logger
	sub                            *entity.Subscription
	serviceService                 services.IServiceService
	contentService                 services.IContentService
	subscriptionService            services.ISubscriptionService
	subscriptionCreditScoreService services.ISubscriptionCreditScoreService
	transactionService             services.ITransactionService
	bettingService                 services.IBettingService
}

func NewCreditScoreHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	subscriptionCreditScoreService services.ISubscriptionCreditScoreService,
	transactionService services.ITransactionService,
	bettingService services.IBettingService,
) *CreditScoreHandler {
	return &CreditScoreHandler{
		rmq:                            rmq,
		logger:                         logger,
		sub:                            sub,
		serviceService:                 serviceService,
		contentService:                 contentService,
		subscriptionService:            subscriptionService,
		subscriptionCreditScoreService: subscriptionCreditScoreService,
		transactionService:             transactionService,
		bettingService:                 bettingService,
	}
}

func (h *CreditScoreHandler) CreditScore() {

}
