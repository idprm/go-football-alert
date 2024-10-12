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
	summaryService                  services.ISummaryService
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
	subscriptionCreditGoalService   services.ISubscriptionCreditGoalService
	subscriptionPredictService      services.ISubscriptionPredictService
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
	summaryService services.ISummaryService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	subscriptionCreditGoalService services.ISubscriptionCreditGoalService,
	subscriptionPredictService services.ISubscriptionPredictService,
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
		summaryService:                  summaryService,
		leagueService:                   leagueService,
		teamService:                     teamService,
		subscriptionCreditGoalService:   subscriptionCreditGoalService,
		subscriptionPredictService:      subscriptionPredictService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		verifyService:                   verifyService,
		req:                             req,
	}
}

const (
	CATEGORY_LIVEMATCH             string = "LIVEMATCH"
	CATEGORY_FLASHNEWS             string = "FLASHNEWS"
	CATEGORY_SMSALERTE             string = "SMSALERTE"
	CATEGORY_CREDIT_GOAL           string = "CREDITGOAL"
	CATEGORY_PREDICT               string = "PREDICTION"
	CATEGORY_PRONOSTIC             string = "PRONOSTIC"
	SUBCATEGORY_FOLLOW_COMPETITION string = "FOLLOW_COMPETITION"
	SUBCATEGORY_FOLLOW_TEAM        string = "FOLLOW_TEAM"
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
	SMS_FOLLOW_COMPETITION_SUB           string = "FOLLOW_COMPETITION_SUB"
	SMS_FOLLOW_COMPETITION_ALREADY_SUB   string = "FOLLOW_COMPETITION_ALREADY_SUB"
	SMS_FOLLOW_COMPETITION_UNVALID_SUB   string = "FOLLOW_COMPETITION_UNVALID_SUB"
	SMS_FOLLOW_COMPETITION_EXPIRE_SUB    string = "FOLLOW_COMPETITION_EXPIRE_SUB"
	SMS_FOLLOW_UNVALID_SUB               string = "FOLLOW_UNVALID_SUB"
	SMS_CONFIRMATION                     string = "CONFIRMATION"
	SMS_INFO                             string = "INFO"
	SMS_STOP                             string = "STOP"
)

func (h *SMSHandler) Registration() {
	l := h.logger.Init("sms", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("SMS")

	/**
	 ** SMS Alerte
	 **/
	h.SMSAlerte()
}

func (h *SMSHandler) SMSAlerte() {
	if h.leagueService.IsLeagueByName(h.req.GetSMS()) {
		league, err := h.leagueService.GetByName(h.req.GetSMS())
		if err != nil {
			log.Println(err.Error())
		}
		if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
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
		if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
			h.SubAlerteEquipe(team)
		} else {
			// SMS-Alerte Equipe
			h.AlreadySubAlerteEquipe(team)
		}
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		if h.req.IsStopAlerte() {
			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
				h.Stop(CATEGORY_SMSALERTE)
			}
		}

		if h.req.IsStopAlive() {
			if h.IsActiveSubByCategory(CATEGORY_LIVEMATCH) {
				h.Stop(CATEGORY_LIVEMATCH)
			}
		}

		if h.req.IsStopFlashNews() {
			if h.IsActiveSubByCategory(CATEGORY_FLASHNEWS) {
				h.Stop(CATEGORY_FLASHNEWS)
			}
		}

	} else {
		h.Unvalid(SMS_FOLLOW_UNVALID_SUB)
	}

}

