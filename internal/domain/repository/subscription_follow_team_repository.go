package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionFollowTeamRepository struct {
	db *gorm.DB
}

func NewSubscriptionFollowTeamRepository(db *gorm.DB) *SubscriptionFollowTeamRepository {
	return &SubscriptionFollowTeamRepository{
		db: db,
	}
}

type ISubscriptionFollowTeamRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionFollowTeam, error)
	Save(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Update(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Delete(*entity.SubscriptionFollowTeam) error
}

func (r *SubscriptionFollowTeamRepository) Count(subId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowTeam{}).Where(&entity.SubscriptionFollowTeam{SubscriptionID: int64(subId), TeamID: int64(teamId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowTeamRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionFollowTeam
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionFollowTeamRepository) Get(subId, teamId int) (*entity.SubscriptionFollowTeam, error) {
	var c entity.SubscriptionFollowTeam
	err := r.db.Where(&entity.SubscriptionFollowTeam{SubscriptionID: int64(subId), TeamID: int64(teamId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionFollowTeamRepository) Save(c *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowTeamRepository) Update(c *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowTeamRepository) Delete(c *entity.SubscriptionFollowTeam) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
