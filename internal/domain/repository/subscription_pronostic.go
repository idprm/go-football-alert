package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionPronosticRepository struct {
	db *gorm.DB
}

func NewSubscriptionPronosticRepository(db *gorm.DB) *SubscriptionPronosticRepository {
	return &SubscriptionPronosticRepository{
		db: db,
	}
}

type ISubscriptionPronosticRepository interface {
	Count(int, int) (int64, error)
	CountBySubId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionPronostic, error)
	Save(*entity.SubscriptionPronostic) error
	Update(*entity.SubscriptionPronostic) error
	Delete(*entity.SubscriptionPronostic) error
}

func (r *SubscriptionPronosticRepository) Count(subId, pronosticId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionPronostic{}).Where(&entity.SubscriptionPronostic{SubscriptionID: int64(subId), PronosticID: int64(pronosticId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionPronosticRepository) CountBySubId(subId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionPronostic{}).Where(&entity.SubscriptionPronostic{SubscriptionID: int64(subId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionPronosticRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var subs []*entity.SubscriptionPronostic
	err := r.db.Scopes(Paginate(subs, p, r.db)).Find(&subs).Error
	if err != nil {
		return nil, err
	}
	p.Rows = subs
	return p, nil
}

func (r *SubscriptionPronosticRepository) Get(subId, pronosticId int) (*entity.SubscriptionPronostic, error) {
	var c entity.SubscriptionPronostic
	err := r.db.Where(&entity.SubscriptionPronostic{SubscriptionID: int64(subId), PronosticID: int64(pronosticId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionPronosticRepository) Save(c *entity.SubscriptionPronostic) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionPronosticRepository) Update(c *entity.SubscriptionPronostic) error {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionPronosticRepository) Delete(c *entity.SubscriptionPronostic) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
