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
}

func NewKannel(
	logger *logger.Logger,
	service *entity.Service,
	content *entity.Content,
	subscription *entity.Subscription,
) *Kannel {
	return &Kannel{
		logger:       logger,
		service:      service,
		content:      content,
		subscription: subscription,
	}
}

type IKannel interface {
	SMS() ([]byte, error)
}

func (p *Kannel) SMS() ([]byte, error) {
	l := p.logger.Init("mt", true)
	start := time.Now()

	req, err := http.NewRequest(http.MethodGet, p.service.GetUrlMT(), nil)
	if err != nil {
		return nil, err
	}

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
	l.WithFields(logrus.Fields{"request": req}).Info("SMS")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	duration := time.Since(start).Milliseconds()
	p.logger.Writer(string(body))
	l.WithFields(logrus.Fields{
		"msisdn":   p.subscription.GetMsisdn(),
		"response": string(body),
		"duration": duration,
	}).Info("SMS")

	return body, nil
}