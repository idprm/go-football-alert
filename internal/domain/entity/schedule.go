package entity

import "time"

type Schedule struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	CountryID  int       `json:"country_id"`
	Country    *Country  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	Name       string    `gorm:"size:45" json:"name"`
	PublishAt  time.Time `json:"publish_at"`
	UnlockedAt time.Time `json:"unlocked_at"`
	IsUnlocked bool      `gorm:"type:boolean;column:is_unlocked" json:"is_unlocked"`
}

func (e *Schedule) GetId() int {
	return e.ID
}

func (e *Schedule) GetName() string {
	return e.Name
}
