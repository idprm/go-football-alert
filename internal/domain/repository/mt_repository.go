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

func (r *MTRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var mts []*entity.MT
	err := r.db.Where("UPPER(msisdn) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?) OR UPPER(content) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateMTs(mts, p, r.db)).Find(&mts).Error
	if err != nil {
		return nil, err
	}
	p.Rows = mts
	return p, nil
}

func (r *MTRepository) Save(c *entity.MT) (*entity.MT, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
