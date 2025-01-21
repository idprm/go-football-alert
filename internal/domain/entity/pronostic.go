package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pronostic struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Category string    `gorm:"size:30" json:"category"`
	Value    string    `gorm:"size:250" json:"value"`
	StartAt  time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"start_at"`
	ExpireAt time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"expire_at"`
	IsSent   bool      `gorm:"type:boolean;default:false;column:is_sent" json:"is_sent"`
	IsActive bool      `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
	gorm.Model
}

func (e *Pronostic) GetCategory() string {
	return e.Category
}

func (e *Pronostic) GetValue() string {
	return e.Value
}

func (e *Pronostic) GetStartAt() time.Time {
	return e.StartAt
}

func (e *Pronostic) GetExpireAt() time.Time {
	return e.ExpireAt
}
