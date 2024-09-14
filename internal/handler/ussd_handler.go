package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type UssdHandler struct {
	menuService services.IMenuService
	ussdService services.IUssdService
}

func NewUssdHandler(
	menuService services.IMenuService,
	ussdService services.IUssdService,
) *UssdHandler {
	return &UssdHandler{
		menuService: menuService,
		ussdService: ussdService,
	}
}

func (h *UssdHandler) Callback(c *fiber.Ctx) error {
	text := c.FormValue("text")

	if text == "" || text == "0" {
		main, err := h.menuService.GetAll()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		var mainData []string
		for i, s := range main {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			mainData = append(mainData, row)
		}
		justString := strings.Join(mainData, "\n")

		return c.Status(fiber.StatusOK).SendString(justString)

	} else {
		if !h.menuService.IsKeyPress(text) {
			return c.Status(fiber.StatusNotFound).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusNotFound,
					Message:    "menu_not_found",
				},
			)
		}
		menu, err := h.menuService.GetMenuByKeyPress(text)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		submenus, err := h.menuService.GetMenuByParentId(menu.GetId())
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		var subData []string
		for i, s := range submenus {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			subData = append(subData, row)
		}
		justString := strings.Join(subData, "\n")

		return c.Status(fiber.StatusOK).SendString(justString)
	}

}
