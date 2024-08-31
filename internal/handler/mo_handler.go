package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
	"github.com/idprm/go-football-alert/internal/providers/telco"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/wiliehidayat87/rmqp"
)

type MOHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	historyService      services.IHistoryService
	req                 *model.MORequest
}

func NewMOHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	req *model.MORequest,
) *MOHandler {
	return &MOHandler{
		rmq:                 rmq,
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		historyService:      historyService,
		req:                 req,
	}
}

func (h *MOHandler) Firstpush() {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}

	content, err := h.getContentFirstpush(service.GetId())
	if err != nil {
		log.Println(err)
	}

	trxId := utils.GenerateTrxId()

	subscription := &entity.Subscription{
		CountryID:     service.GetCountryId(),
		ServiceID:     service.GetId(),
		Category:      service.GetCategory(),
		Msisdn:        h.req.GetMsisdn(),
		LatestTrxId:   trxId,
		LatestKeyword: h.req.GetKeyword(),
		LatestSubject: SUBJECT_FIRSTPUSH,
		IsActive:      true,
		IpAddress:     h.req.GetIpAddress(),
	}

	if h.IsSub() {
		subscription.UpdatedAt = time.Now()
		h.subscriptionService.Update(subscription)

	} else {
		subscription.CreatedAt = time.Now()
		subscription.UpdatedAt = time.Now()
		h.subscriptionService.Save(subscription)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	t := telco.NewTelco(h.logger, service, subscription)
	resp, err := t.DeductFee()
	if err != nil {
		log.Println(err.Error())
	}

	var respDeduct *model.DeductResponse
	xml.Unmarshal(utils.EscapeChar(resp), &respDeduct)

	if !respDeduct.IsFailed() {
		h.subscriptionService.Update(
			&entity.Subscription{
				CountryID:            service.GetCountryId(),
				ServiceID:            service.GetId(),
				Msisdn:               h.req.GetMsisdn(),
				LatestTrxId:          trxId,
				LatestSubject:        SUBJECT_FIRSTPUSH,
				LatestStatus:         STATUS_SUCCESS,
				TotalAmount:          service.GetPrice(),
				RenewalAt:            time.Now().AddDate(0, 0, service.GetRenewalDay()),
				ChargeAt:             time.Now(),
				TotalSuccess:         sub.TotalSuccess + 1,
				IsRetry:              false,
				TotalFirstpush:       sub.TotalFirstpush + 1,
				TotalAmountFirstpush: service.GetPrice(),
				LatestPayload:        string(resp),
				UpdatedAt:            time.Now(),
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:          trxId,
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetKeyword(),
				Amount:         service.GetPrice(),
				Status:         STATUS_SUCCESS,
				StatusCode:     "",
				StatusDetail:   "",
				Subject:        SUBJECT_FIRSTPUSH,
				Payload:        string(resp),
				CreatedAt:      time.Now(),
			},
		)

		h.historyService.Save(
			&entity.History{
				CountryID: service.GetCountryId(),
				ServiceID: service.GetId(),
				Msisdn:    h.req.GetMsisdn(),
				Keyword:   h.req.GetKeyword(),
				Subject:   SUBJECT_FIRSTPUSH,
				Status:    STATUS_SUCCESS,
				CreatedAt: time.Now(),
			},
		)
	} else {
		h.subscriptionService.Update(
			&entity.Subscription{
				CountryID:     service.GetCountryId(),
				ServiceID:     service.GetId(),
				Msisdn:        h.req.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_FIRSTPUSH,
				LatestStatus:  STATUS_FAILED,
				RenewalAt:     time.Now().AddDate(0, 0, 1),
				RetryAt:       time.Now(),
				TotalFailed:   sub.TotalFailed + 1,
				IsRetry:       true,
				LatestPayload: string(resp),
				UpdatedAt:     time.Now(),
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:          trxId,
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetKeyword(),
				Status:         STATUS_FAILED,
				StatusCode:     "",
				StatusDetail:   "",
				Subject:        SUBJECT_FIRSTPUSH,
				Payload:        string(resp),
				CreatedAt:      time.Now(),
			},
		)

		h.historyService.Save(
			&entity.History{
				CountryID: service.GetCountryId(),
				ServiceID: service.GetId(),
				Msisdn:    h.req.GetMsisdn(),
				Keyword:   h.req.GetKeyword(),
				Subject:   SUBJECT_FIRSTPUSH,
				Status:    STATUS_FAILED,
				CreatedAt: time.Now(),
			},
		)
	}

	k := kannel.NewKannel(h.logger, service, content, subscription)
	sms, err := k.SMS()
	if err != nil {
		log.Println(err.Error())
	}

	var respKannel *model.KannelResponse
	json.Unmarshal(sms, &respKannel)

	msg := &entity.Transaction{
		TrxId:          trxId,
		CountryID:      service.GetCountryId(),
		SubscriptionID: sub.GetId(),
		ServiceID:      subscription.GetServiceId(),
		Msisdn:         subscription.GetMsisdn(),
		Amount:         0,
		StatusCode:     "-",
		StatusDetail:   "-",
		Subject:        SUBJECT_FP_SMS,
		IpAddress:      h.req.GetIpAddress(),
		Payload:        string(sms),
		CreatedAt:      time.Now(),
	}
	if utils.IsSMSSuccess(respKannel.Message) {
		msg.SetStatus(STATUS_SUCCESS)
		h.transactionService.Save(msg)
	} else {
		msg.SetStatus(STATUS_FAILED)
		h.transactionService.Save(msg)
	}

	// count total sub
	h.subscriptionService.Update(
		&entity.Subscription{
			CountryID: service.GetCountryId(),
			ServiceID: service.GetId(),
			Msisdn:    h.req.GetMsisdn(),
			TotalSub:  sub.TotalSub + 1,
			UpdatedAt: time.Now(),
		},
	)

}

