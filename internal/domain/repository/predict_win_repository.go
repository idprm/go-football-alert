package repository

import "gorm.io/gorm"

type PredictWinRepository struct {
	db *gorm.DB
}

func NewPredictWinRepository(db *gorm.DB) *PredictWinRepository {
	return &PredictWinRepository{
		db: db,
	}
}

type IPredictWinRepository interface {
}
