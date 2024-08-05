package repository

import "gorm.io/gorm"

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{
		db: db,
	}
}

type ISeasonRepository interface {
}
