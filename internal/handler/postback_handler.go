package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type PostbackHandler struct {
	rmq            rmqp.AMQP
	logger         *logger.Logger
	sub            *entity.Subscription
	serviceService services.IServiceService
}

func NewPostbackHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
) *PostbackHandler {
	return &PostbackHandler{
		rmq:            rmq,
		logger:         logger,
		sub:            sub,
		serviceService: serviceService,
	}
}

func (h *PostbackHandler) MO() {

}

func (h *PostbackHandler) DN() {

}
