package entity

import "gorm.io/gorm"

type Betting struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	gorm.Model `json:"-"`
}
