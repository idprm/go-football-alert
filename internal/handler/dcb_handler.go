package handler

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/wiliehidayat87/rmqp"
)

type DCBHandler struct {
	rmq                             rmqp.AMQP
	summaryService                  services.ISummaryService
	menuService                     services.IMenuService
	ussdService                     services.IUssdService
	scheduleService                 services.IScheduleService
	serviceService                  services.IServiceService
	contentService                  services.IContentService
	subscriptionService             services.ISubscriptionService
	subscriptionCreditGoalService   services.ISubscriptionCreditGoalService
	subscriptionPredictWinService   services.ISubscriptionPredictWinService
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService
	subscriptionFollowTeamService   services.ISubscriptionFollowTeamService
	transactionService              services.ITransactionService
	historyService                  services.IHistoryService
	mtService                       services.IMTService
	smsAlerteService                services.ISMSAlerteService
	pronosticService                services.IPronosticService
}

func NewDCBHandler(
	rmq rmqp.AMQP,
	summaryService services.ISummaryService,
	menuService services.IMenuService,
	ussdService services.IUssdService,
	scheduleService services.IScheduleService,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	subscriptionCreditGoalService services.ISubscriptionCreditGoalService,
	subscriptionPredictWinService services.ISubscriptionPredictWinService,
	subscriptionFollowLeagueService services.ISubscriptionFollowLeagueService,
	subscriptionFollowTeamService services.ISubscriptionFollowTeamService,
	transactionService services.ITransactionService,
	historyService services.IHistoryService,
	mtService services.IMTService,
	smsAlerteService services.ISMSAlerteService,
	pronosticService services.IPronosticService,
) *DCBHandler {
	return &DCBHandler{
		rmq:                             rmq,
		summaryService:                  summaryService,
		menuService:                     menuService,
		ussdService:                     ussdService,
		scheduleService:                 scheduleService,
		serviceService:                  serviceService,
		contentService:                  contentService,
		subscriptionService:             subscriptionService,
		subscriptionCreditGoalService:   subscriptionCreditGoalService,
		subscriptionPredictWinService:   subscriptionPredictWinService,
		subscriptionFollowLeagueService: subscriptionFollowLeagueService,
		subscriptionFollowTeamService:   subscriptionFollowTeamService,
		transactionService:              transactionService,
		historyService:                  historyService,
		mtService:                       mtService,
		smsAlerteService:                smsAlerteService,
		pronosticService:                pronosticService,
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

func (h *DCBHandler) Unsubscription(c *fiber.Ctx) error {
	if h.subscriptionService.IsActiveSubscriptionBySubId(1) {
		trxId := utils.GenerateTrxId()

		sub, _ := h.subscriptionService.GetBySubId(1)

		service, _ := h.serviceService.GetById(sub.GetServiceId())

		content, err := h.getContentService(SMS_STOP, service)
		if err != nil {
			log.Println(err.Error())
		}

		if sub.IsFollowLeague {
			h.subscriptionService.UpdateNotFollowLeague(sub)
		}

		if sub.IsFollowTeam {
			h.subscriptionService.UpdateNotFollowTeam(sub)
		}

		h.subscriptionService.UpdateNotActive(sub)

		mt := &model.MTRequest{
			Smsc:         service.ScUnsubMT,
			Keyword:      "STOP " + sub.GetLatestKeyword(),
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

		return c.Status(fiber.StatusOK).JSON(
			&model.WebResponse{
				Error:      false,
				StatusCode: fiber.StatusOK,
				Message:    "",
			},
		)
	}
	return c.Status(fiber.StatusNotFound).JSON(
		&model.WebResponse{
			Error:      true,
			StatusCode: fiber.StatusNotFound,
			Message:    "not_found",
		},
	)
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

func (h *DCBHandler) GetAllPronosticPaginate(c *fiber.Ctx) error {
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

	pronostics, err := h.pronosticService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(pronostics)
}

func (h *DCBHandler) SavePronostic(c *fiber.Ctx) error {
	req := new(model.PronosticRequest)

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

	if !h.pronosticService.IsPronosticByFixtureId(int(req.FixtureID)) {
		h.pronosticService.Save(
			&entity.Pronostic{
				FixtureID: req.FixtureID,
				Category:  req.Category,
				Value:     req.Value,
				PublishAt: req.PublishAt,
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

	h.pronosticService.Update(
		&entity.Pronostic{
			FixtureID: req.FixtureID,
			Category:  req.Category,
			Value:     req.Value,
			PublishAt: req.PublishAt,
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

func (h *DCBHandler) getContentService(v string, service *entity.Service) (*entity.Content, error) {
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
