package handler

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/telco"
	"github.com/idprm/go-football-alert/internal/utils"
)

type TestHandler struct {
	logger *logger.Logger
}

func NewTestHandler(
	logger *logger.Logger,
) *TestHandler {
	return &TestHandler{
		logger: logger,
	}
}

func (h *TestHandler) TestCharge() {
	t := telco.NewTelco(h.logger, &entity.Service{
		UrlTelco: "http://localhost:9100/test/charge",
	}, &entity.Subscription{Msisdn: "2281299708787"}, utils.GenerateTrxId())
	resp, err := t.DeductFee()
	if err != nil {
		log.Println(err.Error())
	}

	var respDeduct *model.DeductResponse
	errX := xml.Unmarshal(resp, &respDeduct)
	if errX != nil {
		log.Printf("xml: unmarshal: %s", errX)
	}

	fmt.Println(respDeduct.Body.Item.TransactionSN)
}
