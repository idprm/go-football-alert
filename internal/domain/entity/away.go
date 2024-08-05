package entity

type Away struct {
	ID        int64    `json:"id,omitempty"`
	FixtureID int64    `json:"fixture_id"`
	Fixture   *Fixture `json:"fixture"`
	TeamID    int64    `json:"team_id"`
	Team      *Team    `json:"team"`
	Goal      int      `json:"goal"`
	IsWinner  bool     `gorm:"type:bool;default:false" json:"is_winner"`
}
