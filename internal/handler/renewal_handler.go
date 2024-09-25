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
	"github.com/wiliehidayat87/rmqp"
)

type RenewalHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	sub                 *entity.Subscription
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	summaryService      services.ISummaryService
}

func NewRenewalHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	summaryService services.ISummaryService,
) *RenewalHandler {
	return &RenewalHandler{
		rmq:                 rmq,
		logger:              logger,
		sub:                 sub,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		summaryService:      summaryService,
	}
}

func (h *RenewalHandler) Dailypush() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContent(MT_RENEWAL)
		if err != nil {
			log.Println(err.Error())
		}

		trxId := utils.GenerateTrxId()

		summary := &entity.Summary{
			ServiceID: service.GetId(),
			CreatedAt: time.Now(),
		}

		t := telco.NewTelco(h.logger, service, h.sub)
		resp, err := t.DeductFee()
		if err != nil {
			log.Println(err.Error())
		}

		var respDeduct *model.DeductResponse
		xml.Unmarshal(utils.EscapeChar(resp), &respDeduct)

		if respDeduct.IsFailed() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     service.GetId(),
					Msisdn:        h.sub.GetMsisdn(),
					LatestTrxId:   trxId,
					LatestSubject: SUBJECT_RENEWAL,
					LatestStatus:  STATUS_FAILED,
					RenewalAt:     time.Now().AddDate(0, 0, 1),
					RetryAt:       time.Now(),
					TotalFailed:   h.sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(resp),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.sub.GetMsisdn(),
					Keyword:      h.sub.GetLatestKeyword(),
					Status:       STATUS_FAILED,
					StatusCode:   respDeduct.GetFaultCode(),
					StatusDetail: respDeduct.GetFaultString(),
					Subject:      SUBJECT_RENEWAL,
					Payload:      string(resp),
					CreatedAt:    time.Now(),
				},
			)

			// setter summary
			summary.SetTotalChargeFailed(1)
		} else {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:            service.GetId(),
					Msisdn:               h.sub.GetMsisdn(),
					LatestTrxId:          trxId,
					LatestSubject:        SUBJECT_RENEWAL,
					LatestStatus:         STATUS_SUCCESS,
					TotalAmount:          service.GetPrice(),
					RenewalAt:            time.Now().AddDate(0, 0, service.GetRenewalDay()),
					ChargeAt:             time.Now(),
					TotalSuccess:         h.sub.TotalSuccess + 1,
					IsRetry:              false,
					TotalFirstpush:       h.sub.TotalFirstpush + 1,
					TotalAmountFirstpush: service.GetPrice(),
					LatestPayload:        string(resp),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					ServiceID:    service.GetId(),
					Msisdn:       h.sub.GetMsisdn(),
					Keyword:      h.sub.GetLatestKeyword(),
					Amount:       service.GetPrice(),
					Status:       STATUS_SUCCESS,
					StatusCode:   "",
					StatusDetail: "",
					Subject:      SUBJECT_RENEWAL,
					Payload:      string(resp),
					CreatedAt:    time.Now(),
				},
			)

			// setter summary
			summary.SetTotalChargeSuccess(1)
		}

		// setter renewal
		summary.SetTotalRenewal(1)
		// summary save
		h.summaryService.Save(summary)

		// k := kannel.NewKannel(h.logger, service, content, h.sub)
		// sms, err := k.SMS(service.ScSubMT)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// var respKannel *model.KannelResponse
		// json.Unmarshal(sms, &respKannel)

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

func (h *RenewalHandler) getContent(name string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(name)
}
