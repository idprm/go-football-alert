package entity

import (
	"time"

	"gorm.io/gorm"
)

type Lineup struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	LeagueID    int64     `json:"league_id"`
	League      *League   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	FixtureID   int64     `json:"fixture_id"`
	Fixture     *Fixture  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	TeamID      int64     `json:"team_id"`
	Team        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	TeamName    string    `json:"team_name,omitempty"`
	FixtureDate time.Time `json:"fixture_date"`
	Formation   string    `gorm:"size:30" json:"formation,omitempty"`
	gorm.Model  `json:"-"`
}
