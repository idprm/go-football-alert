package maxifoot

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-football-alert/internal/utils"
)

var (
	URL_MAXFOOT string = utils.GetEnv("URL_MAXFOOT")
)

type Maxifoot struct {
}

func NewMaxifoot() *Maxifoot {
	return &Maxifoot{}
}

func (p *Maxifoot) GetNews() ([]byte, error) {
	req, err := http.NewRequest("GET", URL_MAXFOOT, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    20 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   20 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}
