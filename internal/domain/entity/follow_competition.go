package entity

import "gorm.io/gorm"

type FollowCompetition struct {
	ID         int64   `gorm:"primaryKey" json:"id"`
	NewsID     int64   `json:"news_id"`
	News       *News   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	LeagueID   int64   `json:"league_id"`
	League     *League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	gorm.Model `json:"-"`
}
