package repository

import "gorm.io/gorm"

type HomeRepository struct {
	db *gorm.DB
}

func NewHomeRepository(db *gorm.DB) *HomeRepository {
	return &HomeRepository{
		db: db,
	}
}

type IHomeRepository interface {
}
