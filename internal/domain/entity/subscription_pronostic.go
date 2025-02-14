package entity

import "gorm.io/gorm"

type SubscriptionPronostic struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	PronosticID    int64         `json:"pronostic_id"`
	Pronostic      *Pronostic    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pronostic,omitempty"`
	IsActive       bool          `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model
}
