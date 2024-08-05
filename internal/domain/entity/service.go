package entity

type Service struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
