package madeinfoot

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-football-alert/internal/utils"
)

var (
	URL_MADEINFOOT string = utils.GetEnv("URL_MADEINFOOT")
)

type MadeInFoot struct {
}

func NewMadeInFoot() *MadeInFoot {
	return &MadeInFoot{}
}

func (p *MadeInFoot) GetNews() ([]byte, error) {
	req, err := http.NewRequest("GET", URL_MADEINFOOT, nil)
	if err != nil {
		return nil, errors.New(err.Error())
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
