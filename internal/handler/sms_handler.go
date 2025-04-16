package handler

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"strings"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/telco"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/wiliehidayat87/rmqp"
)

type SMSHandler struct {
	rmq                             rmqp.AMQP
	rds                             *redis.Client
	logger                          *logger.Logger
	serviceService                  services.IServiceService
	contentService                  services.IContentService
	subscriptionService             services.ISubscriptionService
	transactionService              services.ITransactionService
	historyService                  services.IHistoryService
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
	subscriptionCreditGoalService   services.ISubscriptionCreditGoalService
	subscriptionPredictWinService   services.ISubscriptionPredictWinService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	verifyService                   services.IVerifyService
	req                             *model.MORequest
}

func NewSMSHandler(
	rmq rmqp.AMQP,
	rds *redis.Client,
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	subscriptionCreditGoalService services.ISubscriptionCreditGoalService,
	subscriptionPredictWinService services.ISubscriptionPredictWinService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	verifyService services.IVerifyService,
	req *model.MORequest,
) *SMSHandler {
	return &SMSHandler{
		rmq:                             rmq,
		rds:                             rds,
		logger:                          logger,
		serviceService:                  serviceService,
		contentService:                  contentService,
		subscriptionService:             subscriptionService,
		transactionService:              transactionService,
		historyService:                  historyService,
		leagueService:                   leagueService,
		teamService:                     teamService,
		subscriptionCreditGoalService:   subscriptionCreditGoalService,
		subscriptionPredictWinService:   subscriptionPredictWinService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		verifyService:                   verifyService,
		req:                             req,
	}
}

const (
	CATEGORY_LIVEMATCH             string = "LIVEMATCH"
	CATEGORY_FLASHNEWS             string = "FLASHNEWS"
	CATEGORY_SMSALERTE_COMPETITION string = "SMSALERTE_COMPETITION"
	CATEGORY_SMSALERTE_EQUIPE      string = "SMSALERTE_EQUIPE"
	CATEGORY_PRONOSTIC             string = "PRONOSTIC"
	CATEGORY_PRONOSTIC_COMBINED    string = "PRONOSTIC_COMBINED"
	CATEGORY_PRONOSTIC_VIP         string = "PRONOSTIC_VIP"
	CATEGORY_CREDIT_GOAL           string = "CREDIT_GOAL"
	CATEGORY_PREDICT_WIN           string = "PREDICT_WIN"
)

