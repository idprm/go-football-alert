package repository

import "gorm.io/gorm"

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

type ISubscriptionRepository interface {
}
