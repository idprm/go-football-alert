package entity

import "gorm.io/gorm"

type FollowTeam struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	NewsID     int64 `json:"news_id"`
	News       *News `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	TeamID     int64 `json:"team_id"`
	Team       *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	gorm.Model `json:"-"`
}
