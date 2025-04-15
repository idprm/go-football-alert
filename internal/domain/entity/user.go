package entity

import "gorm.io/gorm"

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"size:100" json:"email"`
	Password string `gorm:"size:200" json:"password"`
	gorm.Model
}
