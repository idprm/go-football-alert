package repository

import "gorm.io/gorm"

type LeagueRepository struct {
	db *gorm.DB
}

func NewLeagueRepository(db *gorm.DB) *LeagueRepository {
	return &LeagueRepository{
		db: db,
	}
}

type ILeagueRepository interface {
}
