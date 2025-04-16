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
	rmq                             rmqp.AMQP
	logger                          *logger.Logger
	menuService                     services.IMenuService
	ussdService                     services.IUssdService
	serviceService                  services.IServiceService
	contentService                  services.IContentService
	subscriptionService             services.ISubscriptionService
	transactionService              services.ITransactionService
	historyService                  services.IHistoryService
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	req                             *model.UssdRequest
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
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	req *model.UssdRequest,
) *UssdHandler {
	return &UssdHandler{
		rmq:                             rmq,
		logger:                          logger,
		menuService:                     menuService,
		ussdService:                     ussdService,
		serviceService:                  serviceService,
		contentService:                  contentService,
		subscriptionService:             subscriptionService,
		transactionService:              transactionService,
		historyService:                  historyService,
		leagueService:                   leagueService,
		teamService:                     teamService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		req:                             req,
	}
}

const (
	SMS_LIVE_MATCH_SUB         string = "LIVE_MATCH_SUB"
	SMS_LIVE_MATCH_ALREADY_SUB string = "LIVE_MATCH_ALREADY_SUB"
	SMS_FLASH_NEWS_SUB         string = "FLASH_NEWS_SUB"
	SMS_FLASH_NEWS_ALREADY_SUB string = "FLASH_NEWS_ALREADY_SUB"
)

func (h *UssdHandler) Registration() {
	l := h.logger.Init("ussd", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("USSD")

	/**
	 ** LiveMatch & FlashNews & SMSAlerte
	 **/

	if h.req.IsCatLiveMatch() {
		if !h.IsActiveSubByNonSMSAlerte(CATEGORY_LIVEMATCH) {
			h.SubLiveMatch()
		}
	}

	if h.req.IsCatFlashNews() {
		if !h.IsActiveSubByNonSMSAlerte(CATEGORY_FLASHNEWS) {
			h.SubFlashNews()
		}
	}

	if h.req.IsCatPronostic() {
		if !h.IsActiveSubByNonSMSAlerte(CATEGORY_PRONOSTIC) {
			h.SubPronostic()
		}
	}

	if h.req.IsCatSMSAlerteCompetition() {
		if h.leagueService.IsLeagueByCode(h.req.GetUniqueCode()) {
			league, err := h.leagueService.GetByCode(h.req.GetUniqueCode())
			if err != nil {
				log.Println(err.Error())
			}
			if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE_COMPETITION, league.GetCode()) {
				h.SubAlerteCompetition(league)
			}
		}
	}

	if h.req.IsCatSMSAlerteEquipe() {
		if h.teamService.IsTeamByCode(h.req.GetUniqueCode()) {
			team, err := h.teamService.GetByCode(h.req.GetUniqueCode())
			if err != nil {
				log.Println(err.Error())
			}
			if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE_EQUIPE, team.GetCode()) {
				h.SubAlerteEquipe(team)
			}
		}

	}

}

func (h *UssdHandler) UnRegistration() {
	l := h.logger.Init("ussd", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("USSD")

	/**
	 ** LiveMatch & FlashNews & SMSAlerte
	 **/
	if h.req.IsCatLiveMatch() {
		if h.IsActiveSubByNonSMSAlerte(CATEGORY_LIVEMATCH) {
			h.StopNonSMSAlerte(CATEGORY_LIVEMATCH)
		}
	}

	if h.req.IsCatFlashNews() {
		if h.IsActiveSubByNonSMSAlerte(CATEGORY_FLASHNEWS) {
			h.StopNonSMSAlerte(CATEGORY_FLASHNEWS)
		}
	}

	if h.req.IsCatPronostic() {
		if h.IsActiveSubByNonSMSAlerte(CATEGORY_PRONOSTIC) {
			h.StopNonSMSAlerte(CATEGORY_PRONOSTIC)
		}
	}

	if h.req.IsCatSMSAlerteCompetition() {
		if h.leagueService.IsLeagueByCode(h.req.GetUniqueCode()) {
			league, err := h.leagueService.GetByCode(h.req.GetUniqueCode())
			if err != nil {
				log.Println(err.Error())
			}
			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE_COMPETITION, league.GetCode()) {
				h.StopAlerteCompetition(league)
			}
		}
	}

	if h.req.IsCatSMSAlerteEquipe() {
		if h.teamService.IsTeamByCode(h.req.GetUniqueCode()) {
			team, err := h.teamService.GetByCode(h.req.GetUniqueCode())
			if err != nil {
				log.Println(err.Error())
			}
			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE_EQUIPE, team.GetCode()) {
				h.StopAlerteEquipe(team)
			}
		}
	}
}

