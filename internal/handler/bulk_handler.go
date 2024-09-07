package handler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
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

		content, err := h.getContentPrediction(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		k := kannel.NewKannel(h.logger, service, content, h.sub)
		sms, err := k.SMS(service.ScSubMT)
		if err != nil {
			log.Println(err.Error())
		}

		var respKannel *model.KannelResponse
		json.Unmarshal(sms, &respKannel)

		h.subscriptionService.Update(
			&entity.Subscription{
				CountryID:     service.GetCountryId(),
				ServiceID:     service.GetId(),
				Msisdn:        h.sub.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_PREDICTION,
				LatestStatus:  STATUS_SUCCESS,
				UpdatedAt:     time.Now(),
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:          trxId,
				CountryID:      service.GetCountryId(),
				SubscriptionID: h.sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.sub.GetMsisdn(),
				Keyword:        h.sub.GetLatestKeyword(),
				Status:         STATUS_SUCCESS,
				StatusCode:     "",
				StatusDetail:   "",
				Subject:        SUBJECT_PREDICTION,
				Payload:        "-",
				CreatedAt:      time.Now(),
			},
		)
	}
}

func (h *BulkHandler) News() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContentNews(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		k := kannel.NewKannel(h.logger, service, content, h.sub)
		sms, err := k.SMS(service.ScSubMT)
		if err != nil {
			log.Println(err.Error())
		}

		var respKannel *model.KannelResponse
		json.Unmarshal(sms, &respKannel)

		h.subscriptionService.Update(
			&entity.Subscription{
				CountryID:     service.GetCountryId(),
				ServiceID:     service.GetId(),
				Msisdn:        h.sub.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_NEWS,
				LatestStatus:  STATUS_SUCCESS,
				UpdatedAt:     time.Now(),
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:          trxId,
				CountryID:      service.GetCountryId(),
				SubscriptionID: h.sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.sub.GetMsisdn(),
				Keyword:        h.sub.GetLatestKeyword(),
				Status:         STATUS_SUCCESS,
				StatusCode:     "",
				StatusDetail:   "",
				Subject:        SUBJECT_NEWS,
				Payload:        "-",
				CreatedAt:      time.Now(),
			},
		)
	}
}

func (h *BulkHandler) getContentPrediction(serviceId int) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(serviceId, MT_PREDICTION) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(serviceId, MT_PREDICTION)
}

func (h *BulkHandler) getContentNews(serviceId int) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(serviceId, MT_NEWS) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(serviceId, MT_NEWS)
}
