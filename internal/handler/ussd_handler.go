package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
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

	text := c.FormValue("text")

	layerFirst := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. Foot International \n
		2. Foot Europe \n
		3. Foot Afrique \n
		4. Coupe Nationale Guinee \n
		0. Prec \n
		98. Accueil
	`
	layerInternational := `Foot international Credit Goal Mercedi 24 Juil. \n
		1. CCgypte - Dominican Republic \n
		2. Guinea - Nouvelle Zeelande \n
		3. Japon - Paraguay \n
		4. Irag - Ukraine \n
		5. France - Etats Unis \n
		0. Suiv`

	layerEuro := `Foot Europe Credit Goal Mercedi 24 Juil. \n
		no match \n
		0. Prec`

	layerAfrique := `Foot Afrique Credit Goal Mercedi 24 Juil. \n
		no match \n
		0. Prec`

	layerGuinee := `Coupe Nationale Guinee Credit Goal Mercedi 24 Juil. \n
		no match \n
		0. Prec`

	layerMatch1_1 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. CCgypte \n
		2. Dominican Republic \n
		0. Prec \n
		98. Accueil
	`

	layerMatch1_2 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. Guinea \n
		2. Nouvelle Zeelande \n
		0. Prec \n
		98. Accueil
	`

	layerMatch1_3 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. Japon \n
		2. Paraguay \n
		0. Prec \n
		98. Accueil
	`
	layerMatch1_4 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. Irag \n
		2. Ukraine \n
		0. Prec \n
		98. Accueil
	`

	layerMatch1_5 := `Credit Goal Gagnez des lots a chaque but de votre equipe \n
		1. France \n
		2. Etats Unis \n
		0. Prec \n
		98. Accueil
	`

	switch text {
	case "1":
		return c.Status(fiber.StatusOK).SendString(layerInternational)
	case "2":
		return c.Status(fiber.StatusOK).SendString(layerEuro)
	case "3":
		return c.Status(fiber.StatusOK).SendString(layerAfrique)
	case "4":
		return c.Status(fiber.StatusOK).SendString(layerGuinee)
	case "0":
		return c.Status(fiber.StatusOK).SendString(layerFirst)
	case "1*1":
		return c.Status(fiber.StatusOK).SendString(layerMatch1_1)
	case "1*2":
		return c.Status(fiber.StatusOK).SendString(layerMatch1_2)
	case "1*3":
		return c.Status(fiber.StatusOK).SendString(layerMatch1_3)
	case "1*4":
		return c.Status(fiber.StatusOK).SendString(layerMatch1_4)
	case "1*5":
		return c.Status(fiber.StatusOK).SendString(layerMatch1_5)
	}
	return c.Status(fiber.StatusOK).SendString(layerFirst)
}

func (h *UssdHandler) Event(c *fiber.Ctx) error {
	log.Println(c.Body())
	return c.Status(fiber.StatusOK).SendString("OK")
}
