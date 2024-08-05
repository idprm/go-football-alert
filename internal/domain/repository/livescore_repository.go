package repository

import "gorm.io/gorm"

type LiveScoreRepository struct {
	db *gorm.DB
}

func NewLivescoreRepository(db *gorm.DB) *LiveScoreRepository {
	return &LiveScoreRepository{
		db: db,
	}
}

type ILiveScoreRepository interface {
}
