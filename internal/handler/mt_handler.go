package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
	"github.com/wiliehidayat87/rmqp"
)

type MTHandler struct {
	rmq    rmqp.AMQP
	logger *logger.Logger
	req    *model.MTRequest
}

func NewMTHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	req *model.MTRequest,
) *MTHandler {
	return &MTHandler{
		rmq:    rmq,
		logger: logger,
		req:    req,
	}
}

func (h *MTHandler) MessageTerminated() {
	k := kannel.NewKannel(h.logger, &entity.Service{}, h.req.Content, h.req.Subscription)
	k.SMS(h.req.Smsc)
}
