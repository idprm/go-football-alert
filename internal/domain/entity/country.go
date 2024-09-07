package entity

type Country struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:45" json:"name"`
	Code     string `gorm:"size:5" json:"code"`
	TimeZone string `gorm:"size:50" json:"timezone"`
	Currency string `gorm:"size:8" json:"currency"`
}
