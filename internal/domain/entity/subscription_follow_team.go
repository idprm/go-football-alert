package entity

import (
	"time"

	"gorm.io/gorm"
)

type SubscriptionFollowTeam struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	TeamID         int64         `json:"team_id"`
	Team           *Team         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	LimitPerDay    int           `gorm:"size:3;default:4" json:"limit_by_day"`
	Sent           int           `gorm:"size:3;default:0" json:"sent"`
	LatestKeyword  string        `gorm:"size:100" json:"latest_keyword,omitempty"`
	RenewalAt      time.Time     `gorm:"type:TIMESTAMP;null;default:null" json:"renewal_at,omitempty"`
	UnsubAt        time.Time     `gorm:"type:TIMESTAMP;null;default:null" json:"unsub_at,omitempty"`
	ChargeAt       time.Time     `gorm:"type:TIMESTAMP;null;default:null" json:"charge_at,omitempty"`
	RetryAt        time.Time     `gorm:"type:TIMESTAMP;null;default:null" json:"retry_at,omitempty"`
	IsActive       bool          `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model
}
