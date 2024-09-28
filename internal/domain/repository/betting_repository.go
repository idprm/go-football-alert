package repository

import "gorm.io/gorm"

type BettingRepository struct {
	db *gorm.DB
}

func NewBettingRepository(db *gorm.DB) *BettingRepository {
	return &BettingRepository{
		db: db,
	}
}

type IBettingRepository interface {
}
