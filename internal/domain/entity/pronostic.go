package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pronostic struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FixtureID int64     `json:"fixture_id"`
	Fixture   *Fixture  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	Category  string    `gorm:"size:30" json:"category"`
	Value     string    `gorm:"size:250" json:"value"`
	PublishAt time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"publish_at"`
	IsSent    bool      `gorm:"type:boolean;default:false;column:is_sent" json:"is_sent"`
	gorm.Model
}
