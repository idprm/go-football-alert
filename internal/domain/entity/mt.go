package entity

import "gorm.io/gorm"

type MT struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	TrxId      string `gorm:"size:100" json:"trx_id,omitempty"`
	Msisdn     string `gorm:"size:15;not null" json:"msisdn"`
	Keyword    string `gorm:"size:250" json:"keyword"`
	Content    string `gorm:"type:text" json:"content"`
	StatusCode int    `gorm:"size:10" json:"status_code"`
	StatusText string `gorm:"size:100" json:"status_text"`
	Payload    string `gorm:"type:text" json:"payload"`
	gorm.Model
}
