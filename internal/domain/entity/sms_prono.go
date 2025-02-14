package entity

import "gorm.io/gorm"

type SMSProno struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	PronosticID    int64         `json:"pronostic_id"`
	Pronostic      *Pronostic    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pronostic,omitempty"`
	gorm.Model     `json:"-"`
}
