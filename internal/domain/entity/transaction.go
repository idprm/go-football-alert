package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID             int64    `gorm:"primaryKey" json:"id"`
	TrxId          string   `gorm:"size:100" json:"trx_id,omitempty"`
	ServiceID      int      `json:"service_id"`
	Service        *Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Msisdn         string   `gorm:"size:15;not null" json:"msisdn"`
	Code           string   `gorm:"size:25;not null" json:"code"`
	Channel        string   `gorm:"size:15" json:"channel,omitempty"`
	Keyword        string   `gorm:"size:50" json:"keyword"`
	Amount         float64  `gorm:"size:10;default:0" json:"amount"`
	Discount       float64  `gorm:"size:3;default:0" json:"discount"`
	Status         string   `gorm:"size:25" json:"status,omitempty"`
	StatusCode     string   `gorm:"size:85" json:"status_code,omitempty"`
	StatusDetail   string   `gorm:"size:85" json:"status_detail,omitempty"`
	Subject        string   `gorm:"size:25" json:"subject,omitempty"`
	IpAddress      string   `gorm:"size:30" json:"ip_address,omitempty"`
	Payload        string   `gorm:"type:text" json:"payload,omitempty"`
	Note           string   `gorm:"type:text" json:"note,omitempty"`
	IsDiscount     bool     `gorm:"type:boolean;default:false;column:is_discount" json:"is_discount,omitempty"`
	IsUnderpayment bool     `gorm:"type:boolean;default:false;column:is_underpayment" json:"is_underpayment,omitempty"`
	gorm.Model
}

func (e *Transaction) GetId() int64 {
	return e.ID
}

func (e *Transaction) GetServiceId() int {
	return e.ServiceID
}

func (e *Transaction) GetMsisdn() string {
	return e.Msisdn
}

func (e *Transaction) GetCode() string {
	return e.Code
}

func (t *Transaction) SetAmount(v float64) {
	t.Amount = v
}

func (t *Transaction) SetDiscount(v float64) {
	t.Discount = v
}

func (t *Transaction) SetStatus(v string) {
	t.Status = v
}

func (t *Transaction) SetStatusCode(v string) {
	t.StatusCode = v
}

func (t *Transaction) SetStatusDetail(v string) {
	t.StatusDetail = v
}

func (t *Transaction) SetSubject(v string) {
	t.Subject = v
}

func (t *Transaction) SetNote(v string) {
	t.Note = v
}
