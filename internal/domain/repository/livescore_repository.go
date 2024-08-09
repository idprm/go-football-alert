package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type LiveScoreRepository struct {
	db *gorm.DB
}

func NewLivescoreRepository(db *gorm.DB) *LiveScoreRepository {
	return &LiveScoreRepository{
		db: db,
	}
}

type ILiveScoreRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Livescore, error)
	Save(*entity.Livescore) (*entity.Livescore, error)
	Update(*entity.Livescore) (*entity.Livescore, error)
	Delete(*entity.Livescore) error
}

func (r *LiveScoreRepository) Count(fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Livescore{}).Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LiveScoreRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var livescores []*entity.Livescore
	err := r.db.Scopes(Paginate(livescores, pagination, r.db)).Find(&livescores).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = livescores
	return pagination, nil
}

func (r *LiveScoreRepository) Get(fixtureId, teamId int) (*entity.Livescore, error) {
	var c entity.Livescore
	err := r.db.Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LiveScoreRepository) Save(c *entity.Livescore) (*entity.Livescore, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LiveScoreRepository) Update(c *entity.Livescore) (*entity.Livescore, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LiveScoreRepository) Delete(c *entity.Livescore) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