func (h *MOHandler) Unsub() {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}
	trxId := utils.GenerateTrxId()

	h.subscriptionService.Update(
		&entity.Subscription{
			CountryID:     service.GetCountryId(),
			ServiceID:     service.GetId(),
			Msisdn:        h.req.GetMsisdn(),
			LatestTrxId:   trxId,
			LatestSubject: SUBJECT_UNSUB,
			LatestStatus:  STATUS_SUCCESS,
			UnsubAt:       time.Now(),
			IpAddress:     h.req.GetIpAddress(),
			IsRetry:       false,
			IsActive:      false,
			UpdatedAt:     time.Now(),
		},
	)

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err)
	}

	h.subscriptionService.Update(
		&entity.Subscription{
			CountryID:  service.GetCountryId(),
			ServiceID:  service.GetId(),
			Msisdn:     h.req.GetMsisdn(),
			TotalUnsub: sub.TotalUnsub + 1,
			UpdatedAt:  time.Now(),
		},
	)

	h.transactionService.Save(
		&entity.Transaction{
			TrxId:          trxId,
			CountryID:      service.GetCountryId(),
			SubscriptionID: sub.GetId(),
			ServiceID:      service.GetId(),
			Msisdn:         h.req.GetMsisdn(),
			Status:         STATUS_SUCCESS,
			StatusCode:     "",
			StatusDetail:   "",
			Subject:        SUBJECT_UNSUB,
			Payload:        "",
			CreatedAt:      time.Now(),
		},
	)

	h.historyService.Save(
		&entity.History{
			CountryID: service.GetCountryId(),
			ServiceID: service.GetId(),
			Msisdn:    h.req.GetMsisdn(),
			Keyword:   h.req.GetKeyword(),
			Subject:   SUBJECT_UNSUB,
			Status:    STATUS_SUCCESS,
			IpAddress: h.req.GetIpAddress(),
			CreatedAt: time.Now(),
		},
	)
}

func (h *MOHandler) IsActiveSub() bool {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *MOHandler) IsSub() bool {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *MOHandler) IsService() bool {
	return h.serviceService.IsService(h.req.GetSubKeyword())
}

func (h *MOHandler) getService() (*entity.Service, error) {
	return h.serviceService.Get(h.req.GetSubKeyword())
}

func (h *MOHandler) getContentFirstpush(serviceId int) (*entity.Content, error) {
	return h.contentService.Get(serviceId, MT_FIRSTPUSH)
}
