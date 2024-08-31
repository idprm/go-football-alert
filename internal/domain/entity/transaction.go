package entity

import (
	"time"
)

type Transaction struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	TrxId          string        `json:"trx_id,omitempty"`
	CountryID      int           `json:"country_id"`
	Country        *Country      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	ServiceID      int           `json:"service_id"`
	Service        *Service      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Msisdn         string        `gorm:"size:15;not null" json:"msisdn"`
	Keyword        string        `json:"keyword,omitempty"`
	Amount         float64       `json:"amount,omitempty"`
	Status         string        `json:"status,omitempty"`
	StatusCode     string        `json:"status_code,omitempty"`
	StatusDetail   string        `json:"status_detail,omitempty"`
	Subject        string        `json:"subject,omitempty"`
	IpAddress      string        `json:"ip_address,omitempty"`
	Payload        string        `json:"payload,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

func (e *Transaction) GetId() int64 {
	return e.ID
}

func (e *Transaction) GetSubscriptionId() int64 {
	return e.SubscriptionID
}

func (e *Transaction) GetServiceId() int {
	return e.ServiceID
}

func (e *Transaction) GetMsisdn() string {
	return e.Msisdn
}

func (t *Transaction) SetAmount(v float64) {
	t.Amount = v
}

func (t *Transaction) SetStatus(v string) {
	t.Status = v
}

func (t *Transaction) SetStatusCode(v string) {
	t.StatusCode = v
}

func (t *Transaction) SetStatusDetail(v string) {
	t.StatusDetail = v
}

func (t *Transaction) SetSubject(v string) {
	t.Subject = v
}
