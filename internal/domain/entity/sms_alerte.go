package entity

import "gorm.io/gorm"

type SMSAlerte struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	NewsID         int64         `json:"news_id"`
	News           *News         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	gorm.Model     `json:"-"`
}
