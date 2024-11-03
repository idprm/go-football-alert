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

type RetryHandler struct {
	rmq                             rmqp.AMQP
	logger                          *logger.Logger
	sub                             *entity.Subscription
	serviceService                  services.IServiceService
	contentService                  services.IContentService
	subscriptionService             services.ISubscriptionService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	transactionService              services.ITransactionService
	summaryService                  services.ISummaryService
}

func NewRetryHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	transactionService services.ITransactionService,
	summaryService services.ISummaryService,
) *RetryHandler {
	return &RetryHandler{
		rmq:                             rmq,
		logger:                          logger,
		sub:                             sub,
		serviceService:                  serviceService,
		contentService:                  contentService,
		subscriptionService:             subscriptionService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		transactionService:              transactionService,
		summaryService:                  summaryService,
	}
}

func (h *RetryHandler) Firstpush() {
	// check if active sub
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
		// check is retry
		if h.subscriptionService.IsRetry(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.Get(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode())
			if err != nil {
				log.Println(err.Error())
			}

			service, err := h.serviceService.GetById(h.sub.GetServiceId())
			if err != nil {
				log.Println(err.Error())
			}
			// smart billing set discount based on retry
			service.SetPriceWithDiscount(0)

			summary := &entity.Summary{
				ServiceID: service.GetId(),
				CreatedAt: time.Now(),
			}

			t := telco.NewTelco(h.logger, service, h.sub, trxId)

			respBal, err := t.QueryProfileAndBal()
			if err != nil {
				log.Println(err.Error())
			}

			var respBalance *model.QueryProfileAndBalResponse
			xml.Unmarshal(respBal, &respBalance)

			if respBalance.IsEnoughBalance(service) {
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
							Msisdn:               h.sub.GetMsisdn(),
							Code:                 h.sub.GetCode(),
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
					// is_free set to false
					h.subscriptionService.UpdateNotFree(sub)

					h.transactionService.Update(
						&entity.Transaction{
							TrxId:        trxId,
							ServiceID:    service.GetId(),
							Msisdn:       h.sub.GetMsisdn(),
							Code:         h.sub.GetCode(),
							Keyword:      sub.GetLatestKeyword(),
							Amount:       service.GetPrice(),
							Discount:     0,
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
				// summary save
				h.summaryService.UpdateRetry(summary)
			}
		}
	}
}

func (h *RetryHandler) Dailypush() {
	// check if active sub
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
		// check is retry
		if h.subscriptionService.IsRetry(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.Get(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode())
			if err != nil {
				log.Println(err.Error())
			}

			service, err := h.serviceService.GetById(h.sub.GetServiceId())
			if err != nil {
				log.Println(err.Error())
			}
			// smart billing set discount based on retry
			service.SetPriceWithDiscount(0)

			summary := &entity.Summary{
				ServiceID: service.GetId(),
				CreatedAt: time.Now(),
			}

			t := telco.NewTelco(h.logger, service, h.sub, trxId)
			respBal, err := t.QueryProfileAndBal()
			if err != nil {
				log.Println(err.Error())
			}

			var respBalance *model.QueryProfileAndBalResponse
			xml.Unmarshal(respBal, &respBalance)

			if respBalance.IsEnoughBalance(service) {
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
							Code:               h.sub.GetCode(),
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
					// is_free set to false
					h.subscriptionService.UpdateNotFree(sub)

					h.transactionService.Update(
						&entity.Transaction{
							TrxId:        trxId,
							ServiceID:    service.GetId(),
							Msisdn:       h.sub.GetMsisdn(),
							Code:         h.sub.GetCode(),
							Keyword:      sub.GetLatestKeyword(),
							Amount:       service.GetPrice(),
							Discount:     0,
							Status:       STATUS_SUCCESS,
							StatusCode:   respDeduct.GetAcctResCode(),
							StatusDetail: respDeduct.GetAcctResName(),
							Subject:      SUBJECT_RENEWAL,
							Payload:      string(resp),
						},
					)

					// setter summary
					summary.SetTotalChargeSuccess(1)
					summary.SetTotalRevenue(service.GetPrice())
				}

				// summary save
				h.summaryService.UpdateRetry(summary)
			}
		}
	}
}
