package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type PredictWinHandler struct {
	rmq            rmqp.AMQP
	logger         *logger.Logger
	sub            *entity.Subscription
	serviceService services.IServiceService
}

func NewPredictWinHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	sub *entity.Subscription,
	serviceService services.IServiceService,
) *PredictWinHandler {
	return &PredictWinHandler{
		rmq:            rmq,
		logger:         logger,
		sub:            sub,
		serviceService: serviceService,
	}
}
