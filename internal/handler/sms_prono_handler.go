package handler

import (
	"encoding/json"
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/wiliehidayat87/rmqp"
)

type SMSPronoHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
	pronosticService    services.IPronosticService
	smsPronoService     services.ISMSPronoService
	sub                 *entity.SMSProno
}

func NewSMSPronoHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	pronosticService services.IPronosticService,
	smsPronoService services.ISMSPronoService,
	sub *entity.SMSProno,
) *SMSPronoHandler {
	return &SMSPronoHandler{
		rmq:                 rmq,
		logger:              logger,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
		pronosticService:    pronosticService,
		smsPronoService:     smsPronoService,
		sub:                 sub,
	}
}

func (h *SMSPronoHandler) SMSProno() {

	if !h.smsPronoService.ISMSProno(h.sub.SubscriptionID, h.sub.PronosticID) {

		if h.subscriptionService.IsActiveSubscriptionBySubId(h.sub.SubscriptionID) {
			// save
			h.smsPronoService.Save(
				&entity.SMSProno{
					SubscriptionID: h.sub.SubscriptionID,
					PronosticID:    h.sub.PronosticID,
				},
			)

			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.GetBySubId(h.sub.SubscriptionID)
			if err != nil {
				log.Println(err.Error())
			}

			prono, err := h.pronosticService.GetById(h.sub.PronosticID)
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
				Content:      &entity.Content{Value: prono.GetValue()},
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
	}
}
