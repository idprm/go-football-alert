package rmcsport

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-football-alert/internal/utils"
)

var (
	URL_RMCSPORT string = utils.GetEnv("URL_RMCSPORT")
)

type RmcSport struct {
}

func NewRmcSport() *RmcSport {
	return &RmcSport{}
}

func (p *RmcSport) GetNews() ([]byte, error) {
	req, err := http.NewRequest("GET", URL_RMCSPORT, nil)
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
