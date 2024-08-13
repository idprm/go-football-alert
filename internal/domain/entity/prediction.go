package entity

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
