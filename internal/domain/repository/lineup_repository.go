package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type LineupRepository struct {
	db *gorm.DB
}

func NewLineupRepository(db *gorm.DB) *LineupRepository {
	return &LineupRepository{
		db: db,
	}
}

type ILineupRepository interface {
	Count(int) (int64, error)
	CountByPrimaryId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Lineup, error)
	GetByPrimaryId(int) (*entity.Lineup, error)
	Save(*entity.Lineup) (*entity.Lineup, error)
	Update(*entity.Lineup) (*entity.Lineup, error)
	UpdateByPrimaryId(*entity.Lineup) (*entity.Lineup, error)
	Delete(*entity.Lineup) error
}

func (r *LineupRepository) Count(id int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Lineup{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LineupRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Lineup{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LineupRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var lineups []*entity.Lineup
	err := r.db.Scopes(Paginate(lineups, pagination, r.db)).Find(&lineups).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = lineups
	return pagination, nil
}

func (r *LineupRepository) Get(id int) (*entity.Lineup, error) {
	var c entity.Lineup
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LineupRepository) GetByPrimaryId(primaryId int) (*entity.Lineup, error) {
	var c entity.Lineup
	err := r.db.Where("primary_id = ?", primaryId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LineupRepository) Save(c *entity.Lineup) (*entity.Lineup, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LineupRepository) Update(c *entity.Lineup) (*entity.Lineup, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LineupRepository) UpdateByPrimaryId(c *entity.Lineup) (*entity.Lineup, error) {
	err := r.db.Where("primary_id = ?", c.PrimaryID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LineupRepository) Delete(c *entity.Lineup) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
