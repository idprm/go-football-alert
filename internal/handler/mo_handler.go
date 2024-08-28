package handler

import (
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/wiliehidayat87/rmqp"
)

type MOHandler struct {
	rmq    rmqp.AMQP
	logger *logger.Logger
}

func NewMOHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
) *MOHandler {
	return &MOHandler{
		rmq:    rmq,
		logger: logger,
	}
}

func (h *MOHandler) Firstpush() {

}

func (h *MOHandler) Unsub() {

}
