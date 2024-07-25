package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/services"
)

type UssdHandler struct {
	ussdService services.IUssdService
}

func NewUssdHandler(ussdService services.IUssdService) *UssdHandler {
	return &UssdHandler{
		ussdService: ussdService,
	}
}

func (h *UssdHandler) Callback(c *fiber.Ctx) error {
	req := new(model.AfricasTalkingRequest)

	layer1 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. Foot International \n
		2. Foot Europe \n
		3. Foot Afrique \n
		4. Coupe Nationale Guinee \n
		0. Prec \n
		98. Accueil
	`
	layer2 := `Foot international Credit Goal Mercedi 24 Juil. \n
		1. CCgypte - Dominican Republic \n
		2. Guinea - Nouvelle Zeelande \n
		3. Japon - Paraguay \n
		4. Irag - Ukraine \n
		5. France - Etats Unis \n
		0. Suiv`

	if req.IsFirst() {
		return c.Status(fiber.StatusOK).SendString(layer1)
	}
	if req.IsOne() {
		return c.Status(fiber.StatusOK).SendString(layer2)
	}
	if req.IsTwo() {
		return c.Status(fiber.StatusOK).SendString(layer2)
	}
	return c.Status(fiber.StatusOK).SendString(layer1)
}

func (h *UssdHandler) Event(c *fiber.Ctx) error {
	log.Println(c.Body())
	return c.Status(fiber.StatusOK).SendString("OK")
}
