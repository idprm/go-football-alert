package entity

type Transaction struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	ServiceID      int           `json:"service_id"`
	Service        *Service      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Msisdn         string        `gorm:"size:15;not null" json:"msisdn"`
}
