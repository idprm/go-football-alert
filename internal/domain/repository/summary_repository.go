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

var (
// queryCountSummaryPaginate       = "SELECT COUNT(*) FROM summaries"
// querySelectSummaryPaginate      = "SELECT * FROM summaries WHERE DATE(created_at) BETWEEN DATE(?) AND DATE(?) GROUP BY DATE(created_at) ORDER BY DATE(created_at) DESC"
// querySelectSummaryPaginate2     = "SELECT * FROM summaries LIMIT ? OFFSET ?"
// querySelectRevenueInTransaction = "SELECT DATE(created_at) as created_at, subject, status,  COUNT(1) as total, SUM(amount) as revenue FROM transactions WHERE DATE(created_at) BETWEEN DATE(?) AND DATE(?) GROUP BY DATE(created_at), subject, status ORDER BY DATE(created_at) DESC"
// SELECT DATE(created_at), subject, status,  COUNT(1), SUM(amount) as revenue
// FROM fb_alert_test.transactions
// WHERE DATE(created_at) BETWEEN DATE('2025-03-01') AND DATE(NOW())
// GROUP BY  DATE(created_at) , subject, status
// ORDER BY DATE(created_at) DESC;
)

type ISummaryRepository interface {
	Count(int, time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, time.Time) (*entity.Summary, error)
	GetActiveSub() (int, error)
	GetSub(time.Time, time.Time) (int, error)
	GetUnSub(time.Time, time.Time) (int, error)
	GetRenewal(time.Time, time.Time) (int, error)
	GetRevenue(time.Time, time.Time) (float64, error)
	Save(*entity.Summary) (*entity.Summary, error)
	Update(*entity.Summary) (*entity.Summary, error)
	Delete(*entity.Summary) error
	CountSummaryRevenue(*entity.SummaryRevenue) (int64, error)
	SaveSummaryRevenue(*entity.SummaryRevenue) error
	UpdateSummaryRevenue(*entity.SummaryRevenue) error
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
	err := r.db.Select("DATE(created_at) as created_at").Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Group("DATE(created_at)").Scopes(PaginateSummary(summaries, p, r.db)).Find(&summaries).Error
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

func (r *SummaryRepository) GetActiveSub() (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("total_active_sub").Where("DATE(created_at) = DATE(NOW())").Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalActiveSub, nil
}

func (r *SummaryRepository) GetSub(start, end time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_sub) as total_sub").Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", start, end).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalSub, nil
}

func (r *SummaryRepository) GetUnSub(start, end time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_unsub) as total_unsub").Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", start, end).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalUnsub, nil
}

func (r *SummaryRepository) GetRenewal(start, end time.Time) (int, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_renewal) as total_renewal").Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", start, end).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalRenewal, nil
}

func (r *SummaryRepository) GetRevenue(start, end time.Time) (float64, error) {
	var c entity.Summary
	err := r.db.Table("summaries").Select("SUM(total_revenue) as total_revenue").Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", start, end).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.TotalRevenue, nil
}

func (r *SummaryRepository) GetRevenueDaily() (*entity.Transaction, error) {
	var t entity.Transaction
	err := r.db.Table("transactions").Select("").Where("").Scan(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
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

func (r *SummaryRepository) CountSummaryRevenue(c *entity.SummaryRevenue) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SummaryRevenue{}).Where("DATE(created_at) = DATE(?) AND subject = ? AND status = ?", c.CreatedAt, c.Subject, c.Status).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryRepository) SaveSummaryRevenue(c *entity.SummaryRevenue) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryRepository) UpdateSummaryRevenue(c *entity.SummaryRevenue) error {
	err := r.db.Where("DATE(created_at) = DATE(?) AND subject = ? AND status = ?", c.CreatedAt, c.Subject, c.Status).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}
