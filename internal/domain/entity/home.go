package entity

type Home struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	FixtureID int64    `json:"fixture_id"`
	Fixture   *Fixture `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	TeamID    int64    `json:"team_id"`
	Team      *Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	Goal      int      `gorm:"size:10" json:"goal"`
	IsWinner  bool     `gorm:"type:bool;default:false" json:"is_winner"`
}