const (
	SMS_CREDIT_GOAL_SUB                  string = "CREDIT_GOAL_SUB"
	SMS_CREDIT_GOAL_ALREADY_SUB          string = "CREDIT_GOAL_ALREADY_SUB"
	SMS_CREDIT_GOAL_UNVALID_SUB          string = "CREDIT_GOAL_UNVALID_SUB"
	SMS_CREDIT_GOAL_MATCH_END_PAYOUT     string = "CREDIT_GOAL_MATCH_END_PAYOUT"
	SMS_CREDIT_GOAL_MATCH_INCENTIVE      string = "CREDIT_GOAL_MATCH_INCENTIVE"
	SMS_PREDICT_SUB                      string = "PREDICT_SUB"
	SMS_PREDICT_SUB_BET_WIN              string = "PREDICT_SUB_BET_WIN"
	SMS_PREDICT_SUB_BET_DRAW             string = "PREDICT_SUB_BET_DRAW"
	SMS_PREDICT_UNVALID_SUB              string = "PREDICT_UNVALID_SUB"
	SMS_PREDICT_SUB_REJECT_MATCH_END     string = "PREDICT_SUB_REJECT_MATCH_END"
	SMS_PREDICT_SUB_REJECT_MATCH_STARTED string = "PREDICT_SUB_REJECT_MATCH_STARTED"
	SMS_PREDICT_MATCH_END_WINNER_AIRTIME string = "PREDICT_MATCH_END_WINNER_AIRTIME"
	SMS_PREDICT_MATCH_END_WINNER_LOTERY  string = "PREDICT_MATCH_END_WINNER_LOTERY"
	SMS_PREDICT_MATCH_END_LUCKY_LOSER    string = "PREDICT_MATCH_END_LUCKY_LOSER"
	SMS_PREDICT_MATCH_END_LOSER_NOTIF    string = "PREDICT_MATCH_END_LOSER_NOTIF"
	SMS_FOLLOW_TEAM_SUB                  string = "FOLLOW_TEAM_SUB"
	SMS_FOLLOW_TEAM_ALREADY_SUB          string = "FOLLOW_TEAM_ALREADY_SUB"
	SMS_FOLLOW_TEAM_UNVALID_SUB          string = "FOLLOW_TEAM_UNVALID_SUB"
	SMS_FOLLOW_TEAM_EXPIRE_SUB           string = "FOLLOW_TEAM_EXPIRE_SUB"
	SMS_FOLLOW_TEAM_STOP                 string = "FOLLOW_TEAM_STOP"
	SMS_FOLLOW_COMPETITION_SUB           string = "FOLLOW_COMPETITION_SUB"
	SMS_FOLLOW_COMPETITION_ALREADY_SUB   string = "FOLLOW_COMPETITION_ALREADY_SUB"
	SMS_FOLLOW_COMPETITION_UNVALID_SUB   string = "FOLLOW_COMPETITION_UNVALID_SUB"
	SMS_FOLLOW_COMPETITION_EXPIRE_SUB    string = "FOLLOW_COMPETITION_EXPIRE_SUB"
	SMS_FOLLOW_COMPETITION_STOP          string = "FOLLOW_COMPETITION_STOP"
	SMS_FOLLOW_UNVALID_SUB               string = "FOLLOW_UNVALID_SUB"
	SMS_PRONOSTIC_SAFE_SUB               string = "PRONOSTIC_SAFE_SUB"
	SMS_PRONOSTIC_SAFE_ALREADY_SUB       string = "PRONOSTIC_SAFE_ALREADY_SUB"
	SMS_PRONOSTIC_COMBINED_SUB           string = "PRONOSTIC_COMBINED_SUB"
	SMS_PRONOSTIC_COMBINED_ALREADY_SUB   string = "PRONOSTIC_COMBINED_ALREADY_SUB"
	SMS_PRONOSTIC_VIP_SUB                string = "PRONOSTIC_VIP_SUB"
	SMS_PRONOSTIC_VIP_ALREADY_SUB        string = "PRONOSTIC_VIP_ALREADY_SUB"
	SMS_PRONOSTIC_STOP                   string = "PRONOSTIC_STOP"
	SMS_PRONOSTIC_UNVALID_SUB            string = "PRONOSTIC_UNVALID_SUB"
	SMS_LIVE_MATCH_UNVALID_SUB           string = "LIVEMATCH_UNVALID_SUB"
	SMS_CONFIRMATION                     string = "CONFIRMATION"
	SMS_SUCCESS_CHARGING                 string = "SUCCESS_CHARGING"
	SMS_SUCCESS_CHARGING_SMS_ALERTE      string = "SUCCESS_CHARGING_SMS_ALERTE"
	SMS_REMINDER_48H_BEFORE_CHARGING     string = "REMINDER_48H_BEFORE_CHARGING"
	SMS_REMINDER_AFTER_TRIAL_ENDS        string = "REMINDER_AFTER_TRIAL_ENDS"
	SMS_INFO                             string = "INFO"
	SMS_STOP                             string = "STOP"
)

const (
	LIMIT_PER_DAY int = 4
)

