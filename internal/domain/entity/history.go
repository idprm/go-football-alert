package entity

import "time"

type History struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	ServiceID      int           `json:"service_id"`
	Service        *Service      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Msisdn         string        `gorm:"size:15;not null" json:"msisdn"`
	Keyword        string        `json:"keyword"`
	Subject        string        `json:"subject"`
	Status         string        `json:"status"`
	IpAddress      string        `json:"ip_address"`
	CreatedAt      time.Time     `json:"created_at"`
}

func (e *History) GetId() int64 {
	return e.ID
}

func (e *History) GetSubscriptionId() int64 {
	return e.SubscriptionID
}

func (e *History) GetServiceId() int {
	return e.ServiceID
}

func (e *History) GetMsisdn() string {
	return e.Msisdn
}
