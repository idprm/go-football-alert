package handler

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type MTHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	req                 *model.MTRequest
}

func NewMTHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	req *model.MTRequest,
) *MTHandler {
	return &MTHandler{
		rmq:                 rmq,
		logger:              logger,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		req:                 req,
	}
}

func (h *MTHandler) MessageTerminated() {
	k := kannel.NewKannel(h.logger, &entity.Service{}, h.req.Content, h.req.Subscription)
	sms, err := k.SMS(h.req.Smsc)
	if err != nil {
		log.Println(err.Error())
	}
	h.transactionService.Save(
		&entity.Transaction{
			TrxId:        "",
			ServiceID:    h.req.Subscription.GetServiceId(),
			Msisdn:       h.req.Subscription.GetMsisdn(),
			Keyword:      h.req.Subscription.GetLatestKeyword(),
			Status:       "",
			StatusCode:   "",
			StatusDetail: "",
			Subject:      "",
			Payload:      string(sms),
			CreatedAt:    time.Now(),
		},
	)
}
