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
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
}

func NewRenewalHandler(
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
	leagueService services.ILeagueService,
	teamService services.ITeamService,
) *RenewalHandler {
	return &RenewalHandler{
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
		leagueService:                   leagueService,
		teamService:                     teamService,
	}
}

func (h *RenewalHandler) Dailypush() {
	// check is active
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
		// check is renewal
		if h.subscriptionService.IsRenewal(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode()) {
			trxId := utils.GenerateTrxId()

			sub, err := h.subscriptionService.Get(h.sub.GetServiceId(), h.sub.GetMsisdn(), h.sub.GetCode())
			if err != nil {
				log.Println(err.Error())
			}

			service, err := h.serviceService.GetById(h.sub.GetServiceId())
			if err != nil {
				log.Println(err.Error())
			}

			t := telco.NewTelco(h.logger, service, h.sub, trxId)

			respBal, err := t.QueryProfileAndBal()
			if err != nil {
				log.Println(err.Error())
			}

			var respBalance *model.QueryProfileAndBalResponse
			xml.Unmarshal(respBal, &respBalance)

			// if balance enough with service price, then deduct feee
			if respBalance.IsEnoughBalance(service) {

				respDFee, err := t.DeductFee()
				if err != nil {
					log.Println(err.Error())
				}

				var respDeduct *model.DeductResponse
				xml.Unmarshal(respDFee, &respDeduct)

				if respDeduct.IsSuccess() {
					h.subscriptionService.Update(
						&entity.Subscription{
							ServiceID:          service.GetId(),
							Msisdn:             h.sub.GetMsisdn(),
							Code:               h.sub.GetCode(),
							LatestTrxId:        trxId,
							LatestSubject:      SUBJECT_RENEWAL,
							LatestStatus:       STATUS_SUCCESS,
							LatestKeyword:      h.sub.GetLatestKeyword(),
							TotalAmount:        sub.TotalAmount + service.GetPrice(),
							RenewalAt:          time.Now().AddDate(0, 0, service.GetRenewalDay()),
							ChargeAt:           time.Now(),
							TotalSuccess:       sub.TotalSuccess + 1,
							IsRetry:            false,
							TotalRenewal:       sub.TotalRenewal + 1,
							TotalAmountRenewal: sub.TotalAmountRenewal + service.GetPrice(),
							BeforeBalance:      respDeduct.GetBeforeBalanceToFloat(),
							AfterBalance:       respDeduct.GetAfterBalanceToFloat(),
							LatestPayload:      string(respDFee),
						},
					)
					// is_retry set to false
					h.subscriptionService.UpdateNotRetry(sub)
					// is_free set to false
					h.subscriptionService.UpdateNotFree(sub)

					h.transactionService.Save(
						&entity.Transaction{
							TrxId:        trxId,
							ServiceID:    service.GetId(),
							Msisdn:       h.sub.GetMsisdn(),
							Code:         h.sub.GetCode(),
							Channel:      h.sub.GetChannel(),
							Keyword:      h.sub.GetLatestKeyword(),
							Amount:       service.GetPrice(),
							Status:       STATUS_SUCCESS,
							StatusCode:   respDeduct.GetAcctResCode(),
							StatusDetail: respDeduct.GetAcctResName(),
							Subject:      SUBJECT_RENEWAL,
							Payload:      string(respDFee),
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
						Keyword:      h.sub.GetLatestKeyword(),
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

				if respDeduct.IsFailed() {
					h.subscriptionService.Update(
						&entity.Subscription{
							ServiceID:     service.GetId(),
							Msisdn:        h.sub.GetMsisdn(),
							Code:          h.sub.GetCode(),
							LatestTrxId:   trxId,
							LatestSubject: SUBJECT_RENEWAL,
							LatestStatus:  STATUS_FAILED,
							LatestKeyword: h.sub.GetLatestKeyword(),
							RenewalAt:     time.Now().AddDate(0, 0, 1),
							RetryAt:       time.Now(),
							TotalFailed:   sub.TotalFailed + 1,
							IsRetry:       true,
							LatestPayload: string(respDFee),
						},
					)

					// is_free set to false
					h.subscriptionService.UpdateNotFree(sub)

					h.transactionService.Save(
						&entity.Transaction{
							TrxId:        trxId,
							ServiceID:    service.GetId(),
							Msisdn:       h.sub.GetMsisdn(),
							Code:         h.sub.GetCode(),
							Channel:      h.sub.GetChannel(),
							Keyword:      h.sub.GetLatestKeyword(),
							Status:       STATUS_FAILED,
							StatusCode:   respDeduct.GetFaultCode(),
							StatusDetail: respDeduct.GetFaultString(),
							Subject:      SUBJECT_RENEWAL,
							Payload:      string(respDFee),
						},
					)
				}
			} else {
				h.subscriptionService.Update(
					&entity.Subscription{
						ServiceID:     service.GetId(),
						Msisdn:        h.sub.GetMsisdn(),
						Code:          h.sub.GetCode(),
						LatestTrxId:   trxId,
						LatestSubject: SUBJECT_RENEWAL,
						LatestStatus:  STATUS_FAILED,
						LatestKeyword: h.sub.GetLatestKeyword(),
						RenewalAt:     time.Now().AddDate(0, 0, 1),
						RetryAt:       time.Now(),
						TotalFailed:   sub.TotalFailed + 1,
						IsRetry:       true,
						LatestPayload: string(respBal),
					},
				)
				// is_free set to false
				h.subscriptionService.UpdateNotFree(sub)

				h.transactionService.Save(
					&entity.Transaction{
						TrxId:        trxId,
						ServiceID:    service.GetId(),
						Msisdn:       h.sub.GetMsisdn(),
						Code:         h.sub.GetCode(),
						Channel:      h.sub.GetChannel(),
						Keyword:      h.sub.GetLatestKeyword(),
						Status:       STATUS_FAILED,
						StatusCode:   "",
						StatusDetail: "INSUFF_BALANCE",
						Subject:      SUBJECT_RENEWAL,
						Payload:      string(respBal),
					},
				)

			}

		}
	}
}

func (h *RenewalHandler) getContentService(name string, service *entity.Service) (*entity.Content, error) {
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

func (h *RenewalHandler) getContentSmsAlerteService(teamOrLeague string, service *entity.Service) (*entity.Content, error) {
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
