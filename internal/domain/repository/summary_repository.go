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
	GetSubByMonth(time.Time) (int, error)
	GetUnsubByMonth(time.Time) (int, error)
	GetRenewalByMonth(time.Time) (int, error)
	GetRevenueByMonth(time.Time) (float64, error)
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

func (r *SummaryRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var summaries []*entity.Summary
	err := r.db.Where("MONTH(created_at) = MONTH(?) AND YEAR(created_at) = YEAR(?)", p.GetDate(), p.GetDate()).Scopes(Paginate(summaries, p, r.db)).Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	p.Rows = summaries
	return p, nil
}

func (r *SummaryRepository) Get(serviceId int, date time.Time) (*entity.Summary, error) {
	var c entity.Summary
	err := r.db.Where("service_id = ?", serviceId).Where("DATE(created_at) = DATE(?)", date).Preload("Service").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SummaryRepository) GetSubByMonth(date time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_sub) as total_sub").Where("MONTH(created_at) = MONTH(?) AND YEAR(created_at) = YEAR(?)", date, date).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalSub, nil
}

func (r *SummaryRepository) GetUnsubByMonth(date time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_unsub) as total_unsub").Where("MONTH(created_at) = MONTH(?) AND YEAR(created_at) = YEAR(?)", date, date).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalUnsub, nil
}

func (r *SummaryRepository) GetRenewalByMonth(date time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_renewal) as total_renewal").Where("MONTH(created_at) = MONTH(?) AND YEAR(created_at) = YEAR(?)", date, date).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalRenewal, nil
}

func (r *SummaryRepository) GetRevenueByMonth(date time.Time) (float64, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_revenue) as total_revenue").Where("MONTH(created_at) = MONTH(?) AND YEAR(created_at) = YEAR(?)", date, date).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalRevenue, nil
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
