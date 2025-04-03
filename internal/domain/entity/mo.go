package entity

import "gorm.io/gorm"

type MO struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	TrxId   string `gorm:"size:100" json:"trx_id,omitempty"`
	Msisdn  string `gorm:"size:15;not null" json:"msisdn"`
	Channel string `gorm:"size:15" json:"channel,omitempty"`
	Keyword string `gorm:"size:300" json:"keyword,omitempty"`
	Action  string `gorm:"size:15" json:"action,omitempty"`
	gorm.Model
}
