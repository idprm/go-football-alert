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
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/wiliehidayat87/rmqp"
)

type SMSHandler struct {
	rmq                                  rmqp.AMQP
	rds                                  *redis.Client
	logger                               *logger.Logger
	serviceService                       services.IServiceService
	contentService                       services.IContentService
	subscriptionService                  services.ISubscriptionService
	transactionService                   services.ITransactionService
	historyService                       services.IHistoryService
	summaryService                       services.ISummaryService
	leagueService                        services.ILeagueService
	teamService                          services.ITeamService
	subscriptionCreditGoalService        services.ISubscriptionCreditGoalService
	subscriptionPredictService           services.ISubscriptionPredictService
	subscriptionFollowCompetitionService services.ISubscriptionFollowCompetitionService
	subscriptionFollowTeamService        services.ISubscriptionFollowTeamService
	verifyService                        services.IVerifyService
	req                                  *model.MORequest
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
	subscriptionFollowCompetitionService services.ISubscriptionFollowCompetitionService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	verifyService services.IVerifyService,
	req *model.MORequest,
) *SMSHandler {
	return &SMSHandler{
		rmq:                                  rmq,
		rds:                                  rds,
		logger:                               logger,
		serviceService:                       serviceService,
		contentService:                       contentService,
		subscriptionService:                  subscriptionService,
		transactionService:                   transactionService,
		historyService:                       historyService,
		summaryService:                       summaryService,
		leagueService:                        leagueService,
		teamService:                          teamService,
		subscriptionCreditGoalService:        subscriptionCreditGoalService,
		subscriptionPredictService:           subscriptionPredictService,
		subscriptionFollowCompetitionService: subscriptionFollowCompetitionService,
		subscriptionFollowTeamService:        subscriptionFollowTeamService,
		verifyService:                        verifyService,
		req:                                  req,
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
	SMS_FOLLOW_TEAM_UNVALID_SUB          string = "FOLLOW_TEAM_UNVALID_SUB"
	SMS_FOLLOW_TEAM_EXPIRE_SUB           string = "FOLLOW_TEAM_EXPIRE_SUB"
	SMS_FOLLOW_COMPETITION_SUB           string = "FOLLOW_COMPETITION_SUB"
	SMS_FOLLOW_COMPETITION_INVALID_SUB   string = "FOLLOW_COMPETITION_INVALID_SUB"
	SMS_FOLLOW_COMPETITION_EXPIRE_SUB    string = "FOLLOW_COMPETITION_EXPIRE_SUB"
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

func (h *SMSHandler) Confirmation(v string) {
	content, err := h.getContent(SMS_CONFIRMATION)
	if err != nil {
		log.Println(err.Error())
	}

	service, err := h.serviceService.GetById(7)
	if err != nil {
		log.Println(err.Error())
	}

	k := kannel.NewKannel(
		h.logger,
		service,
		content,
		&entity.Subscription{Msisdn: h.req.GetMsisdn()},
	)

	// sent
	k.SMS(h.req.GetTo())

	// set on catch
	h.verifyService.SetCategory(
		&entity.Verify{
			Msisdn:   h.req.GetMsisdn(),
			Category: v,
		},
	)
}

func (h *SMSHandler) SMSAlerte() {
	// user choose 1, 2, 3 package
	if h.req.IsChooseService() {
		h.Subscription(CATEGORY_SMSALERTE)
	} else {
		if h.leagueService.IsLeagueByName(h.req.GetSMS()) {
			if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
				h.Confirmation(SUBCATEGORY_FOLLOW_COMPETITION)
			} else {
				// SMS-Alerte Competition
				h.AlerteCompetition()
				// SMS-Alerte Matchs
			}
		} else if h.teamService.IsTeamByName(h.req.GetSMS()) {
			if !h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
				h.Confirmation(SUBCATEGORY_FOLLOW_TEAM)
			} else {
				// SMS-Alerte Equipe
				h.AlerteEquipe()
				// SMS-Alerte Matchs
			}
		} else if h.req.IsInfo() {
			h.Info()
		} else if h.req.IsStop() {
			if h.IsActiveSubByCategory(CATEGORY_SMSALERTE) {
				h.Stop()
			}
		} else {
			h.Unvalid(SMS_FOLLOW_COMPETITION_INVALID_SUB)
		}
	}
}

