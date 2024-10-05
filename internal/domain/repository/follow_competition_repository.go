package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type FollowCompetitionRepository struct {
	db *gorm.DB
}

func NewFollowCompetitionRepository(db *gorm.DB) *FollowCompetitionRepository {
	return &FollowCompetitionRepository{
		db: db,
	}
}

type IFollowCompetitionRepository interface {
	Save(*entity.FollowCompetition) (*entity.FollowCompetition, error)
	Update(*entity.FollowCompetition) (*entity.FollowCompetition, error)
	Delete(*entity.FollowCompetition) error
}

func (r *FollowCompetitionRepository) Save(c *entity.FollowCompetition) (*entity.FollowCompetition, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowCompetitionRepository) Update(c *entity.FollowCompetition) (*entity.FollowCompetition, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FollowCompetitionRepository) Delete(c *entity.FollowCompetition) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
