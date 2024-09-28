package repository

import "gorm.io/gorm"

type FollowTeamRepository struct {
	db *gorm.DB
}

func NewFollowTeamRepository(db *gorm.DB) *FollowTeamRepository {
	return &FollowTeamRepository{
		db: db,
	}
}

type IFollowTeamRepository interface {
}
