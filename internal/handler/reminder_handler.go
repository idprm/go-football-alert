package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/wiliehidayat87/rmqp"
)

type ReminderHandler struct {
	rmq                 rmqp.AMQP
	rds                 *redis.Client
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	pronosticService    services.IPronosticService
	subPronosticService services.ISubscriptionPronosticService
	sub                 *entity.Subscription
}

func NewReminderHandler(
	rmq rmqp.AMQP,
	rds *redis.Client,
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	pronosticService services.IPronosticService,
	subPronosticService services.ISubscriptionPronosticService,
	sub *entity.Subscription,
) *ReminderHandler {
	return &ReminderHandler{
		rmq:                 rmq,
		rds:                 rds,
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		pronosticService:    pronosticService,
		subPronosticService: subPronosticService,
		sub:                 sub,
	}
}