/**
** SUB ON USSD
**/
func (h *UssdHandler) SubLiveMatch() {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentLiveMatch(service)
	if err != nil {
		log.Println(err)
	}

	h.Firstpush(CATEGORY_LIVEMATCH, service, service.GetCode(), content)
}

func (h *UssdHandler) SubFlashNews() {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFlashNews(service)
	if err != nil {
		log.Println(err)
	}

	h.Firstpush(CATEGORY_FLASHNEWS, service, service.GetCode(), content)
}

func (h *UssdHandler) SubPronostic() {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentPronostic(SMS_PRONOSTIC_SAFE_SUB, service)
	if err != nil {
		log.Println(err)
	}

	h.Firstpush(CATEGORY_PRONOSTIC, service, service.GetCode(), content)
}

func (h *UssdHandler) SubAlerteCompetition(league *entity.League) {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowCompetition(service, league)
	if err != nil {
		log.Println(err)
	}

	h.Firstpush(CATEGORY_SMSALERTE_COMPETITION, service, league.GetCode(), content)

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), league.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	if !h.subscriptionFollowLeagueService.IsSub(sub.GetId(), league.GetId()) {
		// insert follow league
		h.subscriptionFollowLeagueService.Save(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				LimitPerDay:    LIMIT_PER_DAY,
				IsActive:       true,
			},
		)
	} else {
		// update follow league
		h.subscriptionFollowLeagueService.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				LimitPerDay:    LIMIT_PER_DAY,
				IsActive:       true,
			},
		)
	}

}

func (h *UssdHandler) SubAlerteEquipe(team *entity.Team) {
	service, err := h.getServiceByCode(h.req.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowTeam(service, team)
	if err != nil {
		log.Println(err)
	}

	// firstpush
	h.Firstpush(CATEGORY_SMSALERTE_EQUIPE, service, team.GetCode(), content)

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), team.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	if !h.subscriptionFollowTeamService.IsSub(sub.GetId(), team.GetId()) {
		// insert follow team
		h.subscriptionFollowTeamService.Save(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				LimitPerDay:    LIMIT_PER_DAY,
				IsActive:       true,
			},
		)
	} else {
		// update follow team
		h.subscriptionFollowTeamService.Update(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				LimitPerDay:    LIMIT_PER_DAY,
				IsActive:       true,
			},
		)
	}
}

