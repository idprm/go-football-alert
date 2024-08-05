package repository

import "gorm.io/gorm"

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

type IScheduleRepository interface {
}
