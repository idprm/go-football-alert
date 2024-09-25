package handler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/wiliehidayat87/rmqp"
)

type BulkHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	sub                 *entity.Subscription
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	newsService         services.INewsService
}

func NewBulkHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	newsService services.INewsService,
) *BulkHandler {
	return &BulkHandler{
		rmq:                 rmq,
		logger:              logger,
		sub:                 sub,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		newsService:         newsService,
	}
}

func (h *BulkHandler) Prediction() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContent(SMS_PREDICT_SUB)
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     service.GetId(),
				Msisdn:        h.sub.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_PREDICTION,
				LatestStatus:  STATUS_SUCCESS,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.sub.GetMsisdn(),
				Keyword:      h.sub.GetLatestKeyword(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_PREDICTION,
				Payload:      "-",
				CreatedAt:    time.Now(),
			},
		)

		mt := &model.MTRequest{
			Smsc:         "",
			Subscription: h.sub,
			Content:      content,
		}

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

func (h *BulkHandler) FollowCompetition() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContent(SMS_FOLLOW_COMPETITION_SUB)
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     service.GetId(),
				Msisdn:        h.sub.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_NEWS,
				LatestStatus:  STATUS_SUCCESS,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.sub.GetMsisdn(),
				Keyword:      h.sub.GetLatestKeyword(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_NEWS,
				Payload:      "-",
				CreatedAt:    time.Now(),
			},
		)

		mt := &model.MTRequest{
			Smsc:         "",
			Subscription: h.sub,
			Content:      content,
		}

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

func (h *BulkHandler) FollowTeam() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContent(SMS_FOLLOW_TEAM_SUB)
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     service.GetId(),
				Msisdn:        h.sub.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_NEWS,
				LatestStatus:  STATUS_SUCCESS,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.sub.GetMsisdn(),
				Keyword:      h.sub.GetLatestKeyword(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_NEWS,
				Payload:      "-",
				CreatedAt:    time.Now(),
			},
		)

		mt := &model.MTRequest{
			Smsc:         "",
			Subscription: h.sub,
			Content:      content,
		}

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

func (h *BulkHandler) getContent(name string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(name)
}
