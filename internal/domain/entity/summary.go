package entity

import "time"

type SummaryDashboard struct {
	ID             int64     `gorm:"primaryKey" json:"id"`
	TotalActiveSub int       `gorm:"size:10:default:0" json:"total_active_sub"`
	TotalRevenue   float64   `gorm:"size:15:default:0" json:"total_revenue"`
	CreatedAt      time.Time `json:"created_at"`
}

func (e *SummaryDashboard) SetTotalActiveSub(v int) {
	e.TotalActiveSub = v
}

func (e *SummaryDashboard) SetTotalRevenue(v float64) {
	e.TotalRevenue = v
}

func (e *SummaryDashboard) GetCreatedAt() time.Time {
	return e.CreatedAt
}

type SummaryRevenue struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Subject   string    `gorm:"size:45" json:"subject"`
	Status    string    `gorm:"size:45" json:"status"`
	Total     int       `gorm:"size:10:default:0" json:"total"`
	Revenue   float64   `gorm:"size:15:default:0" json:"revenue"`
	CreatedAt time.Time `json:"created_at"`
}
