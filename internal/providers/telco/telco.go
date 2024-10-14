package telco

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/sirupsen/logrus"
)

type Telco struct {
	logger       *logger.Logger
	service      *entity.Service
	subscription *entity.Subscription
	trxId        string
}

func NewTelco(
	logger *logger.Logger,
	service *entity.Service,
	subscription *entity.Subscription,
	trxId string,
) *Telco {
	return &Telco{
		logger:       logger,
		service:      service,
		subscription: subscription,
		trxId:        trxId,
	}
}

type ITelco interface {
	QueryProfileAndBal() ([]byte, error)
	DeductFee() ([]byte, error)
}

func (p *Telco) QueryProfileAndBal() ([]byte, error) {
	l := p.logger.Init("mt", true)
	start := time.Now()

	var reqXml model.QueryProfileAndBalRequest
	reqXml.SetSoap("http://schemas.xmlsoap.org/soap/envelope/")
	reqXml.SetXsd("http://com.ztesoft.zsmart/xsd")
	reqXml.SetUsername(p.service.GetUserTelco())
	reqXml.SetPassword(p.service.GetPassTelco())
	reqXml.SetMsisdn(p.subscription.GetMsisdn())
	reqXml.SetTransactionSN(time.Now().UnixMicro())

	payload, err := xml.Marshal(&reqXml)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, p.service.GetUrlTelco(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	p.logger.Writer(req)

	l.WithFields(logrus.Fields{
		"trx_id":       p.trxId,
		"request_url":  p.service.GetUrlTelco(),
		"request_body": req.Body,
	}).Info("PROFILE_AND_BAL")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	p.logger.Writer(string(body))
	duration := time.Since(start).Milliseconds()

	l.WithFields(
		logrus.Fields{
			"trx_id":      p.trxId,
			"msisdn":      p.subscription.GetMsisdn(),
			"response":    string(body),
			"status_code": resp.StatusCode,
			"status_text": http.StatusText(resp.StatusCode),
			"duration":    duration,
		}).Info("PROFILE_AND_BAL")

	return body, nil
}

func (p *Telco) DeductFee() ([]byte, error) {
	// l := p.logger.Init("mt", true)
	// start := time.Now()

	var reqXml model.DeductRequest

	reqXml.SetSoap("http://schemas.xmlsoap.org/soap/envelope/")
	reqXml.SetXsd("http://com.ztesoft.zsmart/xsd")
	reqXml.SetUsername(p.service.GetUserTelco())
	reqXml.SetPassword(p.service.GetPassTelco())
	reqXml.SetTransactionSN(time.Now().UnixMicro())
	reqXml.SetTransactionDesc("OFCTEST")
	reqXml.SetChannelID("ESERV")
	reqXml.SetMsisdn(p.subscription.GetMsisdn())
	reqXml.SetAccountCode("")
	reqXml.SetAcctResCode("1")
	reqXml.SetDeductBalance(strconv.FormatFloat(p.service.GetPrice(), 'f', 0, 64))

	payload, err := xml.Marshal(&reqXml)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, p.service.GetUrlTelco(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	// p.logger.Writer(req)
	// l.WithFields(logrus.Fields{
	// 	"trx_id":       p.trxId,
	// 	"request_url":  p.service.GetUrlTelco(),
	// 	"request_body": req.Body,
	// }).Info("DEDUCT_FEE")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// p.logger.Writer(string(body))
	// duration := time.Since(start).Milliseconds()
	// l.WithFields(
	// 	logrus.Fields{
	// 		"trx_id":      p.trxId,
	// 		"msisdn":      p.subscription.GetMsisdn(),
	// 		"response":    string(body),
	// 		"status_code": resp.StatusCode,
	// 		"status_text": http.StatusText(resp.StatusCode),
	// 		"duration":    duration,
	// 	}).Info("DEDUCT_FEE")

	return body, nil
}