func (h *SMSHandler) Subscription(category string) {
	trxId := utils.GenerateTrxId()

	service, err := h.serviceService.GetByPackage(category, h.req.GetServiceByNumber())
	if err != nil {
		log.Println(err.Error())
	}

	content, err := h.getContent(SMS_FOLLOW_COMPETITION_SUB)
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
		IpAddress:     h.req.GetIpAddress(),
	}

	if !h.subscriptionService.IsActiveSubscription(service.GetId(), h.req.GetMsisdn()) {
		h.subscriptionService.Save(
			&entity.Subscription{
				ServiceID: service.GetId(),
				Category:  service.GetCategory(),
				Msisdn:    h.req.GetMsisdn(),
				IsActive:  true,
			},
		)
	} else {
		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID: service.GetId(),
				Category:  service.GetCategory(),
				Msisdn:    h.req.GetMsisdn(),
				IsActive:  true,
			},
		)
	}

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err.Error())
	}

	t := telco.NewTelco(h.logger, service, subscription)

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
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FIRSTPUSH,
				Status:         STATUS_FAILED,
				CreatedAt:      time.Now(),
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
				CreatedAt:    time.Now(),
			},
		)

		h.historyService.Save(
			&entity.History{
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FIRSTPUSH,
				Status:         STATUS_SUCCESS,
				CreatedAt:      time.Now(),
			},
		)

		summary.SetTotalChargeSuccess(1)
		summary.SetTotalRevenue(service.GetPrice())
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
		Smsc:         h.req.GetSMS(),
		Service:      service,
		Subscription: sub,
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

func (h *SMSHandler) AlerteCompetition() {
	content, err := h.getContent(SMS_FOLLOW_COMPETITION_SUB)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
		Content:      content,
	}

	jsonData, err := json.Marshal(mt)
	if err != nil {
		log.Println(err.Error())
	}

	//h.Firstpush()

	h.rmq.IntegratePublish(
		RMQ_MT_EXCHANGE,
		RMQ_MT_QUEUE,
		RMQ_DATA_TYPE, "", string(jsonData),
	)
}

