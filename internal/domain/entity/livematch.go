package entity

import (
	"net/url"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type LiveMatch struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	FixtureID   int64     `json:"fixture_id"`
	Fixture     *Fixture  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	FixtureDate time.Time `json:"fixture_date"`
	Goal        string    `json:"goal"`
	Elapsed     int       `json:"elapsed"`
	IsActive    bool      `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model
}

// Libya - Benin (0-0) 44"
func (e *LiveMatch) GetLiveMatchName() string {
	return e.Fixture.Home.GetName() + " - " + e.Fixture.Away.GetName() + " (" + e.GetGoal() + ") " + e.GetElapsed()
}

func (e *LiveMatch) GetLiveMatchNameQueryEscape() string {
	return url.QueryEscape(e.GetLiveMatchName())
}

func (e *LiveMatch) GetGoal() string {
	return e.Goal
}

func (e *LiveMatch) GetElapsed() string {
	return strconv.Itoa(e.Elapsed) + `"`
}
