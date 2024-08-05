package repository

import "gorm.io/gorm"

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

type ITeamRepository interface {
}
