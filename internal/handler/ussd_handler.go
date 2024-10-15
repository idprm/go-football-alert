package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/telco"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/wiliehidayat87/rmqp"
)

type UssdHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	menuService         services.IMenuService
	ussdService         services.IUssdService
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	historyService      services.IHistoryService
	summaryService      services.ISummaryService
	req                 *model.UssdRequest
}

func NewUssdHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	menuService services.IMenuService,
	ussdService services.IUssdService,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	summaryService services.ISummaryService,
	req *model.UssdRequest,
) *UssdHandler {
	return &UssdHandler{
		rmq:                 rmq,
		logger:              logger,
		menuService:         menuService,
		ussdService:         ussdService,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		historyService:      historyService,
		summaryService:      summaryService,
		req:                 req,
	}
}

func (h *UssdHandler) Registration() {
	l := h.logger.Init("ussd", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("USSD")

	/**
	 ** LiveMatch &  FlashNews & SMSAlerte
	 **/
	h.Subscription()
}

func (h *UssdHandler) Subscription() {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	var category = ""
	if h.req.IsCatLiveMatch() {
		category = SMS_LIVE_MATCH_SUB
	}
	if h.req.IsCatFlashNews() {
		category = SMS_FLASH_NEWS_SUB
	}
	if h.req.IsCatSMSAlerte() {
		category = SMS_FOLLOW_COMPETITION_SUB
	}

	content, err := h.getContent(category)
	if err != nil {
		log.Println(err)
	}

	summary := &entity.Summary{
		ServiceID: service.GetId(),
		CreatedAt: time.Now(),
	}

	subscription := &entity.Subscription{
		ServiceID:     service.GetId(),
		Category:      service.GetCategory(),
		Msisdn:        h.req.GetMsisdn(),
		LatestTrxId:   trxId,
		LatestKeyword: service.GetCategory(),
		LatestSubject: SUBJECT_FIRSTPUSH,
		IsActive:      true,
		IpAddress:     "",
	}

	if h.IsSub() {
		h.subscriptionService.Update(subscription)
	} else {
		h.subscriptionService.Save(subscription)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	if sub.IsFirstFreeDay() {
		// free 1 day
		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     service.GetId(),
				Msisdn:        h.req.GetMsisdn(),
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_FREEPUSH,
				LatestStatus:  STATUS_SUCCESS,
				RenewalAt:     time.Now().AddDate(0, 0, service.GetFreeDay()),
				LatestPayload: "-",
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.req.GetMsisdn(),
				Keyword:      service.GetCategory(),
				Amount:       0,
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_FREEPUSH,
				Payload:      "-",
			},
		)

		h.historyService.Save(
			&entity.History{
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        service.GetCategory(),
				Subject:        SUBJECT_FREEPUSH,
				Status:         STATUS_SUCCESS,
			},
		)

	} else {

		// charging if free day >= 1
		t := telco.NewTelco(h.logger, service, subscription, trxId)
		resp, err := t.DeductFee()
		if err != nil {
			log.Println(err.Error())
		}

		var respDeduct *model.DeductResponse
		xml.Unmarshal(resp, &respDeduct)

		if respDeduct.IsSuccess() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:            service.GetId(),
					Msisdn:               sub.GetMsisdn(),
					LatestTrxId:          trxId,
					LatestSubject:        SUBJECT_FIRSTPUSH,
					LatestStatus:         STATUS_SUCCESS,
					TotalAmount:          service.GetPrice(),
					RenewalAt:            time.Now().AddDate(0, 0, service.GetRenewalDay()),
					ChargeAt:             time.Now(),
					TotalSuccess:         sub.TotalSuccess + 1,
					IsRetry:              false,
					TotalFirstpush:       sub.TotalFirstpush + 1,
					TotalAmountFirstpush: sub.TotalAmountFirstpush + service.GetPrice(),
					BeforeBalance:        respDeduct.GetBeforeBalanceToFloat(),
					AfterBalance:         respDeduct.GetAfterBalanceToFloat(),
					LatestPayload:        string(resp),
				},
			)
			// is_retry set to false
			h.subscriptionService.UpdateNotRetry(sub)

			h.transactionService.Save(
				&entity.Transaction{
					ServiceID:    service.GetId(),
					Msisdn:       sub.GetMsisdn(),
					Keyword:      sub.GetLatestKeyword(),
					Amount:       service.GetPrice(),
					Status:       STATUS_SUCCESS,
					StatusCode:   respDeduct.GetAcctResCode(),
					StatusDetail: respDeduct.GetAcctResName(),
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(resp),
				},
			)
			// setter summary
			summary.SetTotalChargeSuccess(1)
			summary.SetTotalRevenue(service.GetPrice())
		}

		if respDeduct.IsFailed() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     service.GetId(),
					Msisdn:        sub.GetMsisdn(),
					LatestTrxId:   trxId,
					LatestSubject: SUBJECT_FIRSTPUSH,
					LatestStatus:  STATUS_FAILED,
					RenewalAt:     time.Now().AddDate(0, 0, 1),
					RetryAt:       time.Now(),
					TotalFailed:   sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(resp),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       sub.GetMsisdn(),
					Keyword:      sub.GetLatestKeyword(),
					Status:       STATUS_FAILED,
					StatusCode:   respDeduct.GetFaultCode(),
					StatusDetail: respDeduct.GetFaultString(),
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(resp),
				},
			)

			// setter summary
			summary.SetTotalChargeFailed(1)
		}
	}

	// setter summary
	summary.SetTotalSub(1)
	// summary save
	h.summaryService.Save(summary)

	// count total sub
	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID: service.GetId(),
			Msisdn:    h.req.GetMsisdn(),
			TotalSub:  sub.TotalSub + 1,
		},
	)

	mt := &model.MTRequest{
		Smsc:         service.ScSubMT,
		Keyword:      service.GetCategory(),
		Service:      service,
		Subscription: sub,
		Content:      content,
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

func (h *UssdHandler) IsSub() bool {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *UssdHandler) getServiceByCode(code string) (*entity.Service, error) {
	return h.serviceService.Get(code)
}

func (h *UssdHandler) getContent(v string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(v)
}
