package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type PronosticRepository struct {
	db *gorm.DB
}

func NewPronosticRepository(db *gorm.DB) *PronosticRepository {
	return &PronosticRepository{
		db: db,
	}
}

type IPronosticRepository interface {
	CountByStartAt(time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(time.Time) (*entity.Pronostic, error)
	GetById(int64) (*entity.Pronostic, error)
	Save(*entity.Pronostic) error
	Update(*entity.Pronostic) error
	Delete(*entity.Pronostic) error
}

func (r *PronosticRepository) CountByStartAt(startAt time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Pronostic{}).Where("start_at = ?", startAt).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *PronosticRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var pronostics []*entity.Pronostic
	err := r.db.Scopes(Paginate(pronostics, pagination, r.db)).Find(&pronostics).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = pronostics
	return pagination, nil
}

func (r *PronosticRepository) Get(start time.Time) (*entity.Pronostic, error) {
	var c entity.Pronostic
	err := r.db.Where("start_at = ?", start).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PronosticRepository) GetById(id int64) (*entity.Pronostic, error) {
	var c entity.Pronostic
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PronosticRepository) Save(c *entity.Pronostic) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PronosticRepository) Update(c *entity.Pronostic) error {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PronosticRepository) Delete(c *entity.Pronostic) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