func (h *SMSHandler) Registration() {
	l := h.logger.Init("sms", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("SMS")

	if h.req.HasLive() {
		// Livematch or Foot
		if !h.IsActiveSubByCategory(CATEGORY_LIVEMATCH, "") {
			if h.req.IsLiveDaily() {
				h.SubLivematch("jour")
			} else if h.req.IsLiveWeekly() {
				h.SubLivematch("semaine")
			} else if h.req.IsLiveMonthly() {
				h.SubLivematch("mois")
			} else {
				h.UnvalidSub(SMS_LIVE_MATCH_UNVALID_SUB)
			}
		} else {
			h.AlreadySubLiveMatch()
		}
	} else if h.req.IsFlashNews() {
		if !h.IsActiveSubByCategory(CATEGORY_FLASHNEWS, "") {
			h.SubFlashNews()
		} else {
			// h.AlreadySubLiveMatch()
		}
	} else if h.req.HasProno() {
		// Pronostic Safe Sub
		if !h.IsActiveSubByCategory(CATEGORY_PRONOSTIC, "") {
			// pronostic general
			if h.req.IsPronoDaily() {
				h.SubProno("jour")
			} else if h.req.IsPronoWeekly() {
				h.SubProno("semaine")
			} else if h.req.IsPronoMonthly() {
				h.SubProno("mois")
			} else {
				h.UnvalidSub(SMS_PRONOSTIC_UNVALID_SUB)
			}
		} else {
			h.AlreadySubSafe()
		}
	} else if h.leagueService.IsLeagueByName(h.req.GetSMS()) {
		league, err := h.leagueService.GetByName(h.req.GetSMS())
		if err != nil {
			log.Println(err.Error())
		}
		if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE_COMPETITION, league.GetCode()) {
			h.SubAlerteCompetition(league)
		} else {
			// SMS-Alerte Competition
			h.AlreadySubAlerteCompetition(league)
		}
	} else if h.teamService.IsTeamByName(h.req.GetSMS()) {
		team, err := h.teamService.GetByName(h.req.GetSMS())
		if err != nil {
			log.Println(err.Error())
		}
		if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE_EQUIPE, team.GetCode()) {
			h.SubAlerteEquipe(team)
		} else {
			// SMS-Alerte Equipe
			h.AlreadySubAlerteEquipe(team)
		}
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsTicket() {
		// Pronostic Combined Sub
		if !h.IsActiveSubByCategory(CATEGORY_PRONOSTIC_COMBINED, "") {
			h.SubCombined()
		} else {
			h.AlreadySubCombined()
		}
	} else if h.req.IsVIP() {
		// Pronostic VIP Sub
		if !h.IsActiveSubByCategory(CATEGORY_PRONOSTIC_VIP, "") {
			h.SubVIP()
		} else {
			h.AlreadySubVIP()
		}
	} else if h.req.IsStop() {
		// Stop alive ussd
		if h.req.IsStopAlive() {
			if h.IsActiveSubByNonSMSAlerte(CATEGORY_LIVEMATCH) {
				h.Unsub(CATEGORY_LIVEMATCH)
			}
		}

		// Stop flashnews ussd
		if h.req.IsStopFlashNews() {
			if h.IsActiveSubByNonSMSAlerte(CATEGORY_FLASHNEWS) {
				h.Unsub(CATEGORY_FLASHNEWS)
			}
		}

		// Stop prono (safe)
		if h.req.IsStopProno() {
			if h.IsActiveSubByNonSMSAlerte(CATEGORY_PRONOSTIC) {
				h.StopPronostic(CATEGORY_PRONOSTIC)
			}

		}

		// Stop ticket (combined)
		if h.req.IsStopTicket() {
			if h.IsActiveSubByNonSMSAlerte(CATEGORY_PRONOSTIC_COMBINED) {
				h.StopPronostic(CATEGORY_PRONOSTIC_COMBINED)
			}
		}

		// Stop VIP
		if h.req.IsStopVIP() {
			if h.IsActiveSubByNonSMSAlerte(CATEGORY_PRONOSTIC_VIP) {
				h.StopPronostic(CATEGORY_PRONOSTIC_VIP)
			}
		}

		if h.leagueService.IsLeagueByName(h.req.GetStopKeyword()) {
			// Stop SMS Alerte Competition
			league, err := h.leagueService.GetByName(h.req.GetStopKeyword())
			if err != nil {
				log.Println(err.Error())
			}

			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE_COMPETITION, league.GetCode()) {
				h.StopAlerteCompetition(CATEGORY_SMSALERTE_COMPETITION, league)
			}
		}

		if h.teamService.IsTeamByName(h.req.GetStopKeyword()) {
			// Stop SMS Alerte Equipe
			team, err := h.teamService.GetByName(h.req.GetStopKeyword())
			if err != nil {
				log.Println(err.Error())
			}

			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE_EQUIPE, team.GetCode()) {
				h.StopAlerteEquipe(CATEGORY_SMSALERTE_EQUIPE, team)
			}
		}

		if h.serviceService.IsService(h.req.GetStopKeyword()) {
			service, _ := h.serviceService.Get(h.req.GetStopKeyword())
			if h.IsActiveSubByNonSMSAlerte(service.GetCategory()) {
				h.Unsub(service.GetCategory())
			}
		}

	} else {
		h.Unvalid(SMS_FOLLOW_UNVALID_SUB)
	}
}

