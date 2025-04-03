package handler

import (
	"encoding/json"
	"log"
	"time"

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
	leagueService                   services.ILeagueService
	teamService                     services.ITeamService
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
	moService                       services.IMOService
	mtService                       services.IMTService
	smsAlerteService                services.ISMSAlerteService
	pronosticService                services.IPronosticService
}

func NewDCBHandler(
	rmq rmqp.AMQP,
	summaryService services.ISummaryService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
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
	moService services.IMOService,
	mtService services.IMTService,
	smsAlerteService services.ISMSAlerteService,
	pronosticService services.IPronosticService,
) *DCBHandler {
	return &DCBHandler{
		rmq:                             rmq,
		summaryService:                  summaryService,
		leagueService:                   leagueService,
		teamService:                     teamService,
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
		moService:                       moService,
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
				Message:    "please set start_date and end_date",
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

	totalActiveSub, err := h.summaryService.GetActiveSub()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalSub, err := h.summaryService.GetSub(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalUnsub, err := h.summaryService.GetUnSub(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalRenewal, err := h.summaryService.GetRenewal(req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}

	totalRevenue, err := h.summaryService.GetRevenue(req.GetStartDate(), req.GetEndDate())
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
			StartDate:      req.GetStartDate().String(),
			EndDate:        req.GetEndDate().String(),
			TotalActiveSub: totalActiveSub,
			TotalSub:       totalSub,
			TotalUnsub:     totalUnsub,
			TotalRenewal:   totalRenewal,
			TotalRevenue:   totalRevenue,
			Results:        summaries,
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

	req := new(model.UnsubRequest)

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

	if h.subscriptionService.IsActiveSubscriptionBySubId(int64(req.GetId())) {
		trxId := utils.GenerateTrxId()

		sub, _ := h.subscriptionService.GetBySubId(int64(req.GetId()))

		service, err := h.serviceService.GetById(sub.GetServiceId())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				&model.WebResponse{
					Error:      false,
					StatusCode: fiber.StatusInternalServerError,
					Message:    err.Error(),
				},
			)
		}

		if sub.ISMSAlerte() {
			// is_follow competition
			if sub.IsCompetition() {
				// check league code
				if h.leagueService.IsLeagueByCode(sub.GetCode()) {
					league, err := h.leagueService.GetByCode(sub.GetCode())
					if err != nil {
						log.Println(err.Error())
					}
					// unfollow league
					if h.subscriptionFollowLeagueService.IsSub(sub.GetId(), league.GetId()) {
						h.subscriptionFollowLeagueService.Update(
							&entity.SubscriptionFollowLeague{
								SubscriptionID: sub.GetId(),
								LeagueID:       league.GetId(),
								LatestKeyword:  "",
							},
						)

						// disable
						h.subscriptionFollowLeagueService.Disable(
							&entity.SubscriptionFollowLeague{
								SubscriptionID: sub.GetId(),
								LeagueID:       league.GetId(),
							},
						)

						content, err := h.getContentUnFollowCompetition(service, league)
						if err != nil {
							log.Println(err.Error())
						}

						sub.SetLatestTrxId(trxId)

						// unsub sms-alerte
						h.subscriptionService.Update(
							&entity.Subscription{
								ServiceID:     sub.GetServiceId(),
								Msisdn:        sub.GetMsisdn(),
								Code:          sub.GetCode(),
								LatestTrxId:   sub.GetLatestTrxId(),
								LatestSubject: SUBJECT_UNSUB,
								LatestStatus:  STATUS_SUCCESS,
								UnsubAt:       time.Now(),
								TotalUnsub:    sub.TotalUnsub + 1,
							},
						)

						h.historyService.Save(
							&entity.History{
								SubscriptionID: sub.GetId(),
								ServiceID:      sub.GetServiceId(),
								Msisdn:         sub.GetMsisdn(),
								Code:           sub.GetCode(),
								Subject:        SUBJECT_UNSUB,
								Status:         STATUS_SUCCESS,
							},
						)

						s := &entity.Subscription{
							ServiceID: sub.GetServiceId(),
							Msisdn:    sub.GetMsisdn(),
							Code:      sub.GetCode(),
						}

						// set false is_active
						h.subscriptionService.UpdateNotActive(s)
						// set false is_retry
						h.subscriptionService.UpdateNotRetry(s)
						// set false is_follow_league
						h.subscriptionService.UpdateNotFollowLeague(sub)

						mt := &model.MTRequest{
							Smsc:         service.ScUnsubMT,
							Keyword:      "STOP " + sub.GetCode(),
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
				}

			}
			// is_follow team
			if sub.IsEquipe() {
				// check team code
				if h.teamService.IsTeamByCode(sub.GetCode()) {
					team, err := h.teamService.GetByCode(sub.GetCode())
					if err != nil {
						log.Println(err.Error())
					}
					// unfollow team
					if h.subscriptionFollowTeamService.IsSub(sub.GetId(), team.GetId()) {
						h.subscriptionFollowTeamService.Disable(
							&entity.SubscriptionFollowTeam{
								SubscriptionID: sub.GetId(),
								TeamID:         team.GetId(),
							},
						)
					}

					content, err := h.getContentUnFollowTeam(service, team)
					if err != nil {
						log.Println(err.Error())
					}

					sub.SetLatestTrxId(trxId)

					// unsub sms-alerte
					h.subscriptionService.Update(
						&entity.Subscription{
							ServiceID:     sub.GetServiceId(),
							Msisdn:        sub.GetMsisdn(),
							Code:          sub.GetCode(),
							LatestTrxId:   sub.GetLatestTrxId(),
							LatestSubject: SUBJECT_UNSUB,
							LatestStatus:  STATUS_SUCCESS,
							LatestKeyword: sub.GetLatestKeyword(),
							UnsubAt:       time.Now(),
							IpAddress:     sub.GetIpAddress(),
							TotalUnsub:    sub.TotalUnsub + 1,
						},
					)

					h.historyService.Save(
						&entity.History{
							SubscriptionID: sub.GetId(),
							ServiceID:      sub.GetServiceId(),
							Msisdn:         sub.GetMsisdn(),
							Code:           sub.GetCode(),
							Keyword:        sub.GetLatestKeyword(),
							Subject:        SUBJECT_UNSUB,
							Status:         STATUS_SUCCESS,
							IpAddress:      sub.GetIpAddress(),
						},
					)

					s := &entity.Subscription{
						ServiceID: sub.GetServiceId(),
						Msisdn:    sub.GetMsisdn(),
						Code:      sub.GetCode(),
					}

					// set false is_active
					h.subscriptionService.UpdateNotActive(s)
					// set false is_retry
					h.subscriptionService.UpdateNotRetry(s)
					// set false is_follow_team
					h.subscriptionService.UpdateNotFollowTeam(sub)

					mt := &model.MTRequest{
						Smsc:         service.ScUnsubMT,
						Keyword:      "STOP " + sub.GetCode(),
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
			}
		} else {
			content, err := h.getContentService("STOP", service)
			if err != nil {
				log.Println(err.Error())
			}

			sub.SetLatestTrxId(trxId)

			// unsub sms-alerte
			h.subscriptionService.Update(
				&entity.Subscription{
					ServiceID:     sub.GetServiceId(),
					Msisdn:        sub.GetMsisdn(),
					Code:          sub.GetCode(),
					LatestTrxId:   sub.GetLatestTrxId(),
					LatestSubject: SUBJECT_UNSUB,
					LatestStatus:  STATUS_SUCCESS,
					LatestKeyword: sub.GetLatestKeyword(),
					UnsubAt:       time.Now(),
					IpAddress:     sub.GetIpAddress(),
					TotalUnsub:    sub.TotalUnsub + 1,
				},
			)

			h.historyService.Save(
				&entity.History{
					SubscriptionID: sub.GetId(),
					ServiceID:      sub.GetServiceId(),
					Msisdn:         sub.GetMsisdn(),
					Code:           sub.GetCode(),
					Keyword:        sub.GetLatestKeyword(),
					Subject:        SUBJECT_UNSUB,
					Status:         STATUS_SUCCESS,
					IpAddress:      sub.GetIpAddress(),
				},
			)

			s := &entity.Subscription{
				ServiceID: sub.GetServiceId(),
				Msisdn:    sub.GetMsisdn(),
				Code:      sub.GetCode(),
			}

			// set false is_active
			h.subscriptionService.UpdateNotActive(s)
			// set false is_retry
			h.subscriptionService.UpdateNotRetry(s)

			mt := &model.MTRequest{
				Smsc:         service.ScUnsubMT,
				Keyword:      "STOP " + sub.GetCode(),
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

		return c.Status(fiber.StatusOK).JSON(
			&model.WebResponse{
				Error:      false,
				StatusCode: fiber.StatusOK,
				Message:    "unsub_success",
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

func (h *DCBHandler) GetAllMOPaginate(c *fiber.Ctx) error {
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

	mos, err := h.moService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(mos)
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

	startAt, err := time.Parse("2006-01-02T15:04:05-0700", req.StartAt)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(err)
	}

	expireAt, err := time.Parse("2006-01-02T15:04:05-0700", req.ExpireAt)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(err)
	}

	if !h.pronosticService.IsPronosticByStartAt(startAt) {
		prono := &entity.Pronostic{
			Category: req.Category,
			Value:    req.Value,
			StartAt:  startAt,
			ExpireAt: expireAt,
		}

		h.pronosticService.Save(prono)

		p, _ := h.pronosticService.Get(startAt)

		go h.SMSPronostic(p.GetId())

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
			Category: req.Category,
			Value:    req.Value,
			StartAt:  startAt,
			ExpireAt: expireAt,
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

func (h *DCBHandler) getContentUnFollowCompetition(service *entity.Service, league *entity.League) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_COMPETITION_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TvEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowCompetition(SMS_FOLLOW_COMPETITION_STOP, service, league)
}

func (h *DCBHandler) getContentUnFollowTeam(service *entity.Service, team *entity.Team) (*entity.Content, error) {
	if !h.contentService.IsContent(SMS_FOLLOW_TEAM_STOP) {
		return &entity.Content{
			Category: "CATEGORY",
			Channel:  "SMS",
			Value:    "SAMPLE_TEXT",
		}, nil
	}
	return h.contentService.GetUnSubFollowTeam(SMS_FOLLOW_TEAM_STOP, service, team)
}

func (h *DCBHandler) SMSPronostic(pronoId int64) {
	// valid in team
	subs := h.subscriptionService.Prono()

	if len(*subs) > 0 {
		for _, s := range *subs {
			jsonData, err := json.Marshal(&entity.SMSProno{SubscriptionID: s.ID, PronosticID: pronoId})
			if err != nil {
				log.Println(err.Error())
			}

			h.rmq.IntegratePublish(
				RMQ_SMS_PRONO_EXCHANGE,
				RMQ_SMS_PRONO_QUEUE,
				RMQ_DATA_TYPE, "", string(jsonData),
			)
		}
	}
}

func (h *DCBHandler) ReportDaily(c *fiber.Ctx) error {
	return c.Render("", "")
}
