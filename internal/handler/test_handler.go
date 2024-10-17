package handler

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/providers/telco"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
)

type TestHandler struct {
	logger              *logger.Logger
	subscriptionService services.ISubscriptionService
}

func NewTestHandler(
	logger *logger.Logger,
	subscriptionService services.ISubscriptionService,
) *TestHandler {
	return &TestHandler{
		logger:              logger,
		subscriptionService: subscriptionService,
	}
}

func (h *TestHandler) TestBalance() {
	t := telco.NewTelco(h.logger, &entity.Service{
		UrlTelco: "http://localhost:9100/test/balance",
	}, &entity.Subscription{Msisdn: "2281299708787"}, utils.GenerateTrxId())
	resp, err := t.QueryProfileAndBal()
	if err != nil {
		log.Println(err.Error())
	}

	var respBal *model.QueryProfileAndBalResponse
	errX := xml.Unmarshal(resp, &respBal)
	if errX != nil {
		log.Printf("xml: unmarshal: %s", errX)
	}

	fmt.Println(respBal.Body.Item.BalDtoList.BalDto[1].Balance)
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

func (h *TestHandler) TestChargeFailed() {
	t := telco.NewTelco(h.logger, &entity.Service{
		UrlTelco: "http://localhost:9100/test/charge-failed",
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

	if respDeduct.IsFailed() {
		fmt.Println(respDeduct.Body.Fault)
	}
}

func (h *TestHandler) TestUpdateToFalse() {
	sub, err := h.subscriptionService.UpdateNotActive(&entity.Subscription{
		ServiceID: 7,
		Msisdn:    "22390869090",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(sub)
}
