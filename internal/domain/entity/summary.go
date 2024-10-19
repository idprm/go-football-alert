package entity

import "time"

type Summary struct {
	ID                 int64     `gorm:"primaryKey" json:"id"`
	ServiceID          int       `json:"service_id"`
	Service            *Service  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	TotalSub           int       `gorm:"size:10:default:0" json:"total_sub"`
	TotalUnsub         int       `gorm:"size:10:default:0" json:"total_unsub"`
	TotalRenewal       int       `gorm:"size:10:default:0" json:"total_renewal"`
	TotalActiveSub     int       `gorm:"size:10:default:0" json:"total_active_sub,omitempty"`
	TotalChargeSuccess int       `gorm:"size:10:default:0" json:"total_charge_success"`
	TotalChargeFailed  int       `gorm:"size:10:default:0" json:"total_charge_failed"`
	TotalRevenue       float64   `gorm:"size:15:default:0" json:"total_revenue"`
	CreatedAt          time.Time `json:"created_at"`
}

func (e *Summary) SetServiceId(v int) {
	e.ServiceID = v
}

func (e *Summary) SetTotalSub(v int) {
	e.TotalSub = v
}

func (e *Summary) SetTotalUnsub(v int) {
	e.TotalUnsub = v
}

func (e *Summary) SetTotalRenewal(v int) {
	e.TotalRenewal = v
}

func (e *Summary) SetTotalChargeSuccess(v int) {
	e.TotalChargeSuccess = v
}

func (e *Summary) SetTotalChargeFailed(v int) {
	e.TotalChargeFailed = v
}

func (e *Summary) SetTotalRevenue(v float64) {
	e.TotalRevenue = v
}

func (e *Summary) GetCreatedAt() time.Time {
	return e.CreatedAt
}
