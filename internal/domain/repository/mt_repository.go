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
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.MT) (*entity.MT, error)
}

func (r *MTRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var mts []*entity.MT
	err := r.db.Scopes(Paginate(mts, pagination, r.db)).Find(&mts).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = mts
	return pagination, nil
}

func (r *MTRepository) Save(c *entity.MT) (*entity.MT, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
