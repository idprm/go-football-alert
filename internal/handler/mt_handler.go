package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/kannel"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/wiliehidayat87/rmqp"
)

type MTHandler struct {
	rmq       rmqp.AMQP
	logger    *logger.Logger
	mtService services.IMTService
	req       *model.MTRequest
}

func NewMTHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	mtService services.IMTService,
	req *model.MTRequest,
) *MTHandler {
	return &MTHandler{
		rmq:       rmq,
		logger:    logger,
		mtService: mtService,
		req:       req,
	}
}

func (h *MTHandler) MessageTerminated() {

	// is time
	if h.IsTime() {
		if h.req.Content != nil {
			k := kannel.NewKannel(
				h.logger,
				h.req.Service,
				h.req.Content,
				h.req.Subscription,
				h.req.TrxId,
			)
			statusCode, sms, err := k.SMS(h.req.Smsc)
			if err != nil {
				log.Println(err.Error())
			}
			h.mtService.Save(
				&entity.MT{
					TrxId:      h.req.TrxId,
					Msisdn:     h.req.Subscription.GetMsisdn(),
					Keyword:    h.req.Keyword,
					Content:    h.req.Content.GetValue(),
					StatusCode: statusCode,
					StatusText: http.StatusText(statusCode),
					Payload:    string(sms),
				},
			)
		}
	} else {
		h.mtService.Save(
			&entity.MT{
				TrxId:      h.req.TrxId,
				Msisdn:     h.req.Subscription.GetMsisdn(),
				Keyword:    h.req.Keyword,
				Content:    h.req.Content.GetValue(),
				StatusCode: http.StatusNoContent,
				StatusText: http.StatusText(http.StatusNoContent),
				Payload:    "FROM_9AM_TO_10PM_NO_CONTENT_SEND",
			},
		)
	}
}

func (h *MTHandler) IsTime() bool {
	// from 9am to 10pm
	start, _ := time.Parse("15:04", "08:00")
	end, _ := time.Parse("15:04", "22:30")

	t, err := time.Parse("15:04", time.Now().Format("15:04"))
	if err != nil {
		log.Fatal(err)
	}
	return start.Before(t) && end.After(t)
}