/**
** STOP ON USSD
**/
func (h *UssdHandler) StopNonSMSAlerte(category string) {

	if h.subscriptionService.IsActiveSubscriptionByNonSMSAlerte(category, h.req.GetMsisdn()) {

		trxId := utils.GenerateTrxId()

		sub, err := h.subscriptionService.GetActiveByNonSMSAlerte(category, h.req.GetMsisdn())
		if err != nil {
			log.Println(err.Error())
		}

		service, err := h.serviceService.GetById(sub.GetServiceId())
		if err != nil {
			log.Println(err.Error())
		}

		content, err := h.getContentService(SMS_STOP, service)
		if err != nil {
			log.Println(err.Error())
		}

		sub.SetLatestTrxId(trxId)

		// unsub not sms-alerte
		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     sub.GetServiceId(),
				Msisdn:        sub.GetMsisdn(),
				Code:          service.GetCode(),
				LatestTrxId:   sub.GetLatestTrxId(),
				LatestSubject: SUBJECT_UNSUB,
				LatestStatus:  STATUS_SUCCESS,
				LatestKeyword: SMS_STOP + " " + service.GetCode(),
				UnsubAt:       time.Now(),
				IpAddress:     "",
			},
		)

		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:  sub.GetServiceId(),
				Msisdn:     sub.GetMsisdn(),
				Code:       service.GetCode(),
				TotalUnsub: sub.TotalUnsub + 1,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        sub.GetLatestTrxId(),
				ServiceID:    sub.GetServiceId(),
				Msisdn:       h.req.GetMsisdn(),
				Code:         service.GetCode(),
				Keyword:      SMS_STOP + " " + service.GetCode(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_UNSUB,
				IpAddress:    "",
				Payload:      "",
			},
		)

		h.historyService.Save(
			&entity.History{
				SubscriptionID: sub.GetId(),
				ServiceID:      sub.GetServiceId(),
				Msisdn:         sub.GetMsisdn(),
				Code:           SMS_STOP + " " + service.GetCode(),
				Keyword:        service.GetCode(),
				Subject:        SUBJECT_UNSUB,
				Status:         STATUS_SUCCESS,
				IpAddress:      "",
			},
		)

		s := &entity.Subscription{
			ServiceID: sub.GetServiceId(),
			Msisdn:    sub.GetMsisdn(),
			Code:      sub.GetCode(),
		}

		// set false is_active
		h.subscriptionService.UpdateNotActive(s)
		// set false is_retry
		h.subscriptionService.UpdateNotRetry(s)

		mt := &model.MTRequest{
			Smsc:         service.ScSubMT,
			Keyword:      SMS_STOP + " " + service.GetCode(),
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

}

func (h *UssdHandler) StopAlerteCompetition(league *entity.League) {

	if h.subscriptionService.IsActiveSubscriptionByCategory(CATEGORY_SMSALERTE_COMPETITION, h.req.GetMsisdn(), league.GetCode()) {

		trxId := utils.GenerateTrxId()

		sub, err := h.subscriptionService.GetActiveByCategory(CATEGORY_SMSALERTE_COMPETITION, h.req.GetMsisdn(), league.GetCode())
		if err != nil {
			log.Println(err.Error())
		}

		// unfollow league
		if h.subscriptionFollowLeagueService.IsSub(sub.GetId(), league.GetId()) {
			h.subscriptionFollowLeagueService.Update(
				&entity.SubscriptionFollowLeague{
					SubscriptionID: sub.GetId(),
					LeagueID:       league.GetId(),
				},
			)

			// disable
			h.subscriptionFollowLeagueService.Disable(
				&entity.SubscriptionFollowLeague{
					SubscriptionID: sub.GetId(),
					LeagueID:       league.GetId(),
				},
			)

			service, err := h.serviceService.GetById(sub.GetServiceId())
			if err != nil {
				log.Println(err.Error())
			}

			content, err := h.getContentUnFollowCompetition(service, league)
			if err != nil {
				log.Println(err.Error())
			}

			sub.SetLatestTrxId(trxId)

			// unsub sms-alerte
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     sub.GetServiceId(),
					Msisdn:        sub.GetMsisdn(),
					Code:          sub.GetCode(),
					LatestTrxId:   sub.GetLatestTrxId(),
					LatestSubject: SUBJECT_UNSUB,
					LatestStatus:  STATUS_SUCCESS,
					LatestKeyword: SMS_STOP + " " + league.GetCode(),
					UnsubAt:       time.Now(),
					IpAddress:     "-",
					TotalUnsub:    sub.TotalUnsub + 1,
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        sub.GetLatestTrxId(),
					ServiceID:    sub.GetServiceId(),
					Msisdn:       h.req.GetMsisdn(),
					Code:         service.GetCode(),
					Keyword:      SMS_STOP + " " + league.GetCode(),
					Status:       STATUS_SUCCESS,
					StatusCode:   "",
					StatusDetail: "",
					Subject:      SUBJECT_UNSUB,
					IpAddress:    "",
					Payload:      "",
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      sub.GetServiceId(),
					Msisdn:         sub.GetMsisdn(),
					Code:           sub.GetCode(),
					Keyword:        SMS_STOP + " " + league.GetCode(),
					Subject:        SUBJECT_UNSUB,
					Status:         STATUS_SUCCESS,
					IpAddress:      "-",
				},
			)

			s := &entity.Subscription{
				ServiceID: sub.GetServiceId(),
				Msisdn:    sub.GetMsisdn(),
				Code:      sub.GetCode(),
			}

			// set false is_active
			h.subscriptionService.UpdateNotActive(s)
			// set false is_retry
			h.subscriptionService.UpdateNotRetry(s)

			mt := &model.MTRequest{
				Smsc:         service.ScUnsubMT,
				Keyword:      SMS_STOP + " " + league.GetCode(),
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
	}

}

