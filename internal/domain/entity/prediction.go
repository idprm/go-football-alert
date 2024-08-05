package entity

type Prediction struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	FixtureID int64    `json:"fixture_id"`
	Fixture   *Fixture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
}
