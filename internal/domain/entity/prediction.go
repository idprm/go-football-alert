package entity

import "time"

type Prediction struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	FixtureID int64    `json:"fixture_id"`
	Fixture   *Fixture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
}

func (e *Prediction) GetId() int64 {
	return e.ID
}

func (e *Prediction) GetFixtureId() int64 {
	return e.FixtureID
}

type PredictionSubsciption struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	TeamID         int64         `json:"team_id"`
	Team           *Team         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
}
