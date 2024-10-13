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

type SMSAlerteHandler struct {
	rmq                             rmqp.AMQP
	logger                          *logger.Logger
	sub                             *entity.Subscription
	serviceService                  services.IServiceService
	subscriptionService             services.ISubscriptionService
	newsService                     services.INewsService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	smsAlerteService                services.ISMSAlerteService
}

func NewSMSAlerteHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	newsService services.INewsService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	smsAlerteService services.ISMSAlerteService,
) *SMSAlerteHandler {
	return &SMSAlerteHandler{
		rmq:                             rmq,
		logger:                          logger,
		sub:                             sub,
		serviceService:                  serviceService,
		subscriptionService:             subscriptionService,
		newsService:                     newsService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		smsAlerteService:                smsAlerteService,
	}
}

func (h *SMSAlerteHandler) SMSAlerte() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		trxId := utils.GenerateTrxId()

		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		if h.subscriptionFollowLeagueService.IsSub(h.sub.GetId()) {
			sl, err := h.subscriptionFollowLeagueService.GetBySub(h.sub.GetId())
			if err != nil {
				log.Println(err.Error())
			}

			// news league
			newsLeague, _ := h.newsService.GetAllNewsLeague(sl.LeagueID)
			for _, n := range newsLeague {
				// subId, newsId
				if !h.smsAlerteService.ISMSAlerte(h.sub.GetId(), n.NewsID) {
					h.smsAlerteService.Save(
						&entity.SMSAlerte{
							SubscriptionID: h.sub.GetId(),
							NewsID:         n.NewsID,
						},
					)
					mt := &model.MTRequest{
						Smsc:         service.ScSubMT,
						Service:      service,
						Subscription: h.sub,
						Content:      &entity.Content{Value: n.News.GetTitleWithoutAccents()},
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

		if h.subscriptionFollowTeamService.IsSub(h.sub.GetId()) {
			st, err := h.subscriptionFollowTeamService.GetBySub(h.sub.GetId())
			if err != nil {
				log.Println(err.Error())
			}

			// news team
			newsTeam, _ := h.newsService.GetAllNewsTeam(st.TeamID)
			for _, n := range newsTeam {
				// subId, newsId
				if !h.smsAlerteService.ISMSAlerte(h.sub.GetId(), n.NewsID) {
					h.smsAlerteService.Save(
						&entity.SMSAlerte{
							SubscriptionID: h.sub.GetId(),
							NewsID:         n.NewsID,
						},
					)
					mt := &model.MTRequest{
						Smsc:         service.ScSubMT,
						Service:      service,
						Subscription: h.sub,
						Content:      &entity.Content{Value: n.News.GetTitleWithoutAccents()},
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
	}
}
