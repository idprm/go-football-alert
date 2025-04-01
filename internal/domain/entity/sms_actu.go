package entity

import "gorm.io/gorm"

type SMSActu struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Msisdn     string `json:"msisdn"`
	NewsID     int64  `json:"news_id"`
	News       *News  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	gorm.Model `json:"-"`
}
