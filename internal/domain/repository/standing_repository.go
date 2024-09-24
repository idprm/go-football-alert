package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type StandingRepository struct {
	db *gorm.DB
}

func NewStandingRepository(db *gorm.DB) *StandingRepository {
	return &StandingRepository{
		db: db,
	}
}

type IStandingRepository interface {
	Count(int) (int64, error)
	CountByRank(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Standing, error)
	GetByRank(int, int) (*entity.Standing, error)
	Save(*entity.Standing) (*entity.Standing, error)
	Update(*entity.Standing) (*entity.Standing, error)
	UpdateByRank(*entity.Standing) (*entity.Standing, error)
	Delete(*entity.Standing) error
}

func (r *StandingRepository) Count(id int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Standing{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *StandingRepository) CountByRank(leagueId, rank int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Standing{}).Where("league_id = ?", leagueId).Where("ranking = ?", rank).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *StandingRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var standings []*entity.Standing
	err := r.db.Scopes(Paginate(standings, pagination, r.db)).Find(&standings).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = standings
	return pagination, nil
}

func (r *StandingRepository) Get(id int) (*entity.Standing, error) {
	var c entity.Standing
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *StandingRepository) GetByRank(leagueId, rank int) (*entity.Standing, error) {
	var c entity.Standing
	err := r.db.Where("league_id = ?", leagueId).Where("ranking = ?", rank).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *StandingRepository) Save(c *entity.Standing) (*entity.Standing, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *StandingRepository) Update(c *entity.Standing) (*entity.Standing, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *StandingRepository) UpdateByRank(c *entity.Standing) (*entity.Standing, error) {
	err := r.db.Where("league_id = ?", c.LeagueID).Where("ranking = ?", c.Ranking).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *StandingRepository) Delete(c *entity.Standing) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
