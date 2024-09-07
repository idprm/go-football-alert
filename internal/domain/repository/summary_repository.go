package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{
		db: db,
	}
}

type ISummaryRepository interface {
	Count(int, time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, time.Time) (*entity.Summary, error)
	Save(*entity.Summary) (*entity.Summary, error)
	Update(*entity.Summary) (*entity.Summary, error)
	Delete(*entity.Summary) error
}

func (r *SummaryRepository) Count(serviceId int, date time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Summary{}).Where("service_id = ?", serviceId).Where("DATE(created_at) = DATE(?)", date).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var summaries []*entity.Summary
	err := r.db.Scopes(Paginate(summaries, pagination, r.db)).Preload("Service").Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = summaries
	return pagination, nil
}

func (r *SummaryRepository) Get(serviceId int, date time.Time) (*entity.Summary, error) {
	var c entity.Summary
	err := r.db.Where("service_id = ?", serviceId).Where("DATE(created_at) = DATE(?)", date).Preload("Service").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SummaryRepository) Save(c *entity.Summary) (*entity.Summary, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SummaryRepository) Update(c *entity.Summary) (*entity.Summary, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("DATE(created_at) = DATE(?)", c.CreatedAt).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SummaryRepository) Delete(c *entity.Summary) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
