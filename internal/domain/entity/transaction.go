package entity

type Transaction struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	CountryID      int           `json:"country_id"`
	Country        *Country      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	ServiceID      int           `json:"service_id"`
	Service        *Service      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Msisdn         string        `gorm:"size:15;not null" json:"msisdn"`
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
