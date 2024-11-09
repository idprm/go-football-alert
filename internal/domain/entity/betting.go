package entity

import "gorm.io/gorm"

type Betting struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	FixtureID      int64         `json:"fixture_id"`
	Fixture        *Fixture      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	Bet            float64       `gorm:"size:8" json:"bet"`
	Profit         float64       `gorm:"size:8" json:"profit"`
	gorm.Model
}
