package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionPredictRepository struct {
	db *gorm.DB
}

func NewSubscriptionPredictRepository(db *gorm.DB) *SubscriptionPredictRepository {
	return &SubscriptionPredictRepository{
		db: db,
	}
}

type ISubscriptionPredictRepository interface {
	Count(int, int, int) (int64, error)
	CountBySubId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionPredict, error)
	Save(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Update(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Delete(*entity.SubscriptionPredict) error
}

func (r *SubscriptionPredictRepository) Count(subId, fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionPredict{}).Where(&entity.SubscriptionPredict{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionPredictRepository) CountBySubId(subId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionPredict{}).Where(&entity.SubscriptionPredict{SubscriptionID: int64(subId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionPredictRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionPredict
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionPredictRepository) Get(subId, fixtureId, teamId int) (*entity.SubscriptionPredict, error) {
	var c entity.SubscriptionPredict
	err := r.db.Where(&entity.SubscriptionPredict{SubscriptionID: int64(subId), FixtureID: int64(fixtureId), TeamID: int64(teamId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionPredictRepository) Save(c *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionPredictRepository) Update(c *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionPredictRepository) Delete(c *entity.SubscriptionPredict) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}