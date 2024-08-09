package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{
		db: db,
	}
}

type IRewardRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Reward, error)
	Save(*entity.Reward) (*entity.Reward, error)
	Update(*entity.Reward) (*entity.Reward, error)
	Delete(*entity.Reward) error
}

func (r *RewardRepository) Count(fixtureId, subscriptionId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Reward{}).Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subscriptionId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *RewardRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var rewards []*entity.Reward
	err := r.db.Scopes(Paginate(rewards, pagination, r.db)).Find(&rewards).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = rewards
	return pagination, nil
}

func (r *RewardRepository) Get(fixtureId, subscriptionId int) (*entity.Reward, error) {
	var c entity.Reward
	err := r.db.Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subscriptionId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *RewardRepository) Save(c *entity.Reward) (*entity.Reward, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *RewardRepository) Update(c *entity.Reward) (*entity.Reward, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *RewardRepository) Delete(c *entity.Reward) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
