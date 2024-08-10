package entity

type Service struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:50;not null" json:"name"`
	Code string `gorm:"size:15;not null" json:"code"`
}
