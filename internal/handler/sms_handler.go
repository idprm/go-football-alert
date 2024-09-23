package handler

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/redis/go-redis/v9"
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
	req                                  *model.SMSRequest
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
	req *model.SMSRequest,
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
		req:                                  req,
	}
}

const (
	CATEGORY_CREDIT_GOAL                 string = "CREDIT-GOAL"
	CATEGORY_PREDICT                     string = "PREDICTION"
	CATEGORY_FOLLOW_TEAM                 string = "FOLLOW-TEAM"
	CATEGORY_FOLLOW_COMPETITION          string = "FOLLOW-COMPETITION"
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
	/**
	 ** Credit Goal
	 **/
	if h.req.IsCreditGoal() {
		h.CreditGoal()
	}

	/**
	 ** Prediction
	 **/
	if h.req.IsPrediction() {
		h.Prediction()
	}

	/**
	 ** Follow Team
	 **/
	if h.req.IsFollowTeam() {
		h.FollowTeam()
	}

	/**
	 ** Follow Competition
	 **/
	if h.req.IsFollowCompetition() {
		h.FollowCompetition()
	}
}

func (h *SMSHandler) CreditGoal() {
	if h.leagueService.IsLeagueByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_CREDIT_GOAL) {
			h.Confirmation()
		} else {
			// SMS-Alerte Competition
			h.AlerteCompetition()
		}
		// SMS-Alerte Matchs
	} else if h.teamService.IsTeamByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_CREDIT_GOAL) {
			h.Confirmation()
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
			h.Unvalid()
		}
	}
}

func (h *SMSHandler) Prediction() {
	if h.leagueService.IsLeagueByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_PREDICT) {
			h.Confirmation()
		} else {
			// SMS-Alerte Competition
			h.AlerteCompetition()
			// SMS-Alerte Matchs
		}
	} else if h.teamService.IsTeamByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_PREDICT) {
			h.Confirmation()
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
			h.Unvalid()
		}
	}
}

func (h *SMSHandler) FollowTeam() {
	if h.teamService.IsTeamByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_FOLLOW_TEAM) {
			h.Confirmation()
		} else {
			// SMS-Alerte Equipe
			h.AlerteEquipe()
			// SMS-Alerte Matchs
		}
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		if h.IsActiveSubByCategory(CATEGORY_FOLLOW_TEAM) {
			h.Stop()
		}
	} else {
		// user choose 1, 2, 3 package
		if h.req.IsChooseService() {
			h.Subscription(CATEGORY_FOLLOW_TEAM)
		} else {
			h.Unvalid()
		}
	}
}

func (h *SMSHandler) FollowCompetition() {
	if h.leagueService.IsLeagueByName(h.req.GetText()) {
		if !h.IsActiveSubByCategory(CATEGORY_FOLLOW_COMPETITION) {
			h.Confirmation()
		} else {
			// SMS-Alerte Competition
			h.AlerteCompetition()
			// SMS-Alerte Matchs
		}
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		if h.IsActiveSubByCategory(CATEGORY_FOLLOW_COMPETITION) {
			h.Stop()
		}
	} else {
		// user choose 1, 2, 3 package
		if h.req.IsChooseService() {
			h.Subscription(CATEGORY_FOLLOW_COMPETITION)
		} else {
			h.Unvalid()
		}
	}
}

func (h *SMSHandler) Confirmation() {
	content, err := h.getContent(SMS_CONFIRMATION)
	if err != nil {
		log.Println(err.Error())
	}
	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	// sent
	k.SMS(h.req.GetSmsc())
}

func (h *SMSHandler) Subscription(category string) {
	trxId := utils.GenerateTrxId()

	service, err := h.serviceService.GetByPackage(category, h.req.GetServiceByNumber())
	if err != nil {
		log.Println(err.Error())
	}
	if !h.subscriptionService.IsActiveSubscription(service.GetId(), h.req.GetTo()) {
		h.subscriptionService.Save(
			&entity.Subscription{
				ServiceID: service.GetId(),
				Category:  service.GetCategory(),
				Msisdn:    h.req.GetTo(),
				IsActive:  true,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.req.GetTo(),
				Keyword:      "",
				Status:       "",
				StatusCode:   "",
				StatusDetail: "",
				Subject:      "",
				Payload:      "",
				CreatedAt:    time.Now(),
			},
		)
	} else {
		h.subscriptionService.Update(
			&entity.Subscription{
				ServiceID: service.GetId(),
				Category:  service.GetCategory(),
				Msisdn:    h.req.GetTo(),
				IsActive:  true,
			},
		)

		h.transactionService.Save(
			&entity.Transaction{
				TrxId:        trxId,
				ServiceID:    service.GetId(),
				Msisdn:       h.req.GetTo(),
				Keyword:      "",
				Status:       "",
				StatusCode:   "",
				StatusDetail: "",
				Subject:      "",
				Payload:      "",
				CreatedAt:    time.Now(),
			},
		)
	}
}

func (h *SMSHandler) AlerteCompetition() {
	// league, err := h.leagueService.Get(h.req.GetText())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	content, err := h.getContent(SMS_CREDIT_GOAL_SUB)
	if err != nil {
		log.Println(err.Error())
	}
	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	// sent
	k.SMS(h.req.GetSmsc())
}

func (h *SMSHandler) AlerteEquipe() {
	// team, err := h.teamService.Get(h.req.GetText())
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	content, err := h.getContent(SMS_CREDIT_GOAL_SUB)
	if err != nil {
		log.Println(err.Error())
	}
	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	// sent
	k.SMS(h.req.GetSmsc())

}

func (h *SMSHandler) AlerteMatchs() {

}

func (h *SMSHandler) Info() {
	content, err := h.getContent(SMS_INFO)
	if err != nil {
		log.Println(err.Error())
	}

	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	k.SMS("")
}

func (h *SMSHandler) Stop() {
	content, err := h.getContent(SMS_STOP)
	if err != nil {
		log.Println(err.Error())
	}

	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	k.SMS("")
}

func (h *SMSHandler) Unvalid() {
	content, err := h.getContent(SMS_CREDIT_GOAL_UNVALID_SUB)
	if err != nil {
		log.Println(err.Error())
	}

	k := kannel.NewKannel(h.logger, &entity.Service{}, content, &entity.Subscription{Msisdn: h.req.GetTo()})
	k.SMS(h.req.GetSmsc())
}

func (h *SMSHandler) IsActiveSubByCategory(v string) bool {
	return h.subscriptionService.IsActiveSubscriptionByCategory(v, h.req.GetTo())
}

func (h *SMSHandler) IsSub() bool {
	return h.subscriptionService.IsSubscription(1, h.req.GetTo())
}

func (h *SMSHandler) getContent(name string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(name)
}
