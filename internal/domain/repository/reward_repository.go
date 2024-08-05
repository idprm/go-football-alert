package repository

import "gorm.io/gorm"

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{
		db: db,
	}
}

type IRewardRepository interface {
}
