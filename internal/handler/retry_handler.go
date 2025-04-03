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

type RetryHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	sub                 *entity.Subscription
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	summaryService      services.ISummaryService
	leagueService       services.ILeagueService
	teamService         services.ITeamService
}

func NewRetryHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	summaryService services.ISummaryService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
) *RetryHandler {
	return &RetryHandler{
		rmq:                 rmq,
		logger:              logger,
		sub:                 sub,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		summaryService:      summaryService,
		leagueService:       leagueService,
		teamService:         teamService,
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
							Channel:      h.sub.GetChannel(),
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
				}

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
							Channel:      h.sub.GetChannel(),
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

					var content *entity.Content

					if service.IsSmsAlerteCompetition() {
						if h.leagueService.IsLeagueByCode(h.sub.GetCode()) {
							league, err := h.leagueService.GetByCode(h.sub.GetCode())
							if err != nil {
								log.Println(err.Error())
							}
							content, err = h.getContentSmsAlerteService(league.GetName(), service)
							if err != nil {
								log.Println(err.Error())
							}
						}
					} else if service.IsSmsAlerteEquipe() {
						if h.teamService.IsTeamByCode(h.sub.GetCode()) {
							team, err := h.teamService.GetByCode(h.sub.GetCode())
							if err != nil {
								log.Println(err.Error())
							}
							content, err = h.getContentSmsAlerteService(team.GetName(), service)
							if err != nil {
								log.Println(err.Error())
							}
						}
					} else {
						content, err = h.getContentService(SMS_SUCCESS_CHARGING, service)
						if err != nil {
							log.Println(err.Error())
						}
					}

					mt := &model.MTRequest{
						Smsc:         service.ScSubMT,
						Service:      service,
						Keyword:      sub.GetLatestKeyword(),
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

			}
		}
	}
}

func (h *RetryHandler) getContentService(name string, service *entity.Service) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetService(name, service)
}

func (h *RetryHandler) getContentSmsAlerteService(teamOrLeague string, service *entity.Service) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(SMS_SUCCESS_CHARGING_SMS_ALERTE) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetSMSAlerte(SMS_SUCCESS_CHARGING_SMS_ALERTE, teamOrLeague, service)
}
