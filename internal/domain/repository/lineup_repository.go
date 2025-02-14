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
	CountByFixtureId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllUSSD() ([]*entity.Lineup, error)
	Get(int) (*entity.Lineup, error)
	GetByFixtureId(int) (*entity.Lineup, error)
	Save(*entity.Lineup) (*entity.Lineup, error)
	Update(*entity.Lineup) (*entity.Lineup, error)
	UpdateByFixtureId(*entity.Lineup) (*entity.Lineup, error)
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

func (r *LineupRepository) CountByFixtureId(fixtureId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Lineup{}).Where("fixture_id = ?", fixtureId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LineupRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var lineups []*entity.Lineup
	err := r.db.Where("UPPER(team_name) LIKE UPPER(?)", "%"+p.GetSearch()+"%").Scopes(Paginate(lineups, p, r.db)).Find(&lineups).Error
	if err != nil {
		return nil, err
	}
	p.Rows = lineups
	return p, nil
}

func (r *LineupRepository) GetAllUSSD() ([]*entity.Lineup, error) {
	var c []*entity.Lineup
	err := r.db.Where("DATE(fixture_date) <= DATE(NOW())").Order("DATE(fixture_date) DESC").Limit(10).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LineupRepository) Get(id int) (*entity.Lineup, error) {
	var c entity.Lineup
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LineupRepository) GetByFixtureId(fixtureId int) (*entity.Lineup, error) {
	var c entity.Lineup
	err := r.db.Where("fixture_id = ?", fixtureId).Take(&c).Error
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

func (r *LineupRepository) UpdateByFixtureId(c *entity.Lineup) (*entity.Lineup, error) {
	err := r.db.Where("fixture_id = ?", c.FixtureID).Updates(&c).Error
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
