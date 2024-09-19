package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type FixtureRepository struct {
	db *gorm.DB
}

func NewFixtureRepository(db *gorm.DB) *FixtureRepository {
	return &FixtureRepository{
		db: db,
	}
}

type IFixtureRepository interface {
	Count(int, int) (int64, error)
	CountByPrimaryId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Fixture, error)
	Save(*entity.Fixture) (*entity.Fixture, error)
	Update(*entity.Fixture) (*entity.Fixture, error)
	UpdateByPrimaryId(*entity.Fixture) (*entity.Fixture, error)
	Delete(*entity.Fixture) error
}

func (r *FixtureRepository) Count(homeId, awayId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Fixture{}).Where("home_id = ?", homeId).Where("away_id = ?", awayId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *FixtureRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Fixture{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *FixtureRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var fixtures []*entity.Fixture
	err := r.db.Scopes(Paginate(fixtures, pagination, r.db)).Preload("Home").Preload("Away").Find(&fixtures).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = fixtures
	return pagination, nil
}

func (r *FixtureRepository) Get(homeId, awayId int) (*entity.Fixture, error) {
	var c entity.Fixture
	err := r.db.Where("home_id = ?", homeId).Where("away_id = ?", awayId).Preload("Home").Preload("Away").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FixtureRepository) Save(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) Update(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) UpdateByPrimaryId(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Where("primary_id = ?", c.PrimaryID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) Delete(c *entity.Fixture) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