func (h *SMSHandler) SubAlerteCompetition(league *entity.League) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowCompetition(SMS_FOLLOW_COMPETITION_SUB, service, league)
	if err != nil {
		log.Println(err)
	}

	summary := &entity.Summary{
		ServiceID: service.GetId(),
		CreatedAt: time.Now(),
	}

	subscription := &entity.Subscription{
		ServiceID:      service.GetId(),
		Category:       service.GetCategory(),
		Msisdn:         h.req.GetMsisdn(),
		LatestTrxId:    trxId,
		LatestKeyword:  h.req.GetSMS(),
		LatestSubject:  SUBJECT_FIRSTPUSH,
		IsActive:       true,
		IsFollowLeague: true,
		IpAddress:      h.req.GetIpAddress(),
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

	if !h.subscriptionFollowLeagueService.IsSub(sub.GetId()) {
		// insert follow league
		h.subscriptionFollowLeagueService.Save(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				IsActive:       true,
			},
		)
	} else {
		// update follow league
		h.subscriptionFollowLeagueService.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				IsActive:       true,
			},
		)
	}

	// disable if follow team
	if h.subscriptionFollowTeamService.IsSub(sub.GetId()) {
		h.subscriptionFollowTeamService.Disable(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
			},
		)
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
				Keyword:      h.req.GetSMS(),
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
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FREEPUSH,
				Status:         STATUS_SUCCESS,
			},
		)

	} else {

		// charging if free day >= 1
		t := telco.NewTelco(h.logger, service, subscription, trxId)

		deductFee, err := t.DeductFee()
		if err != nil {
			log.Println(err.Error())
		}

		var respDeduct *model.DeductResponse
		xml.Unmarshal(utils.EscapeChar(deductFee), &respDeduct)

		if respDeduct.IsFailed() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     service.GetId(),
					Msisdn:        h.req.GetMsisdn(),
					LatestTrxId:   trxId,
					LatestSubject: SUBJECT_FIRSTPUSH,
					LatestStatus:  STATUS_FAILED,
					RenewalAt:     time.Now().AddDate(0, 0, 1),
					RetryAt:       time.Now(),
					TotalFailed:   sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(deductFee),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.req.GetMsisdn(),
					Keyword:      h.req.GetSMS(),
					Status:       STATUS_FAILED,
					StatusCode:   respDeduct.GetFaultCode(),
					StatusDetail: respDeduct.GetFaultString(),
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(deductFee),
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      service.GetId(),
					Msisdn:         h.req.GetMsisdn(),
					Keyword:        h.req.GetSMS(),
					Subject:        SUBJECT_FIRSTPUSH,
					Status:         STATUS_FAILED,
				},
			)

			// setter summary
			summary.SetTotalChargeFailed(1)
		} else {
			h.subscriptionService.Update(
				&entity.Subscription{
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
					LatestPayload:        string(deductFee),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.req.GetMsisdn(),
					Keyword:      h.req.GetSMS(),
					Amount:       service.GetPrice(),
					Status:       STATUS_SUCCESS,
					StatusCode:   "",
					StatusDetail: "",
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(deductFee),
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      service.GetId(),
					Msisdn:         h.req.GetMsisdn(),
					Keyword:        h.req.GetSMS(),
					Subject:        SUBJECT_FIRSTPUSH,
					Status:         STATUS_SUCCESS,
				},
			)

			summary.SetTotalChargeSuccess(1)
			summary.SetTotalRevenue(service.GetPrice())
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

func (h *SMSHandler) SubAlerteEquipe(team *entity.Team) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowTeam(SMS_FOLLOW_TEAM_SUB, service, team)
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
		LatestKeyword: h.req.GetSMS(),
		LatestSubject: SUBJECT_FIRSTPUSH,
		IsActive:      true,
		IsFollowTeam:  true,
		IpAddress:     h.req.GetIpAddress(),
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	if !h.subscriptionFollowTeamService.IsSub(sub.GetId()) {
		// insert follow team
		h.subscriptionFollowTeamService.Save(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				IsActive:       true,
			},
		)
	} else {
		// update follow team
		h.subscriptionFollowTeamService.Update(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				IsActive:       true,
			},
		)
	}

	// disable if follow league
	if h.subscriptionFollowLeagueService.IsSub(sub.GetId()) {
		h.subscriptionFollowLeagueService.Disable(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
			},
		)
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
				Keyword:      h.req.GetSMS(),
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
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FREEPUSH,
				Status:         STATUS_SUCCESS,
			},
		)

	} else {

		// charging if free day >= 1
		t := telco.NewTelco(h.logger, service, subscription, trxId)

		deductFee, err := t.DeductFee()
		if err != nil {
			log.Println(err.Error())
		}

		var respDeduct *model.DeductResponse
		xml.Unmarshal(utils.EscapeChar(deductFee), &respDeduct)

		if respDeduct.IsFailed() {
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     service.GetId(),
					Msisdn:        h.req.GetMsisdn(),
					LatestTrxId:   trxId,
					LatestSubject: SUBJECT_FIRSTPUSH,
					LatestStatus:  STATUS_FAILED,
					RenewalAt:     time.Now().AddDate(0, 0, 1),
					RetryAt:       time.Now(),
					TotalFailed:   sub.TotalFailed + 1,
					IsRetry:       true,
					LatestPayload: string(deductFee),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.req.GetMsisdn(),
					Keyword:      h.req.GetSMS(),
					Status:       STATUS_FAILED,
					StatusCode:   respDeduct.GetFaultCode(),
					StatusDetail: respDeduct.GetFaultString(),
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(deductFee),
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      service.GetId(),
					Msisdn:         h.req.GetMsisdn(),
					Keyword:        h.req.GetSMS(),
					Subject:        SUBJECT_FIRSTPUSH,
					Status:         STATUS_FAILED,
				},
			)

			// setter summary
			summary.SetTotalChargeFailed(1)
		} else {
			h.subscriptionService.Update(
				&entity.Subscription{
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
					LatestPayload:        string(deductFee),
				},
			)

			h.transactionService.Save(
				&entity.Transaction{
					TrxId:        trxId,
					ServiceID:    service.GetId(),
					Msisdn:       h.req.GetMsisdn(),
					Keyword:      h.req.GetSMS(),
					Amount:       service.GetPrice(),
					Status:       STATUS_SUCCESS,
					StatusCode:   "",
					StatusDetail: "",
					Subject:      SUBJECT_FIRSTPUSH,
					Payload:      string(deductFee),
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      service.GetId(),
					Msisdn:         h.req.GetMsisdn(),
					Keyword:        h.req.GetSMS(),
					Subject:        SUBJECT_FIRSTPUSH,
					Status:         STATUS_SUCCESS,
				},
			)

			summary.SetTotalChargeSuccess(1)
			summary.SetTotalRevenue(service.GetPrice())
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

func (h *SMSHandler) AlreadySubAlerteCompetition(league *entity.League) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowCompetition(SMS_FOLLOW_COMPETITION_SUB, service, league)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	// update keyword in sub
	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID:     service.GetId(),
			Msisdn:        h.req.GetMsisdn(),
			LatestTrxId:   trxId,
			LatestKeyword: h.req.GetSMS(),
		},
	)

	// update in follow league
	if h.subscriptionFollowLeagueService.IsSub(sub.GetId()) {
		h.subscriptionFollowLeagueService.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
				LeagueID:       league.GetId(),
				IsActive:       true,
			},
		)
	}

	// disable if follow team
	if h.subscriptionFollowTeamService.IsSub(sub.GetId()) {
		h.subscriptionFollowTeamService.Disable(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
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

func (h *SMSHandler) AlreadySubAlerteEquipe(team *entity.Team) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContentFollowTeam(SMS_FOLLOW_TEAM_SUB, service, team)
	if err != nil {
		log.Println(err)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	// update keyword in sub
	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID:     service.GetId(),
			Msisdn:        h.req.GetMsisdn(),
			LatestTrxId:   trxId,
			LatestKeyword: h.req.GetSMS(),
		},
	)

	// update in follow team
	if h.subscriptionFollowTeamService.IsSub(sub.GetId()) {
		h.subscriptionFollowTeamService.Update(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: sub.GetId(),
				TeamID:         team.GetId(),
				IsActive:       true,
			},
		)
	}

	// disable if follow league
	if h.subscriptionFollowLeagueService.IsSub(sub.GetId()) {
		h.subscriptionFollowLeagueService.Disable(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: sub.GetId(),
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

func (h *SMSHandler) Stop(category string) {
	trxId := utils.GenerateTrxId()

	sub, err := h.subscriptionService.GetByCategory(category, h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	service, err := h.serviceService.GetById(sub.GetServiceId())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_STOP)
	if err != nil {
		log.Println(err.Error())
	}

	sub.SetLatestTrxId(trxId)
	// unsub
	h.Unsub(sub)

	// If SMS Alerte
	if category == CATEGORY_SMSALERTE {
		// unfollow league
		if h.subscriptionFollowLeagueService.IsSub(sub.GetId()) {
			h.subscriptionFollowLeagueService.Disable(
				&entity.SubscriptionFollowLeague{
					SubscriptionID: sub.GetId(),
				},
			)
		}

		// unfollow team
		if h.subscriptionFollowTeamService.IsSub(sub.GetId()) {
			h.subscriptionFollowTeamService.Disable(
				&entity.SubscriptionFollowTeam{
					SubscriptionID: sub.GetId(),
				},
			)
		}
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

func (h *SMSHandler) Unvalid(v string) {
	trxId := utils.GenerateTrxId()

	service, err := h.getServiceSMSAlerteDaily()
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

func (h *SMSHandler) IsActiveSubByCategory(v string) bool {
	return h.subscriptionService.IsActiveSubscriptionByCategory(v, h.req.GetMsisdn())
}

func (h *SMSHandler) Unsub(sub *entity.Subscription) {
	summary := &entity.Summary{
		ServiceID: sub.GetServiceId(),
		CreatedAt: time.Now(),
	}

	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID:     sub.GetServiceId(),
			Msisdn:        sub.GetMsisdn(),
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
			TotalUnsub: sub.TotalUnsub + 1,
		},
	)

	h.transactionService.Save(
		&entity.Transaction{
			TrxId:        sub.GetLatestTrxId(),
			ServiceID:    sub.GetServiceId(),
			Msisdn:       h.req.GetMsisdn(),
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
			Keyword:        h.req.GetSMS(),
			Subject:        SUBJECT_UNSUB,
			Status:         STATUS_SUCCESS,
			IpAddress:      h.req.GetIpAddress(),
		},
	)

	// setter summary
	summary.SetTotalUnsub(1)
	// save summary
	h.summaryService.Save(summary)

	s := &entity.Subscription{
		ServiceID: sub.GetServiceId(),
		Msisdn:    sub.GetMsisdn(),
	}

	// set false is_active
	h.subscriptionService.UpdateNotActive(s)
	// set false is_retry
	h.subscriptionService.UpdateNotRetry(s)
}

func (h *SMSHandler) IsActiveSub() bool {
	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsActiveSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *SMSHandler) IsSub() bool {
	service, err := h.getServiceSMSAlerteDaily()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *SMSHandler) getServiceSMSAlerteDaily() (*entity.Service, error) {
	return h.serviceService.Get("SA1")
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

func (h *SMSHandler) getContentFollowCompetition(v string, service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowCompetition(v, service, league)
}

func (h *SMSHandler) getContentFollowTeam(v string, service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(v) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetFollowTeam(v, service, team)
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
