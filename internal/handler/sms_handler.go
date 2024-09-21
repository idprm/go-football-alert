package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type SMSHandler struct {
	rmq                                  rmqp.AMQP
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

func (h *SMSHandler) Registration() {

	if h.leagueService.IsLeagueByName(h.req.GetText()) {
		// SMS-Alerte Competition
		h.AlerteCompetition()
		// SMS-Alerte Matchs
	} else if h.teamService.IsTeamByName(h.req.GetText()) {
		// SMS-Alerte Equipe
		h.AlerteEquipe()
		// SMS-Alerte Matchs
	} else if h.req.IsInfo() {
		h.Info()
	} else if h.req.IsStop() {
		h.Stop()
	} else {
		h.Unvalid()
	}
}

func (h *SMSHandler) AlerteCompetition() {
	// league, err := h.leagueService.Get(h.req.GetText())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
}

func (h *SMSHandler) AlerteEquipe() {
	// team, err := h.teamService.Get(h.req.GetText())
	// if err != nil {
	// 	log.Println(err.Error())
	// }

}

func (h *SMSHandler) AlerteMatchs() {

}

func (h *SMSHandler) Info() {

}

func (h *SMSHandler) Stop() {

}

func (h *SMSHandler) Unvalid() {

}

func (h *SMSHandler) IsSub() {

}

func (h *SMSHandler) IsActiveSub() {

}

func (h *SMSHandler) getContent(name string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(name)
}
