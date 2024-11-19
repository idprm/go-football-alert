package entity

import (
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
