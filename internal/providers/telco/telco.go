package telco

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Telco struct {
}

func NewTelco() *Telco {
	return &Telco{}
}

type ITelco interface {
}

func (p *Telco) Token() ([]byte, error) {
	dataJson, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(dataJson))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	timeStamp := strconv.Itoa(int(now.Unix()))

	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("X-Api-Key", "")
	req.Header.Add("X-Signature", "")
	req.Header.Add("X-Trxtime", timeStamp)

	tr := &http.Transport{
		MaxIdleConns:       40,
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