func (h *SMSHandler) AlerteEquipe() {
	content, err := h.getContent(SMS_CREDIT_GOAL_SUB)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Subscription: &entity.Subscription{},
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

func (h *SMSHandler) Info() {
	content, err := h.getContent(SMS_INFO)
	if err != nil {
		log.Println(err.Error())
	}
	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) Stop() {
	content, err := h.getContent(SMS_STOP)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) Unvalid(v string) {
	content, err := h.getContent(v)
	if err != nil {
		log.Println(err.Error())
	}

	mt := &model.MTRequest{
		Smsc:         h.req.GetTo(),
		Subscription: &entity.Subscription{Msisdn: h.req.GetMsisdn()},
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

func (h *SMSHandler) IsActiveSubByCategory(v string) bool {
	return h.subscriptionService.IsActiveSubscriptionByCategory(v, h.req.GetMsisdn())
}

func (h *SMSHandler) Firstpush() {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}

	content, err := h.getContent(MT_FIRSTPUSH)
	if err != nil {
		log.Println(err)
	}

	trxId := utils.GenerateTrxId()

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
	profileBall, err := t.QueryProfileAndBal()
	if err != nil {
		log.Println(err.Error())
	}

	deductFee, err := t.DeductFee()
	if err != nil {
		log.Println(err.Error())
	}

	var respDeduct *model.DeductResponse
	xml.Unmarshal(utils.EscapeChar(deductFee), &respDeduct)

	var respProfileBall *model.QueryProfileAndBalResponse
	xml.Unmarshal(utils.EscapeChar(profileBall), &respProfileBall)

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
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FIRSTPUSH,
				Status:         STATUS_FAILED,
				CreatedAt:      time.Now(),
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
				CreatedAt:    time.Now(),
			},
		)

		h.historyService.Save(
			&entity.History{
				CountryID:      service.GetCountryId(),
				SubscriptionID: sub.GetId(),
				ServiceID:      service.GetId(),
				Msisdn:         h.req.GetMsisdn(),
				Keyword:        h.req.GetSMS(),
				Subject:        SUBJECT_FIRSTPUSH,
				Status:         STATUS_SUCCESS,
				CreatedAt:      time.Now(),
			},
		)

		summary.SetTotalChargeSuccess(1)
		summary.SetTotalRevenue(service.GetPrice())
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
		Smsc:         "",
		Subscription: &entity.Subscription{},
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

func (h *SMSHandler) Unsub() {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}

	content, err := h.getContent(MT_UNSUB)
	if err != nil {
		log.Println(err)
	}

	trxId := utils.GenerateTrxId()

	summary := &entity.Summary{
		ServiceID: service.GetId(),
		CreatedAt: time.Now(),
	}

	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID:     service.GetId(),
			Msisdn:        h.req.GetMsisdn(),
			LatestTrxId:   trxId,
			LatestSubject: SUBJECT_UNSUB,
			LatestStatus:  STATUS_SUCCESS,
			LatestKeyword: h.req.GetSMS(),
			UnsubAt:       time.Now(),
			IpAddress:     h.req.GetIpAddress(),
		},
	)

	s := &entity.Subscription{
		ServiceID: service.GetId(),
		Msisdn:    h.req.GetMsisdn(),
	}

	// set false is_active
	h.subscriptionService.IsNotActive(s)
	// set false is_retry
	h.subscriptionService.IsNotRetry(s)

	sub, err := h.subscriptionService.Get(service.GetId(), h.req.GetMsisdn())
	if err != nil {
		log.Println(err)
	}

	h.subscriptionService.Update(
		&entity.Subscription{
			ServiceID:  service.GetId(),
			Msisdn:     h.req.GetMsisdn(),
			TotalUnsub: sub.TotalUnsub + 1,
		},
	)

	h.transactionService.Save(
		&entity.Transaction{
			TrxId:        trxId,
			ServiceID:    service.GetId(),
			Msisdn:       h.req.GetMsisdn(),
			Keyword:      h.req.GetSMS(),
			Status:       STATUS_SUCCESS,
			StatusCode:   "",
			StatusDetail: "",
			Subject:      SUBJECT_UNSUB,
			IpAddress:    h.req.GetIpAddress(),
			Payload:      "",
			CreatedAt:    time.Now(),
		},
	)

	h.historyService.Save(
		&entity.History{
			CountryID:      service.GetCountryId(),
			SubscriptionID: sub.GetId(),
			ServiceID:      service.GetId(),
			Msisdn:         h.req.GetMsisdn(),
			Keyword:        h.req.GetSMS(),
			Subject:        SUBJECT_UNSUB,
			Status:         STATUS_SUCCESS,
			IpAddress:      h.req.GetIpAddress(),
			CreatedAt:      time.Now(),
		},
	)

	// setter summary
	summary.SetTotalUnsub(1)
	// save summary
	h.summaryService.Save(summary)

	mt := &model.MTRequest{
		Smsc:         "",
		Subscription: sub,
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

func (h *SMSHandler) IsActiveSub() bool {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsActiveSubscription(service.GetId(), h.req.GetMsisdn())
}

func (h *SMSHandler) IsSub() bool {
	service, err := h.getService()
	if err != nil {
		log.Println(err)
	}
	return h.subscriptionService.IsSubscription(service.GetId(), h.req.GetMsisdn())
}

// empty service
func (h *SMSHandler) getService() (*entity.Service, error) {
	return h.serviceService.Get("")
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
	} else {
		// user choose 1, 2, 3 package
		if h.req.IsChooseService() {
			h.Subscription(CATEGORY_CREDIT_GOAL)
		} else {
			h.Unvalid(SMS_CREDIT_GOAL_UNVALID_SUB)
		}
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
	} else {
		// user choose 1, 2, 3 package
		if h.req.IsChooseService() {
			h.Subscription(CATEGORY_PREDICT)
		} else {
			h.Unvalid(SMS_PREDICT_UNVALID_SUB)
		}
	}
}
**/
