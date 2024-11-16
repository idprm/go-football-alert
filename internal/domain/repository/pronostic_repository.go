package repository

import (
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
	Count(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Pronostic, error)
	Save(*entity.Pronostic) (*entity.Pronostic, error)
	Update(*entity.Pronostic) (*entity.Pronostic, error)
	Delete(*entity.Pronostic) error
}

func (r *PronosticRepository) Count(fixtureId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Pronostic{}).Where("fixture_id = ?", fixtureId).Count(&count).Error
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

func (r *PronosticRepository) Get(fixtureId int) (*entity.Pronostic, error) {
	var c entity.Pronostic
	err := r.db.Where("fixture_id = ?", fixtureId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PronosticRepository) Save(c *entity.Pronostic) (*entity.Pronostic, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PronosticRepository) Update(c *entity.Pronostic) (*entity.Pronostic, error) {
	err := r.db.Where("fixture_id = ?", c.FixtureID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PronosticRepository) Delete(c *entity.Pronostic) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
