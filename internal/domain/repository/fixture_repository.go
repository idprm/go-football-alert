package repository

import "gorm.io/gorm"

type FixtureRepository struct {
	db *gorm.DB
}

func NewFixtureRepository(db *gorm.DB) *FixtureRepository {
	return &FixtureRepository{
		db: db,
	}
}

type IFixtureRepository interface {
}
