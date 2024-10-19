package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type DCBHandler struct {
	summaryService      services.ISummaryService
	menuService         services.IMenuService
	ussdService         services.IUssdService
	scheduleService     services.IScheduleService
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	historyService      services.IHistoryService
	mtService           services.IMTService
	smsAlerteService    services.ISMSAlerteService
}

func NewDCBHandler(
	summaryService services.ISummaryService,
	menuService services.IMenuService,
	ussdService services.IUssdService,
	scheduleService services.IScheduleService,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	mtService services.IMTService,
	smsAlerteService services.ISMSAlerteService,
) *DCBHandler {
	return &DCBHandler{
		summaryService:      summaryService,
		menuService:         menuService,
		ussdService:         ussdService,
		scheduleService:     scheduleService,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		historyService:      historyService,
		mtService:           mtService,
		smsAlerteService:    smsAlerteService,
	}
}

func (h *DCBHandler) GetAllSummaryPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	if !req.IsDate() {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    "please set month of date",
			},
		)
	}

	summaries, err := h.summaryService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalSub, err := h.summaryService.GetSubByMonth(req.GetDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalUnsub, err := h.summaryService.GetUnsubByMonth(req.GetDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalRenewal, err := h.summaryService.GetRenewalByMonth(req.GetDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalRevenue, err := h.summaryService.GetRevenueByMonth(req.GetDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		&model.SummaryResponse{
			Month:        req.GetDate().Month().String(),
			Year:         req.GetDate().Year(),
			TotalSub:     totalSub,
			TotalUnsub:   totalUnsub,
			TotalRenewal: totalRenewal,
			TotalRevenue: totalRevenue,
			Results:      summaries,
		},
	)
}

func (h *DCBHandler) GetAllMenuPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	menus, err := h.menuService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(menus)
}

func (h *DCBHandler) GetAllUssdPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	ussds, err := h.ussdService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(ussds)
}

func (h *DCBHandler) GetAllSchedulePaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	schedules, err := h.scheduleService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(schedules)
}

func (h *DCBHandler) GetAllServicePaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	services, err := h.serviceService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(services)
}

func (h *DCBHandler) GetAllContentPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	contents, err := h.contentService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(contents)
}

func (h *DCBHandler) GetAllSubscriptionPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	subs, err := h.subscriptionService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(subs)
}

func (h *DCBHandler) GetAllTransactionPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	trans, err := h.transactionService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(trans)
}

func (h *DCBHandler) GetAllHistoryPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	histories, err := h.historyService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(histories)
}

func (h *DCBHandler) GetAllMTPaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	mts, err := h.mtService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(mts)
}

func (h *DCBHandler) GetAllSMSAlertePaginate(c *fiber.Ctx) error {
	req := new(entity.Pagination)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	alerts, err := h.smsAlerteService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(alerts)
}
