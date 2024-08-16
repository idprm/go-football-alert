package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{
		db: db,
	}
}

type IHistoryRepository interface {
	Count(int, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string) (*entity.History, error)
	Save(*entity.History) (*entity.History, error)
	Update(*entity.History) (*entity.History, error)
	Delete(*entity.History) error
}

func (r *HistoryRepository) Count(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.History{}).Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *HistoryRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var histories []*entity.History
	err := r.db.Scopes(Paginate(histories, pagination, r.db)).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = histories
	return pagination, nil
}

func (r *HistoryRepository) Get(serviceId int, msisdn string) (*entity.History, error) {
	var c entity.History
	err := r.db.Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *HistoryRepository) Save(c *entity.History) (*entity.History, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *HistoryRepository) Update(c *entity.History) (*entity.History, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *HistoryRepository) Delete(c *entity.History) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
