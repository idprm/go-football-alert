package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type FollowTeamRepository struct {
	db *gorm.DB
}

func NewFollowTeamRepository(db *gorm.DB) *FollowTeamRepository {
	return &FollowTeamRepository{
		db: db,
	}
}

type IFollowTeamRepository interface {
	Save(*entity.FollowTeam) (*entity.FollowTeam, error)
	Update(*entity.FollowTeam) (*entity.FollowTeam, error)
	Delete(*entity.FollowTeam) error
}

func (r *FollowTeamRepository) Save(c *entity.FollowTeam) (*entity.FollowTeam, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowTeamRepository) Update(c *entity.FollowTeam) (*entity.FollowTeam, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowTeamRepository) Delete(c *entity.FollowTeam) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
