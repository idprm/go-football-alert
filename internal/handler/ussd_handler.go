package handler

import (
	"github.com/idprm/go-football-alert/internal/services"
)

type UssdHandler struct {
	ussdService services.IUssdService
}

func NewUssdHandler(
	ussdService services.IUssdService,
) *UssdHandler {
	return &UssdHandler{
		ussdService: ussdService,
	}
}
