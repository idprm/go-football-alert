package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
)

type MOHandler struct {
	logger    *logger.Logger
	moService services.IMOService
	req       *entity.MO
}

func NewMOHandler(
	logger *logger.Logger,
	moService services.IMOService,
	req *entity.MO,
) *MOHandler {
	return &MOHandler{
		logger:    logger,
		moService: moService,
		req:       req,
	}
}

func (h *MOHandler) Insert() {
	h.moService.Save(
		&entity.MO{
			TrxId:   h.req.TrxId,
			Msisdn:  h.req.Msisdn,
			Channel: h.req.Channel,
			Keyword: h.req.Keyword,
			Action:  h.req.Action,
		},
	)
}
