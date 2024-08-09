package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type HomeRepository struct {
	db *gorm.DB
}

func NewHomeRepository(db *gorm.DB) *HomeRepository {
	return &HomeRepository{
		db: db,
	}
}

type IHomeRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Home, error)
	Save(*entity.Home) (*entity.Home, error)
	Update(*entity.Home) (*entity.Home, error)
	Delete(*entity.Home) error
}

func (r *HomeRepository) Count(fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Home{}).Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *HomeRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var homes []*entity.Home
	err := r.db.Scopes(Paginate(homes, pagination, r.db)).Preload("Fixture").Preload("Team").Find(&homes).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = homes
	return pagination, nil
}

func (r *HomeRepository) Get(fixtureId, teamId int) (*entity.Home, error) {
	var c entity.Home
	err := r.db.Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Preload("Fixture").Preload("Team").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *HomeRepository) Save(c *entity.Home) (*entity.Home, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *HomeRepository) Update(c *entity.Home) (*entity.Home, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *HomeRepository) Delete(c *entity.Home) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
