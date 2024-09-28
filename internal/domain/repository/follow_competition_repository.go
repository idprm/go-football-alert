package repository

import "gorm.io/gorm"

type FollowCompetitionRepository struct {
	db *gorm.DB
}

func NewFollowCompetitionRepository(db *gorm.DB) *FollowCompetitionRepository {
	return &FollowCompetitionRepository{
		db: db,
	}
}

type IFollowCompetitionRepository interface {
}