func (h *UssdHandler) StopAlerteEquipe(team *entity.Team) {

	if h.subscriptionService.IsActiveSubscriptionByCategory(CATEGORY_SMSALERTE_EQUIPE, h.req.GetMsisdn(), team.GetCode()) {
		trxId := utils.GenerateTrxId()

		sub, err := h.subscriptionService.GetActiveByCategory(CATEGORY_SMSALERTE_EQUIPE, h.req.GetMsisdn(), team.GetCode())
		if err != nil {
			log.Println(err.Error())
		}

		// unfollow team
		if h.subscriptionFollowTeamService.IsSub(sub.GetId(), team.GetId()) {
			h.subscriptionFollowTeamService.Update(
				&entity.SubscriptionFollowTeam{
					SubscriptionID: sub.GetId(),
					TeamID:         team.GetId(),
					LatestKeyword:  SMS_STOP,
				},
			)

			h.subscriptionFollowTeamService.Disable(
				&entity.SubscriptionFollowTeam{
					SubscriptionID: sub.GetId(),
					TeamID:         team.GetId(),
				},
			)

			service, err := h.serviceService.GetById(sub.GetServiceId())
			if err != nil {
				log.Println(err.Error())
			}

			content, err := h.getContentUnFollowTeam(service, team)
			if err != nil {
				log.Println(err.Error())
			}
			// 123
			sub.SetLatestTrxId(trxId)

			// unsub sms-alerte
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     sub.GetServiceId(),
					Msisdn:        sub.GetMsisdn(),
					Code:          sub.GetCode(),
					LatestTrxId:   sub.GetLatestTrxId(),
					LatestSubject: SUBJECT_UNSUB,
					LatestStatus:  STATUS_SUCCESS,
					LatestKeyword: SMS_STOP + " " + team.GetCode(),
					LatestNote:    "",
					UnsubAt:       time.Now(),
					TotalUnsub:    sub.TotalUnsub + 1,
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        sub.GetLatestTrxId(),
					ServiceID:    sub.GetServiceId(),
					Msisdn:       h.req.GetMsisdn(),
					Code:         service.GetCode(),
					Keyword:      SMS_STOP + " " + team.GetCode(),
					Status:       STATUS_SUCCESS,
					StatusCode:   "",
					StatusDetail: "",
					Subject:      SUBJECT_UNSUB,
					IpAddress:    "",
					Payload:      "",
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      sub.GetServiceId(),
					Code:           sub.GetCode(),
					Msisdn:         sub.GetMsisdn(),
					Keyword:        SMS_STOP + " " + team.GetCode(),
					Subject:        SUBJECT_UNSUB,
					Status:         STATUS_SUCCESS,
				},
			)

			s := &entity.Subscription{
				ServiceID: sub.GetServiceId(),
				Msisdn:    sub.GetMsisdn(),
				Code:      sub.GetCode(),
			}

			// set false is_active
			h.subscriptionService.UpdateNotActive(s)
			// set false is_retry
			h.subscriptionService.UpdateNotRetry(s)

			mt := &model.MTRequest{
				Smsc:         service.ScUnsubMT,
				Keyword:      SMS_STOP + " " + team.GetCode(),
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
	}

}

