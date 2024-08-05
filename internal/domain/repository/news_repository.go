package repository

import "gorm.io/gorm"

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{
		db: db,
	}
}

type INewsRepository interface {
}
