package handler

import (
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

		trxId := utils.GenerateTrxId()

		sub, err := h.subscriptionService.Get(h.sub.GetServiceId(), h.sub.GetMsisdn())
		if err != nil {
			log.Println(err.Error())
		}

		service, err := h.serviceService.GetById(h.sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		summary := &entity.Summary{
			ServiceID: service.GetId(),
			CreatedAt: time.Now(),
		}

		t := telco.NewTelco(h.logger, service, h.sub, trxId)
		resp, err := t.DeductFee()
		if err != nil {
			log.Println(err.Error())
		}

		var respDeduct *model.DeductResponse
		xml.Unmarshal(resp, &respDeduct)

		if respDeduct.IsSuccess() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:          service.GetId(),
					Msisdn:             h.sub.GetMsisdn(),
					LatestTrxId:        trxId,
					LatestSubject:      SUBJECT_RENEWAL,
					LatestStatus:       STATUS_SUCCESS,
					TotalAmount:        service.GetPrice(),
					RenewalAt:          time.Now().AddDate(0, 0, service.GetRenewalDay()),
					ChargeAt:           time.Now(),
					TotalSuccess:       sub.TotalSuccess + 1,
					IsRetry:            false,
					TotalRenewal:       sub.TotalRenewal + 1,
					TotalAmountRenewal: sub.TotalAmountRenewal + service.GetPrice(),
					BeforeBalance:      respDeduct.GetBeforeBalanceToFloat(),
					AfterBalance:       respDeduct.GetAfterBalanceToFloat(),
					LatestPayload:      string(resp),
				},
			)
			// is_retry set to false
			h.subscriptionService.UpdateNotRetry(sub)

			h.transactionService.Save(
				&entity.Transaction{
					ServiceID:    service.GetId(),
					Msisdn:       h.sub.GetMsisdn(),
					Keyword:      sub.GetLatestKeyword(),
					Amount:       service.GetPrice(),
					Status:       STATUS_SUCCESS,
					StatusCode:   respDeduct.GetAcctResCode(),
					StatusDetail: respDeduct.GetAcctResName(),
					Subject:      SUBJECT_RENEWAL,
					Payload:      string(resp),
				},
			)
			// setter summary
			summary.SetTotalChargeSuccess(1)
		}

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
					TotalFailed:   sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(resp),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.sub.GetMsisdn(),
					Keyword:      sub.GetLatestKeyword(),
					Status:       STATUS_FAILED,
					StatusCode:   respDeduct.GetFaultCode(),
					StatusDetail: respDeduct.GetFaultString(),
					Subject:      SUBJECT_RENEWAL,
					Payload:      string(resp),
				},
			)

			// setter summary
			summary.SetTotalChargeFailed(1)
		}

		// setter renewal
		summary.SetTotalRenewal(1)
		// summary save
		h.summaryService.Save(summary)

	}
}