func (h *UssdHandler) Firstpush(category string, service *entity.Service, code string, content *entity.Content) {
	trxId := utils.GenerateTrxId()

	var note = ""

	subscription := &entity.Subscription{
		ServiceID:     service.GetId(),
		Category:      service.GetCategory(),
		Msisdn:        h.req.GetMsisdn(),
		Code:          code,
		Channel:       CHANNEL_USSD,
		LatestTrxId:   trxId,
		LatestKeyword: code,
		LatestSubject: SUBJECT_FIRSTPUSH,
		IsActive:      true,
		IpAddress:     "",
	}

	if h.IsSub(service, code) {
		h.subscriptionService.Update(subscription)
	} else {
		h.subscriptionService.Save(subscription)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), code)
	if err != nil {
		log.Println(err.Error())
	}

	if !h.subscriptionService.IsAfter24Hour(service.GetId(), h.req.GetMsisdn(), code) {
		// free 1 day
		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID:     service.GetId(),
				Msisdn:        h.req.GetMsisdn(),
				Code:          code,
				LatestTrxId:   trxId,
				LatestSubject: SUBJECT_FREEPUSH,
				LatestStatus:  STATUS_SUCCESS,
				RenewalAt:     time.Now().AddDate(0, 0, service.GetFreeDay()),
				FreeAt:        time.Now(),
				LatestPayload: "-",
				LatestNote:    note,
				IsFree:        true,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.req.GetMsisdn(),
				Code:         code,
				Channel:      CHANNEL_USSD,
				Keyword:      code,
				Amount:       0,
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_FREEPUSH,
				Payload:      "-",
				Note:         note,
			},
		)

		h.historyService.Save(
			&entity.History{
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Code:           code,
				Channel:        CHANNEL_USSD,
				Keyword:        code,
				Subject:        SUBJECT_FREEPUSH,
				Status:         STATUS_SUCCESS,
			},
		)

	} else {
		// charging if free day >= 1
		t := telco.NewTelco(h.logger, service, subscription, trxId)

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
						Msisdn:               h.req.GetMsisdn(),
						Code:                 code,
						LatestTrxId:          trxId,
						LatestSubject:        SUBJECT_FIRSTPUSH,
						LatestStatus:         STATUS_SUCCESS,
						TotalAmount:          sub.TotalAmount + service.GetPrice(),
						RenewalAt:            time.Now().AddDate(0, 0, service.GetRenewalDay()),
						ChargeAt:             time.Now(),
						TotalSuccess:         sub.TotalSuccess + 1,
						IsRetry:              false,
						TotalFirstpush:       sub.TotalFirstpush + 1,
						TotalAmountFirstpush: sub.TotalAmountFirstpush + service.GetPrice(),
						LatestPayload:        string(resp),
						LatestNote:           note,
					},
				)

				// is_retry set to false
				h.subscriptionService.UpdateNotRetry(sub)

				h.transactionService.Save(
					&entity.Transaction{
						TrxId:        trxId,
						ServiceID:    service.GetId(),
						Msisdn:       h.req.GetMsisdn(),
						Code:         code,
						Channel:      CHANNEL_USSD,
						Keyword:      code,
						Amount:       service.GetPrice(),
						Status:       STATUS_SUCCESS,
						StatusCode:   "",
						StatusDetail: "",
						Subject:      SUBJECT_FIRSTPUSH,
						Payload:      string(resp),
						Note:         note,
					},
				)

				h.historyService.Save(
					&entity.History{
						SubscriptionID: sub.GetId(),
						ServiceID:      service.GetId(),
						Msisdn:         h.req.GetMsisdn(),
						Code:           code,
						Channel:        CHANNEL_USSD,
						Keyword:        code,
						Subject:        SUBJECT_FIRSTPUSH,
						Status:         STATUS_SUCCESS,
					},
				)

				content, err = h.getContentService(SMS_SUCCESS_CHARGING, service)
				if err != nil {
					log.Println(err.Error())
				}

				// send notif success charge
				mt := &model.MTRequest{
					Smsc:         service.ScSubMT,
					Service:      service,
					Keyword:      code,
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
						Msisdn:        h.req.GetMsisdn(),
						Code:          code,
						LatestTrxId:   trxId,
						LatestSubject: SUBJECT_FIRSTPUSH,
						LatestStatus:  STATUS_FAILED,
						RenewalAt:     time.Now().AddDate(0, 0, 1),
						RetryAt:       time.Now(),
						TotalFailed:   sub.TotalFailed + 1,
						IsRetry:       true,
						LatestPayload: string(resp),
						LatestNote:    note,
					},
				)

				h.transactionService.Save(
					&entity.Transaction{
						TrxId:        trxId,
						ServiceID:    service.GetId(),
						Msisdn:       h.req.GetMsisdn(),
						Code:         code,
						Channel:      CHANNEL_USSD,
						Keyword:      code,
						Status:       STATUS_FAILED,
						StatusCode:   respDeduct.GetFaultCode(),
						StatusDetail: respDeduct.GetFaultString(),
						Subject:      SUBJECT_FIRSTPUSH,
						Payload:      string(resp),
						Note:         note,
					},
				)

				h.historyService.Save(
					&entity.History{
						SubscriptionID: sub.GetId(),
						ServiceID:      service.GetId(),
						Msisdn:         h.req.GetMsisdn(),
						Code:           code,
						Channel:        CHANNEL_USSD,
						Keyword:        code,
						Subject:        SUBJECT_FIRSTPUSH,
						Status:         STATUS_FAILED,
					},
				)
			}
		} else {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     service.GetId(),
					Msisdn:        h.req.GetMsisdn(),
					Code:          code,
					LatestTrxId:   trxId,
					LatestSubject: SUBJECT_FIRSTPUSH,
					LatestStatus:  STATUS_FAILED,
					RenewalAt:     time.Now().AddDate(0, 0, 1),
					RetryAt:       time.Now(),
					TotalFailed:   sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(respBal),
					LatestNote:    note,
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.req.GetMsisdn(),
					Code:         code,
					Channel:      CHANNEL_USSD,
					Keyword:      code,
					Status:       STATUS_FAILED,
					StatusCode:   "",
					StatusDetail: "INSUFF BALANCE",
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(respBal),
					Note:         note,
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      service.GetId(),
					Msisdn:         h.req.GetMsisdn(),
					Code:           code,
					Channel:        CHANNEL_USSD,
					Keyword:        code,
					Subject:        SUBJECT_FIRSTPUSH,
					Status:         STATUS_FAILED,
				},
			)
		}

	}

	// count total sub
	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID: service.GetId(),
			Msisdn:    h.req.GetMsisdn(),
			Code:      code,
			TotalSub:  sub.TotalSub + 1,
		},
	)

	mt := &model.MTRequest{
		Smsc:         service.ScSubMT,
		Keyword:      code,
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

func (h *UssdHandler) IsActiveSubByCategory(v, code string) bool {
	return h.subscriptionService.IsActiveSubscriptionByCategory(v, h.req.GetMsisdn(), code)
}

func (h *UssdHandler) IsActiveSubByNonSMSAlerte(v string) bool {
	return h.subscriptionService.IsActiveSubscriptionByNonSMSAlerte(v, h.req.GetMsisdn())
}

func (h *UssdHandler) IsSub(service *entity.Service, code string) bool {
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn(), code)
}

