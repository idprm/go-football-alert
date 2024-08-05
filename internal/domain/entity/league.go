package entity

type League struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
