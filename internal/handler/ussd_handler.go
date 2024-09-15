package handler

import (
	"log"

	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type UssdHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	menuService         services.IMenuService
	ussdService         services.IUssdService
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	historyService      services.IHistoryService
	summaryService      services.ISummaryService
	req                 *model.UssdRequest
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
	summaryService services.ISummaryService,
	req *model.UssdRequest,
) *UssdHandler {
	return &UssdHandler{
		rmq:                 rmq,
		logger:              logger,
		menuService:         menuService,
		ussdService:         ussdService,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		historyService:      historyService,
		summaryService:      summaryService,
		req:                 req,
	}
}

func (h *UssdHandler) USSD() {
	log.Println(h.req)
}