func (h *SMSHandler) Firstpush(category string, service *entity.Service, code string, content *entity.Content) {
	trxId := utils.GenerateTrxId()

	var note = ""
	if category == CATEGORY_SMSALERTE_COMPETITION {
		if h.leagueService.IsLeagueByCode(code) {
			league, _ := h.leagueService.GetByCode(code)
			note = strings.ToUpper("FOLLOW " + league.GetName())
		}
	}

	if category == CATEGORY_SMSALERTE_EQUIPE {
		if h.teamService.IsTeamByCode(code) {
			team, _ := h.teamService.GetByCode(code)
			note = strings.ToUpper("FOLLOW " + team.GetName())
		}
	}

	subscription := &entity.Subscription{
		ServiceID:     service.GetId(),
		Category:      service.GetCategory(),
		Msisdn:        h.req.GetMsisdn(),
		Code:          code,
		Channel:       CHANNEL_SMS,
		LatestTrxId:   trxId,
		LatestKeyword: h.req.GetSMS(),
		LatestSubject: SUBJECT_FIRSTPUSH,
		IsActive:      true,
		IsFollowTeam:  true,
		IpAddress:     h.req.GetIpAddress(),
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
				Channel:      CHANNEL_SMS,
				Keyword:      h.req.GetSMS(),
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
				Channel:        CHANNEL_SMS,
				Keyword:        h.req.GetSMS(),
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
						Channel:      CHANNEL_SMS,
						Keyword:      h.req.GetSMS(),
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
						Channel:        CHANNEL_SMS,
						Keyword:        h.req.GetSMS(),
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
					Keyword:      h.req.GetSMS(),
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
						Channel:      CHANNEL_SMS,
						Keyword:      h.req.GetSMS(),
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
						Channel:        CHANNEL_SMS,
						Keyword:        h.req.GetSMS(),
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
					Channel:      CHANNEL_SMS,
					Keyword:      h.req.GetSMS(),
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
					Channel:        CHANNEL_SMS,
					Keyword:        h.req.GetSMS(),
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
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) SubAlerteCompetition(league *entity.League) {
	service, err := h.getServiceSMSAlerteCompetitionDaily()
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

func (h *SMSHandler) SubAlerteEquipe(team *entity.Team) {
	service, err := h.getServiceSMSAlerteEquipeDaily()
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

func (h *SMSHandler) AlreadySubAlerteCompetition(league *entity.League) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteCompetitionDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowCompetition(service, league)
	if err != nil {
		log.Println(err)
	}

	if h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn(), league.GetCode()) {
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

		mt := &model.MTRequest{
			Smsc:         h.req.GetTo(),
			Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) AlreadySubAlerteEquipe(team *entity.Team) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteEquipeDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowTeam(service, team)
	if err != nil {
		log.Println(err)
	}

	if h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn(), team.GetCode()) {
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

		mt := &model.MTRequest{
			Smsc:         h.req.GetTo(),
			Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) SubLivematch(p string) {
	service, err := h.getServiceLiveMatch(p)
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentLiveMatch(service)
	if err != nil {
		log.Println(err)
	}

	// GET_CODE
	code := service.GetCode()

	h.Firstpush(CATEGORY_LIVEMATCH, service, code, content)
}

func (h *SMSHandler) SubFlashNews() {
	service, err := h.getServiceFlashNewsDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFlashNews(SMS_FLASH_NEWS_SUB, service)
	if err != nil {
		log.Println(err)
	}

	// GET_CODE
	code := service.GetCode()

	h.Firstpush(CATEGORY_FLASHNEWS, service, code, content)
}

func (h *SMSHandler) SubProno(p string) {
	service, err := h.getServicePronostic(p)
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentPronostic(SMS_PRONOSTIC_SAFE_SUB, service)
	if err != nil {
		log.Println(err)
	}

	// GET_CODE
	code := service.GetCode()

	h.Firstpush(CATEGORY_PRONOSTIC, service, code, content)
}

func (h *SMSHandler) SubCombined() {
	service, err := h.getServicePronosticCombinedDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentPronostic(SMS_PRONOSTIC_COMBINED_SUB, service)
	if err != nil {
		log.Println(err)
	}

	// GET_CODE
	code := service.GetCode()

	h.Firstpush(CATEGORY_PRONOSTIC_COMBINED, service, code, content)
}

func (h *SMSHandler) SubVIP() {
	service, err := h.getServicePronosticVIPDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentPronostic(SMS_PRONOSTIC_VIP_SUB, service)
	if err != nil {
		log.Println(err)
	}

	// GET_CODE
	code := service.GetCode()

	h.Firstpush(CATEGORY_PRONOSTIC_VIP, service, code, content)
}

func (h *SMSHandler) AlreadySubLiveMatch() {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceLiveMatch("")
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_LIVE_MATCH_ALREADY_SUB)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), service.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) AlreadySubSafe() {
	trxId := utils.GenerateTrxId()

	service, err := h.getServicePronosticSafeDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_PRONOSTIC_SAFE_ALREADY_SUB)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), service.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) AlreadySubCombined() {
	trxId := utils.GenerateTrxId()

	service, err := h.getServicePronosticCombinedDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_PRONOSTIC_COMBINED_ALREADY_SUB)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), service.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) AlreadySubVIP() {
	trxId := utils.GenerateTrxId()

	service, err := h.getServicePronosticVIPDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_PRONOSTIC_VIP_ALREADY_SUB)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn(), service.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) Info() {
	trxId := utils.GenerateTrxId()

	content, err := h.getContent(SMS_INFO)
	if err != nil {
		log.Println(err.Error())
	}
	mt := &model.MTRequest{
		Smsc:    h.req.GetTo(),
		Keyword: h.req.GetSMS(),
		Service: &entity.Service{
			UrlMT:  URL_MT,
			UserMT: USER_MT,
			PassMT: PASS_MT,
		},
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) Unvalid(v string) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteEquipeDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentSMSAlerteUnvalid(v, service)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:    h.req.GetTo(),
		Keyword: h.req.GetSMS(),
		Service: &entity.Service{
			UrlMT:  URL_MT,
			UserMT: USER_MT,
			PassMT: PASS_MT,
		},
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) UnvalidCompetition(v string) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteCompetitionDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentSMSAlerteUnvalid(v, service)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:    h.req.GetTo(),
		Keyword: h.req.GetSMS(),
		Service: &entity.Service{
			UrlMT:  URL_MT,
			UserMT: USER_MT,
			PassMT: PASS_MT,
		},
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) UnvalidEquipe(v string) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteEquipeDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentSMSAlerteUnvalid(v, service)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:    h.req.GetTo(),
		Keyword: h.req.GetSMS(),
		Service: &entity.Service{
			UrlMT:  URL_MT,
			UserMT: USER_MT,
			PassMT: PASS_MT,
		},
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) UnvalidSub(v string) {
	trxId := utils.GenerateTrxId()

	content, err := h.getContent(v)
	if err != nil {
		log.Println(err.Error())
	}
	mt := &model.MTRequest{
		Smsc:    h.req.GetTo(),
		Keyword: h.req.GetSMS(),
		Service: &entity.Service{
			UrlMT:  URL_MT,
			UserMT: USER_MT,
			PassMT: PASS_MT,
		},
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) IsActiveSubByCategory(v, code string) bool {
	return h.subscriptionService.IsActiveSubscriptionByCategory(v, h.req.GetMsisdn(), code)
}

func (h *SMSHandler) IsActiveSubByNonSMSAlerte(v string) bool {
	return h.subscriptionService.IsActiveSubscriptionByNonSMSAlerte(v, h.req.GetMsisdn())
}

func (h *SMSHandler) StopAlerteCompetition(category string, league *entity.League) {

	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetByCategory(category, h.req.GetMsisdn(), league.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	// unfollow league
	if h.subscriptionFollowLeagueService.IsSub(sub.GetId(), league.GetId()) {
		h.subscriptionFollowLeagueService.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				LatestKeyword:  h.req.GetSMS(),
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
				LatestKeyword: h.req.GetSMS(),
				UnsubAt:       time.Now(),
				IpAddress:     h.req.GetIpAddress(),
				TotalUnsub:    sub.TotalUnsub + 1,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        sub.GetLatestTrxId(),
				ServiceID:    sub.GetServiceId(),
				Msisdn:       h.req.GetMsisdn(),
				Code:         service.GetCode(),
				Keyword:      h.req.GetSMS(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_UNSUB,
				IpAddress:    h.req.GetIpAddress(),
				Payload:      "",
			},
		)

		h.historyService.Save(
			&entity.History{
				SubscriptionID: sub.GetId(),
				ServiceID:      sub.GetServiceId(),
				Msisdn:         sub.GetMsisdn(),
				Code:           sub.GetCode(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_UNSUB,
				Status:         STATUS_SUCCESS,
				IpAddress:      h.req.GetIpAddress(),
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
			Smsc:         h.req.GetTo(),
			Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) StopAlerteEquipe(category string, team *entity.Team) {

	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetByCategory(category, h.req.GetMsisdn(), team.GetCode())
	if err != nil {
		log.Println(err.Error())
	}

	// unfollow team
	if h.subscriptionFollowTeamService.IsSub(sub.GetId(), team.GetId()) {
		h.subscriptionFollowTeamService.Update(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				LatestKeyword:  h.req.GetSMS(),
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
				LatestKeyword: h.req.GetSMS(),
				LatestNote:    "",
				UnsubAt:       time.Now(),
				IpAddress:     h.req.GetIpAddress(),
				TotalUnsub:    sub.TotalUnsub + 1,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        sub.GetLatestTrxId(),
				ServiceID:    sub.GetServiceId(),
				Msisdn:       h.req.GetMsisdn(),
				Code:         service.GetCode(),
				Keyword:      h.req.GetSMS(),
				Status:       STATUS_SUCCESS,
				StatusCode:   "",
				StatusDetail: "",
				Subject:      SUBJECT_UNSUB,
				IpAddress:    h.req.GetIpAddress(),
				Payload:      "",
			},
		)

		h.historyService.Save(
			&entity.History{
				SubscriptionID: sub.GetId(),
				ServiceID:      sub.GetServiceId(),
				Code:           sub.GetCode(),
				Msisdn:         sub.GetMsisdn(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_UNSUB,
				Status:         STATUS_SUCCESS,
				IpAddress:      h.req.GetIpAddress(),
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
			Smsc:         h.req.GetTo(),
			Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) StopPronostic(category string) {
	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetByNonSMSAlerte(category, h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	service, err := h.serviceService.GetById(sub.GetServiceId())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentStopPronostic()
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
			LatestKeyword: h.req.GetSMS(),
			UnsubAt:       time.Now(),
			IpAddress:     h.req.GetIpAddress(),
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
			Keyword:      h.req.GetSMS(),
			Status:       STATUS_SUCCESS,
			StatusCode:   "",
			StatusDetail: "",
			Subject:      SUBJECT_UNSUB,
			IpAddress:    h.req.GetIpAddress(),
			Payload:      "",
		},
	)

	h.historyService.Save(
		&entity.History{
			SubscriptionID: sub.GetId(),
			ServiceID:      sub.GetServiceId(),
			Msisdn:         sub.GetMsisdn(),
			Code:           service.GetCode(),
			Keyword:        h.req.GetSMS(),
			Subject:        SUBJECT_UNSUB,
			Status:         STATUS_SUCCESS,
			IpAddress:      h.req.GetIpAddress(),
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
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) Unsub(category string) {
	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetByNonSMSAlerte(category, h.req.GetMsisdn())
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
			LatestKeyword: h.req.GetSMS(),
			UnsubAt:       time.Now(),
			IpAddress:     h.req.GetIpAddress(),
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
			Keyword:      h.req.GetSMS(),
			Status:       STATUS_SUCCESS,
			StatusCode:   "",
			StatusDetail: "",
			Subject:      SUBJECT_UNSUB,
			IpAddress:    h.req.GetIpAddress(),
			Payload:      "",
		},
	)

	h.historyService.Save(
		&entity.History{
			SubscriptionID: sub.GetId(),
			ServiceID:      sub.GetServiceId(),
			Msisdn:         sub.GetMsisdn(),
			Code:           service.GetCode(),
			Keyword:        h.req.GetSMS(),
			Subject:        SUBJECT_UNSUB,
			Status:         STATUS_SUCCESS,
			IpAddress:      h.req.GetIpAddress(),
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
		Smsc:         h.req.GetTo(),
		Keyword:      h.req.GetSMS(),
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

func (h *SMSHandler) IsSub(service *entity.Service, code string) bool {
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn(), code)
}

func (h *SMSHandler) getServiceSMSAlerteCompetitionDaily() (*entity.Service, error) {
	return h.serviceService.Get("SAC1")
}

func (h *SMSHandler) getServiceSMSAlerteEquipeDaily() (*entity.Service, error) {
	return h.serviceService.Get("SAE1")
}

func (h *SMSHandler) getServiceLiveMatch(p string) (*entity.Service, error) {
	return h.serviceService.GetByPackage("LIVEMATCH", p)
}

func (h *SMSHandler) getServicePronostic(p string) (*entity.Service, error) {
	return h.serviceService.GetByPackage("PRONOSTIC", p)
}

func (h *SMSHandler) getServiceFlashNewsDaily() (*entity.Service, error) {
	return h.serviceService.Get("FN1")
}

func (h *SMSHandler) getServicePronosticSafeDaily() (*entity.Service, error) {
	return h.serviceService.Get("PS1")
}

func (h *SMSHandler) getServicePronosticCombinedDaily() (*entity.Service, error) {
	return h.serviceService.Get("PC1")
}

func (h *SMSHandler) getServicePronosticVIPDaily() (*entity.Service, error) {
	return h.serviceService.Get("PV1")
}

func (h *SMSHandler) getContent(v string) (*entity.Content, error) {
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

func (h *SMSHandler) getContentService(v string, service *entity.Service) (*entity.Content, error) {
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

func (h *SMSHandler) getContentFollowCompetition(service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_COMPETITION_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowCompetition(SMS_FOLLOW_COMPETITION_SUB, service, league)
}

func (h *SMSHandler) getContentFollowTeam(service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_TEAM_SUB) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowTeam(SMS_FOLLOW_TEAM_SUB, service, team)
}

func (h *SMSHandler) getContentUnFollowCompetition(service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_COMPETITION_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowCompetition(SMS_FOLLOW_COMPETITION_STOP, service, league)
}

func (h *SMSHandler) getContentUnFollowTeam(service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_TEAM_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowTeam(SMS_FOLLOW_TEAM_STOP, service, team)
}

func (h *SMSHandler) getContentFlashNews(v string, service *entity.Service) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFlashNews(v, service)
}

func (h *SMSHandler) getContentPronostic(v string, service *entity.Service) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetPronostic(v, service)
}

func (h *SMSHandler) getContentStopPronostic() (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_PRONOSTIC_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(SMS_PRONOSTIC_STOP)
}

func (h *SMSHandler) getContentSMSAlerteUnvalid(v string, service *entity.Service) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetSMSAlerteUnvalid(v, service)
}

func (h *SMSHandler) getContentLiveMatch(service *entity.Service) (*entity.Content, error) {
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

/**
** UNUSED
**/
/**
func (h *SMSHandler) CreditGoal() {

	if h.leagueService.IsLeagueByName(h.req.GetSMS()) {
		if !h.IsActiveSubByCategory(CATEGORY_CREDIT_GOAL) {
			h.Confirmation(CATEGORY_CREDIT_GOAL)
		} else {
			// SMS-Alerte Competition
			h.AlerteCompetition()
		}
		// SMS-Alerte Matchs
	} else if h.teamService.IsTeamByName(h.req.GetSMS()) {
		if !h.IsActiveSubByCategory(CATEGORY_CREDIT_GOAL) {
			h.Confirmation(CATEGORY_CREDIT_GOAL)
		} else {
			// SMS-Alerte Equipe
			h.AlerteEquipe()
		}
		// SMS-Alerte Matchs
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		if h.IsActiveSubByCategory(CATEGORY_CREDIT_GOAL) {
			h.Stop()
		}
	}

func (h *SMSHandler) Prediction() {
	if h.leagueService.IsLeagueByName(h.req.GetSMS()) {
		if !h.IsActiveSubByCategory(CATEGORY_PREDICT) {
			h.Confirmation(CATEGORY_PREDICT)
		} else {
			// SMS-Alerte Competition
			h.AlerteCompetition()
			// SMS-Alerte Matchs
		}
	} else if h.teamService.IsTeamByName(h.req.GetSMS()) {
		if !h.IsActiveSubByCategory(CATEGORY_PREDICT) {
			h.Confirmation(CATEGORY_PREDICT)
		} else {
			// SMS-Alerte Equipe
			h.AlerteEquipe()
			// SMS-Alerte Matchs
		}
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		if h.IsActiveSubByCategory(CATEGORY_PREDICT) {
			h.Stop()
		}
	}
}
**/
