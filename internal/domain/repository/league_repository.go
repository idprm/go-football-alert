package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type LeagueRepository struct {
	db *gorm.DB
}

func NewLeagueRepository(db *gorm.DB) *LeagueRepository {
	return &LeagueRepository{
		db: db,
	}
}

type ILeagueRepository interface {
	Count(string) (int64, error)
	CountByPrimaryId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.League, error)
	GetByPrimaryId(int) (*entity.League, error)
	Save(*entity.League) (*entity.League, error)
	Update(*entity.League) (*entity.League, error)
	Delete(*entity.League) error
}

func (r *LeagueRepository) Count(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var leagues []*entity.League
	err := r.db.Scopes(Paginate(leagues, pagination, r.db)).Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = leagues
	return pagination, nil
}

func (r *LeagueRepository) Get(slug string) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) GetByPrimaryId(primaryId int) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("primary_id = ?", primaryId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) Save(c *entity.League) (*entity.League, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) Update(c *entity.League) (*entity.League, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) Delete(c *entity.League) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
