package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionCreditScoreRepository struct {
	db *gorm.DB
}

func NewSubscriptionCreditScoreRepository(db *gorm.DB) *SubscriptionCreditScoreRepository {
	return &SubscriptionCreditScoreRepository{
		db: db,
	}
}

type ISubscriptionCreditScoreRepository interface {
	Count(int, int, int) (int64, error)
	CountBySubId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionCreditScore, error)
	Save(*entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error)
	Update(*entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error)
	Delete(*entity.SubscriptionCreditScore) error
}

func (r *SubscriptionCreditScoreRepository) Count(subId, fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionCreditScore{}).Where(&entity.SubscriptionCreditScore{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionCreditScoreRepository) CountBySubId(subId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionCreditScore{}).Where(&entity.SubscriptionCreditScore{SubscriptionID: int64(subId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionCreditScoreRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionCreditScore
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionCreditScoreRepository) Get(subId, fixtureId, teamId int) (*entity.SubscriptionCreditScore, error) {
	var c entity.SubscriptionCreditScore
	err := r.db.Where(&entity.SubscriptionCreditScore{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionCreditScoreRepository) Save(c *entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionCreditScoreRepository) Update(c *entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionCreditScoreRepository) Delete(c *entity.SubscriptionCreditScore) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
