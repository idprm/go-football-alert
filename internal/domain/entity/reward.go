package entity

type Reward struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	FixtureID      int64         `json:"fixture_id"`
	Fixture        *Fixture      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	Msisdn         string        `gorm:"size:15;not null" json:"msisdn"`
	Amount         float64       `gorm:"size:8;default:0" json:"amount"`
}

func (e *Reward) GetId() int64 {
	return e.ID
}

func (e *Reward) GetFixtureId() int64 {
	return e.FixtureID
}

func (e *Reward) GetSubscriptionId() int64 {
	return e.SubscriptionID
}
