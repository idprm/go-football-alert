package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type MTRepository struct {
	db *gorm.DB
}

func NewMTRepository(db *gorm.DB) *MTRepository {
	return &MTRepository{
		db: db,
	}
}

type IMTRepository interface {
	Save(*entity.MT) (*entity.MT, error)
}

func (r *MTRepository) Save(c *entity.MT) (*entity.MT, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
