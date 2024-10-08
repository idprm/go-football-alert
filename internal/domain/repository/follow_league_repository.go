package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type FollowLeagueRepository struct {
	db *gorm.DB
}

func NewFollowLeagueRepository(db *gorm.DB) *FollowLeagueRepository {
	return &FollowLeagueRepository{
		db: db,
	}
}

type IFollowLeagueRepository interface {
	Save(*entity.FollowLeague) (*entity.FollowLeague, error)
	Update(*entity.FollowLeague) (*entity.FollowLeague, error)
	Delete(*entity.FollowLeague) error
}

func (r *FollowLeagueRepository) Save(c *entity.FollowLeague) (*entity.FollowLeague, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowLeagueRepository) Update(c *entity.FollowLeague) (*entity.FollowLeague, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowLeagueRepository) Delete(c *entity.FollowLeague) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
