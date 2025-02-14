package kannel

import (
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/sirupsen/logrus"
)

type Kannel struct {
	logger       *logger.Logger
	service      *entity.Service
	content      *entity.Content
	subscription *entity.Subscription
	trxId        string
}

func NewKannel(
	logger *logger.Logger,
	service *entity.Service,
	content *entity.Content,
	subscription *entity.Subscription,
	trxId string,
) *Kannel {
	return &Kannel{
		logger:       logger,
		service:      service,
		content:      content,
		subscription: subscription,
		trxId:        trxId,
	}
}

type IKannel interface {
	SMS(string) ([]byte, error)
}

func (p *Kannel) SMS(sc string) (int, []byte, error) {
	l := p.logger.Init("mt", true)

	start := time.Now()
	p.service.SetUrlMT(
		"MOBIMIUM",
		p.service.UserMT,
		p.service.PassMT,
		sc,
		p.subscription.GetMsisdn(),
		p.content.GetValue(),
	)

	req, err := http.NewRequest(http.MethodGet, p.service.GetUrlMT(), nil)
	if err != nil {
		return 0, nil, err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: tr,
	}

	p.logger.Writer(req)
	l.WithFields(logrus.Fields{
		"trx_id":  p.trxId,
		"msisdn":  p.subscription.GetMsisdn(),
		"request": p.service.GetUrlMT(),
	}).Info("SMS")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	p.logger.Writer(string(body))

	duration := time.Since(start).Milliseconds()
	p.logger.Writer(string(body))
	l.WithFields(logrus.Fields{
		"trx_id":      p.trxId,
		"msisdn":      p.subscription.GetMsisdn(),
		"response":    string(body),
		"status_code": resp.StatusCode,
		"status_text": http.StatusText(resp.StatusCode),
		"duration":    duration,
	}).Info("SMS")

	return resp.StatusCode, body, nil
}
