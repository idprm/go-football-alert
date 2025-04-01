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

type SMSAlerteHandler struct {
	rmq                             rmqp.AMQP
	logger                          *logger.Logger
	serviceService                  services.IServiceService
	subscriptionService             services.ISubscriptionService
	newsService                     services.INewsService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	smsAlerteService                services.ISMSAlerteService
	sub                             *entity.SMSAlerte
}

func NewSMSAlerteHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	newsService services.INewsService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	smsAlerteService services.ISMSAlerteService,
	sub *entity.SMSAlerte,
) *SMSAlerteHandler {
	return &SMSAlerteHandler{
		rmq:                             rmq,
		logger:                          logger,
		serviceService:                  serviceService,
		subscriptionService:             subscriptionService,
		newsService:                     newsService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		smsAlerteService:                smsAlerteService,
		sub:                             sub,
	}
}

func (h *SMSAlerteHandler) SMSAlerte() error {

	if !h.smsAlerteService.ISMSAlerte(h.sub.SubscriptionID, h.sub.NewsID) {
		if h.subscriptionService.IsActiveSubscriptionBySubId(h.sub.SubscriptionID) {
			// save
			h.smsAlerteService.Save(
				&entity.SMSAlerte{
					SubscriptionID: h.sub.SubscriptionID,
					NewsID:         h.sub.NewsID,
				},
			)

			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.GetBySubId(h.sub.SubscriptionID)
			if err != nil {
				return err
			}

			news, err := h.newsService.GetById(h.sub.NewsID)
			if err != nil {
				return err
			}

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

	return nil
}
