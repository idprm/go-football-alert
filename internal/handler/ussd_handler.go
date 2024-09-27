package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/sirupsen/logrus"
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

func (h *UssdHandler) Registration() {
	l := h.logger.Init("ussd", true)
	l.WithFields(logrus.Fields{"request": h.req}).Info("USSD")
}
func (h *UssdHandler) FlashNews(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) CreditGoal(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) LiveMatch(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) LineUp(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) Schedule(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) MatchStats(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampResults(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampStandings(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampSchedules(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampCreditVictoire(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampCreditGoal(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampSmsAlerte(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) ChampSmsAlerteEquipe(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) PredictionDailySafeOfTheDay(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) PredictionDailyCombinedBets(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) PredictionChampVipProno(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) SmsAlerteEurope(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) SmsAlerteAfrique(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) KitFoot(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *UssdHandler) FootEurope(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}
