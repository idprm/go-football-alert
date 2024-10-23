package entity

import "gorm.io/gorm"

type SubscriptionFollowLeague struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	LeagueID       int64         `json:"league_id"`
	League         *League       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	LimitPerDay    int           `gorm:"size:3;default:4" json:"limit_by_day"`
	Sent           int           `gorm:"size:3;default:0" json:"sent"`
	IsActive       bool          `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model     `json:"-"`
}
