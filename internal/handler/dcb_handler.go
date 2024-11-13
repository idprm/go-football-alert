package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
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

func (h *DCBHandler) SaveMenu(c *fiber.Ctx) error {
	req := new(model.MenuRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	errors := ValidateStruct(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if !h.menuService.IsSlug(slug.Make(req.Name)) {
		h.menuService.Save(
			&entity.Menu{
				Category:    req.Category,
				Name:        req.GetName(),
				Slug:        slug.Make(req.GetName()),
				TemplateXML: req.TemplateXML,
				IsConfirm:   false,
				IsActive:    req.IsActive,
			},
		)

		return c.Status(fiber.StatusCreated).JSON(
			&model.WebResponse{
				Error:      false,
				StatusCode: fiber.StatusCreated,
				Message:    "created",
			},
		)
	}

	h.menuService.Update(
		&entity.Menu{
			Category:    req.Category,
			Slug:        slug.Make(req.GetName()),
			TemplateXML: req.TemplateXML,
			IsActive:    req.IsActive,
		},
	)

	return c.Status(fiber.StatusOK).JSON(
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "updated",
		},
	)
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

func (h *DCBHandler) SaveService(c *fiber.Ctx) error {
	req := new(model.ServiceRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	errors := ValidateStruct(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if !h.serviceService.IsService(req.Code) {
		h.serviceService.Save(
			&entity.Service{
				Channel:    req.Channel,
				Category:   req.Category,
				Name:       req.Name,
				Code:       req.GetCode(),
				Package:    req.Package,
				Price:      req.Price,
				Currency:   req.Currency,
				RewardGoal: req.RewardGoal,
				RenewalDay: req.RenewalDay,
				FreeDay:    req.FreeDay,
				UrlTelco:   req.UrlTelco,
				UserTelco:  req.UserTelco,
				PassTelco:  req.PassTelco,
				UrlMT:      req.UrlMT,
				UserMT:     req.UserMT,
				PassMT:     req.PassMT,
				ScSubMT:    req.ScSubMT,
				ScUnsubMT:  req.ScUnsubMT,
				ShortCode:  req.ShortCode,
				UssdCode:   req.UssdCode,
			},
		)

		return c.Status(fiber.StatusCreated).JSON(
			&model.WebResponse{
				Error:      false,
				StatusCode: fiber.StatusCreated,
				Message:    "created",
			},
		)
	}

	h.serviceService.Update(
		&entity.Service{
			Channel:    req.Channel,
			Category:   req.Category,
			Name:       req.Name,
			Code:       req.GetCode(),
			Package:    req.Package,
			Price:      req.Price,
			Currency:   req.Currency,
			RewardGoal: req.RewardGoal,
			RenewalDay: req.RenewalDay,
			FreeDay:    req.FreeDay,
			UrlTelco:   req.UrlTelco,
			UserTelco:  req.UserTelco,
			PassTelco:  req.PassTelco,
			UrlMT:      req.UrlMT,
			UserMT:     req.UserMT,
			PassMT:     req.PassMT,
			ScSubMT:    req.ScSubMT,
			ScUnsubMT:  req.ScUnsubMT,
			ShortCode:  req.ShortCode,
			UssdCode:   req.UssdCode,
		},
	)

	return c.Status(fiber.StatusOK).JSON(
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "updated",
		},
	)
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

func (h *DCBHandler) SaveContent(c *fiber.Ctx) error {
	req := new(model.ContentRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	errors := ValidateStruct(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if !h.contentService.IsContent(req.Name) {
		h.contentService.Save(
			&entity.Content{
				Category: req.Category,
				Name:     req.GetName(),
				Channel:  req.Channel,
				Value:    req.Value,
			},
		)

		return c.Status(fiber.StatusCreated).JSON(
			&model.WebResponse{
				Error:      false,
				StatusCode: fiber.StatusCreated,
				Message:    "created",
			},
		)
	}

	h.contentService.Update(
		&entity.Content{
			Category: req.Category,
			Name:     req.GetName(),
			Channel:  req.Channel,
			Value:    req.Value,
		},
	)

	return c.Status(fiber.StatusOK).JSON(
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "updated",
		},
	)
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
