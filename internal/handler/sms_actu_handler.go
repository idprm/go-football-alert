package handler

import (
	"encoding/json"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/wiliehidayat87/rmqp"
)

type SMSActuHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
	newsService         services.INewsService
	smsActuService      services.ISMSActuService
	sub                 *entity.SMSActu
}

func NewSMSActuHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	newsService services.INewsService,
	smsActuService services.ISMSActuService,
	sub *entity.SMSActu,
) *SMSActuHandler {
	return &SMSActuHandler{
		rmq:                 rmq,
		logger:              logger,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
		newsService:         newsService,
		smsActuService:      smsActuService,
		sub:                 sub,
	}
}

func (h *SMSActuHandler) SMSActu() error {
	// xxxxx
	if !h.smsActuService.ISMSActu(h.sub.Msisdn, h.sub.NewsID) {
		// xxxxx
		if h.subscriptionService.IsActiveSubscriptionBySMSAlerteMsisdn(h.sub.Msisdn) {
			// save
			h.smsActuService.Save(
				&entity.SMSActu{
					Msisdn: h.sub.Msisdn,
					NewsID: h.sub.NewsID,
				},
			)

			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.GetActiveBySMSAlerteMsisdn(h.sub.Msisdn)
			if err != nil {
				return err
			}

			news, err := h.newsService.GetById(h.sub.NewsID)
			if err != nil {
				return err
			}

			if h.serviceService.IsServiceById(sub.GetServiceId()) {
				service, err := h.serviceService.GetById(sub.GetServiceId())
				if err != nil {
					return err
				}

				mt := &model.MTRequest{
					Smsc:         service.ScSubMT,
					Service:      service,
					Keyword:      sub.GetLatestKeyword(),
					Subscription: sub,
					Content:      &entity.Content{Value: news.GetTitleWithoutAccents()},
				}
				mt.SetTrxId(trxId)

				jsonData, err := json.Marshal(mt)
				if err != nil {
					return err
				}

				h.rmq.IntegratePublish(
					RMQ_MT_EXCHANGE,
					RMQ_MT_QUEUE,
					RMQ_DATA_TYPE, "", string(jsonData),
				)
			}
		}
	}

	return nil
}
