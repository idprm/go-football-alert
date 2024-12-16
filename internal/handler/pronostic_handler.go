package handler

import (
	"encoding/json"
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/wiliehidayat87/rmqp"
)

type PronosticHandler struct {
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

func NewPronosticHandler(
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
) *PronosticHandler {
	return &PronosticHandler{
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

func (h *PronosticHandler) Pronostic() {

	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetBySubId(h.sub.ID)
	if err != nil {
		log.Println(err.Error())
	}

	service, err := h.serviceService.GetById(sub.GetServiceId())
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         service.ScSubMT,
		Service:      service,
		Keyword:      sub.GetLatestKeyword(),
		Subscription: sub,
		Content:      &entity.Content{Value: ""},
	}
	mt.SetTrxId(trxId)

	jsonData, err := json.Marshal(mt)
	if err != nil {
		log.Println(err.Error())
	}

	h.rmq.IntegratePublish(
		RMQ_MT_EXCHANGE,
		RMQ_MT_QUEUE,
		RMQ_DATA_TYPE, "", string(jsonData),
	)
}
