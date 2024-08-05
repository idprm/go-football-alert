package repository

import "gorm.io/gorm"

type PredictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) *PredictionRepository {
	return &PredictionRepository{
		db: db,
	}
}

type IPredictionRepository interface {
}
