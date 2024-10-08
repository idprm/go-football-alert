package entity

import "gorm.io/gorm"

type SubscriptionFollowLeague struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	LeagueID       int64         `json:"league_id"`
	League         *League       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	LimitPerDay    int           `json:"limit_by_day"`
	gorm.Model     `json:"-"`
}
