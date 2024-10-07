package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type CreditGoalHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	sub                 *entity.Subscription
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	rewardService       services.IRewardService
	summaryService      services.ISummaryService
}

func NewCreditGoalHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	rewardService services.IRewardService,
	summaryService services.ISummaryService,
) *CreditGoalHandler {
	return &CreditGoalHandler{
		rmq:                 rmq,
		logger:              logger,
		sub:                 sub,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		rewardService:       rewardService,
		summaryService:      summaryService,
	}
}

func (h *CreditGoalHandler) CreditGoal() {
	if h.subscriptionService.IsActiveSubscription(h.sub.GetServiceId(), h.sub.GetMsisdn()) {
		// service, err := h.serviceService.GetById(h.sub.GetServiceId())
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// content, err := h.getContent(MT_CREDIT_GOAL)
		// if err != nil {
		// 	log.Println(err.Error())
		// }

		// trxId := utils.GenerateTrxId()

		// summary := &entity.Summary{
		// 	ServiceID: service.GetId(),
		// 	CreatedAt: time.Now(),
		// }

		// h.subscriptionService.Update(
		// 	&entity.Subscription{
		// 		ServiceID:     service.GetId(),
		// 		Msisdn:        h.sub.GetMsisdn(),
		// 		LatestTrxId:   trxId,
		// 		LatestSubject: SUBJECT_CREDIT_GOAL,
		// 		LatestStatus:  STATUS_SUCCESS,
		// 	},
		// )

		// h.transactionService.Save(
		// 	&entity.Transaction{
		// 		TrxId:        trxId,
		// 		ServiceID:    service.GetId(),
		// 		Msisdn:       h.sub.GetMsisdn(),
		// 		Keyword:      h.sub.GetLatestKeyword(),
		// 		Status:       STATUS_SUCCESS,
		// 		StatusCode:   "",
		// 		StatusDetail: "",
		// 		Subject:      SUBJECT_CREDIT_GOAL,
		// 		Payload:      "-",
		// 		CreatedAt:    time.Now(),
		// 	},
		// )

		// // setter summary
		// summary.SetTotalCreditGoal(1)

		// // summary save
		// h.summaryService.Save(summary)
	}
}

func (h *CreditGoalHandler) getContent(name string) (*entity.Content, error) {
	// if data not exist in table contents
	if !h.contentService.IsContent(name) {
		return &entity.Content{
			Value: "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.Get(name)
}
