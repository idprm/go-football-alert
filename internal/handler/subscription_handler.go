package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type SubscriptionHandler struct {
	subscriptionService services.ISubscriptionService
}

func NewSubscriptionHandler(
	subscriptionService services.ISubscriptionService,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

func (h *SubscriptionHandler) GetAllPaginate(c *fiber.Ctx) error {
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

	subscriptions, err := h.subscriptionService.GetAllPaginate(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(subscriptions)
}

func (h *SubscriptionHandler) Get(c *fiber.Ctx) error {
	service, _ := strconv.Atoi(c.Params("service"))
	msisdn := c.Params("msisdn")

	if !h.subscriptionService.IsSubscription(service, msisdn) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	subscription, err := h.subscriptionService.Get(service, msisdn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusInternalServerError,
				Message:    err.Error(),
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(subscription)
}

func (h *SubscriptionHandler) Save(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *SubscriptionHandler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *SubscriptionHandler) Delete(c *fiber.Ctx) error {
	service, _ := strconv.Atoi(c.Params("service"))
	msisdn := c.Params("msisdn")

	if !h.subscriptionService.IsSubscription(service, msisdn) {
		return c.Status(fiber.StatusNotFound).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusNotFound,
				Message:    "not_found",
			},
		)
	}

	data := &entity.Subscription{
		ServiceID: service,
		Msisdn:    msisdn,
	}

	err := h.subscriptionService.Delete(data)
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
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "deleted",
		},
	)
}
