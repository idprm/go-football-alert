package entity

type Lineup struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	LeagueID  int64    `json:"league_id"`
	League    *League  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	FixtureID int64    `json:"fixture_id"`
	Fixture   *Fixture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
}
