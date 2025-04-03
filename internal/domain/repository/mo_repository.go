package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type MORepository struct {
	db *gorm.DB
}

func NewMORepository(db *gorm.DB) *MORepository {
	return &MORepository{
		db: db,
	}
}

type IMORepository interface {
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.MO) (*entity.MO, error)
}

func (r *MORepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var mos []*entity.MO
	err := r.db.Where("UPPER(msisdn) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateMTs(mos, p, r.db)).Find(&mos).Error
	if err != nil {
		return nil, err
	}
	p.Rows = mos
	return p, nil
}

func (r *MORepository) Save(c *entity.MO) (*entity.MO, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