func (h *UssdHandler) getServiceByCode(code string) (*entity.Service, error) {
	return h.serviceService.Get(code)
}

func (h *UssdHandler) getContentService(v string, service *entity.Service) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetService(v, service)
}

func (h *UssdHandler) getContentLiveMatch(service *entity.Service) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(SMS_LIVE_MATCH_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetLiveMatch(SMS_LIVE_MATCH_SUB, service)
}

func (h *UssdHandler) getContentFlashNews(service *entity.Service) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(SMS_FLASH_NEWS_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFlashNews(SMS_FLASH_NEWS_SUB, service)
}

func (h *UssdHandler) getContentFollowCompetition(service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_COMPETITION_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowCompetition(SMS_FOLLOW_COMPETITION_SUB, service, league)
}

func (h *UssdHandler) getContentFollowTeam(service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_TEAM_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowTeam(SMS_FOLLOW_TEAM_SUB, service, team)
}

func (h *UssdHandler) getContentPronostic(v string, service *entity.Service) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetPronostic(v, service)
}

func (h *UssdHandler) getContentUnFollowCompetition(service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_COMPETITION_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowCompetition(SMS_FOLLOW_COMPETITION_STOP, service, league)
}

func (h *UssdHandler) getContentUnFollowTeam(service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_TEAM_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowTeam(SMS_FOLLOW_TEAM_STOP, service, team)
}
