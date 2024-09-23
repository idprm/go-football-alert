package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionCreditGoalRepository struct {
	db *gorm.DB
}

func NewSubscriptionCreditGoalRepository(db *gorm.DB) *SubscriptionCreditGoalRepository {
	return &SubscriptionCreditGoalRepository{
		db: db,
	}
}

type ISubscriptionCreditGoalRepository interface {
	Count(int, int, int) (int64, error)
	CountBySubId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionCreditGoal, error)
	Save(*entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error)
	Update(*entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error)
	Delete(*entity.SubscriptionCreditGoal) error
}

func (r *SubscriptionCreditGoalRepository) Count(subId, fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionCreditGoal{}).Where(&entity.SubscriptionCreditGoal{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionCreditGoalRepository) CountBySubId(subId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionCreditGoal{}).Where(&entity.SubscriptionCreditGoal{SubscriptionID: int64(subId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionCreditGoalRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionCreditGoal
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionCreditGoalRepository) Get(subId, fixtureId, teamId int) (*entity.SubscriptionCreditGoal, error) {
	var c entity.SubscriptionCreditGoal
	err := r.db.Where(&entity.SubscriptionCreditGoal{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionCreditGoalRepository) Save(c *entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionCreditGoalRepository) Update(c *entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionCreditGoalRepository) Delete(c *entity.SubscriptionCreditGoal) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
