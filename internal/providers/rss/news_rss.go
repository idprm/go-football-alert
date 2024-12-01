package rss

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type NewsRSS struct {
}

func NewNewsRSS() *NewsRSS {
	return &NewsRSS{}
}

func (p *NewsRSS) GetNews(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
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
