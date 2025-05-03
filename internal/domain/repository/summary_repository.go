package repository

import (
	"database/sql"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"gorm.io/gorm"
)

var (
	querySelectTotalActiveSub     = "SELECT COUNT(1) as total_sub FROM subscriptions WHERE is_active = true"
	querySelectTotalRevenue       = "SELECT SUM(total_amount) as total_revenue FROM subscriptions"
	querySelectPopulateRevenue    = "SELECT DATE(created_at) as created_at, subject, status, COUNT(1) as total, SUM(amount) as revenue FROM transactions WHERE DATE(created_at) BETWEEN DATE_ADD(NOW(), INTERVAL -1 DAY) AND DATE(NOW()) GROUP BY DATE(created_at), subject, status ORDER BY DATE(created_at) ASC"
	querySelectPopulateTotalDaily = "SELECT DATE(a.created_at) as created_at, IFNULL(b.total, 0) + IFNUL(i.total, 0) + IFNULL(c.total, 0) as total_sub, IFNULL(d.total, 0) as total_unsub, IFNULL(e.total, 0) + IFNULL(f.total, 0) as total_renewal, IFNULL(h.revenue, 0) + IFNULL(g.revenue, 0) as total_revenue, IFNULL(j.total_active_sub, 0) as total_active_sub FROM summary_revenues a LEFT JOIN summary_revenues b ON b.created_at = a.created_at AND b.subject = 'FIRSTPUSH' AND b.status = 'SUCCESS' LEFT JOIN summary_revenues i ON i.created_at = a.created_at AND i.subject = 'FIRSTPUSH' AND i.status = 'FAILED' LEFT JOIN summary_revenues c ON c.created_at = a.created_at AND c.subject = 'FREE' LEFT JOIN summary_revenues d ON d.created_at = a.created_at AND d.subject = 'UNSUB' LEFT JOIN summary_revenues e ON e.created_at = a.created_at AND e.subject = 'RENEWAL' AND e.status = 'SUCCESS' LEFT JOIN summary_revenues f ON f.created_at = a.created_at AND f.subject = 'RENEWAL' AND f.status = 'FAILED' LEFT JOIN summary_revenues g ON g.created_at = a.created_at AND g.subject = 'FIRSTPUSH' AND g.status = 'SUCCESS' LEFT JOIN summary_revenues h ON h.created_at = a.created_at AND h.subject = 'RENEWAL' AND h.status = 'SUCCESS' LEFT JOIN summary_dashboards j ON DATE(j.created_at) = DATE(a.created_at) WHERE DATE(a.created_at) BETWEEN DATE_ADD(NOW(), INTERVAL -1 DAY) AND DATE(NOW()) GROUP BY DATE(a.created_at) ORDER BY DATE(a.created_at) ASC"
)

type SummaryDashboardRepository struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

type SummaryRevenueRepository struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

type SummaryTotalDailyRepository struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

func NewSummaryDashboardRepository(db *gorm.DB, sqlDB *sql.DB) *SummaryDashboardRepository {
	return &SummaryDashboardRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

func NewSummaryRevenueRepository(db *gorm.DB, sqlDB *sql.DB) *SummaryRevenueRepository {
	return &SummaryRevenueRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

func NewSummaryTotalDailyRepository(db *gorm.DB, sqlDB *sql.DB) *SummaryTotalDailyRepository {
	return &SummaryTotalDailyRepository{
		db:    db,
		sqlDB: sqlDB,
	}
}

type ISummaryDashboardRepository interface {
	Count(time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(time.Time) (*entity.SummaryDashboard, error)
	Save(*entity.SummaryDashboard) error
	Update(*entity.SummaryDashboard) error
	Delete(*entity.SummaryDashboard) error
	GetTotalActiveSub() (int, error)
	GetTotalRevenue() (float64, error)
}

type ISummaryRevenueRepository interface {
	Count(time.Time, string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByRange(*model.RangeDateRequest) ([]*entity.SummaryRevenue, error)
	Get(time.Time, string, string) (*entity.SummaryRevenue, error)
	Save(*entity.SummaryRevenue) error
	Update(*entity.SummaryRevenue) error
	Delete(*entity.SummaryRevenue) error
	SelectRevenue() (*[]entity.SummaryRevenue, error)
}

type ISummaryTotalDailyRepository interface {
	Count(time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByRange(*model.RangeDateRequest) ([]*entity.SummaryTotalDaily, error)
	Get(time.Time) (*entity.SummaryTotalDaily, error)
	Save(*entity.SummaryTotalDaily) error
	Update(*entity.SummaryTotalDaily) error
	Delete(*entity.SummaryTotalDaily) error
	SelectTotalDaily() (*[]entity.SummaryTotalDaily, error)
}

func (r *SummaryDashboardRepository) Count(date time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SummaryDashboard{}).Where("DATE(created_at) = DATE(?)", date).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryDashboardRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var summaries []*entity.SummaryDashboard
	err := r.db.Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Scopes(PaginateSummary(summaries, p, r.db)).Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	p.Rows = summaries
	return p, nil
}

func (r *SummaryDashboardRepository) Get(date time.Time) (*entity.SummaryDashboard, error) {
	var c entity.SummaryDashboard
	err := r.db.Where("DATE(created_at) = DATE(?)", date).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SummaryDashboardRepository) Save(c *entity.SummaryDashboard) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryDashboardRepository) Update(c *entity.SummaryDashboard) error {
	err := r.db.Where("DATE(created_at) = DATE(?)", c.CreatedAt).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryDashboardRepository) Delete(c *entity.SummaryDashboard) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryRevenueRepository) Count(date time.Time, subject, status string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SummaryRevenue{}).Where("DATE(created_at) = DATE(?) AND subject = ? AND status = ?", date, subject, status).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryRevenueRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var summaries []*entity.SummaryRevenue
	err := r.db.Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Scopes(PaginateSummary(summaries, p, r.db)).Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	p.Rows = summaries
	return p, nil
}

func (r *SummaryRevenueRepository) GetAllByRange(p *model.RangeDateRequest) ([]*entity.SummaryRevenue, error) {
	var summaries []*entity.SummaryRevenue
	err := r.db.Where("subject IN('FIRSTPUSH', 'RENEWAL') AND status = 'SUCCESS' AND DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Order("DATE(created_at) DESC").Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	return summaries, nil
}

func (r *SummaryRevenueRepository) Get(date time.Time, subject, status string) (*entity.SummaryRevenue, error) {
	var c entity.SummaryRevenue
	err := r.db.Where("DATE(created_at) = DATE(?) AND subject = ? AND status = ?", date, subject, status).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SummaryRevenueRepository) Save(c *entity.SummaryRevenue) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryRevenueRepository) Update(c *entity.SummaryRevenue) error {
	err := r.db.Where("DATE(created_at) = DATE(?) AND subject = ? AND status = ?", c.CreatedAt, c.Subject, c.Status).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryRevenueRepository) Delete(c *entity.SummaryRevenue) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil

}

func (r *SummaryDashboardRepository) GetTotalActiveSub() (int, error) {
	var count int
	err := r.sqlDB.QueryRow(querySelectTotalActiveSub).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryDashboardRepository) GetTotalRevenue() (float64, error) {
	var total float64
	err := r.sqlDB.QueryRow(querySelectTotalRevenue).Scan(&total)
	if err != nil {
		return total, err
	}
	return total, nil
}

func (r *SummaryRevenueRepository) SelectRevenue() (*[]entity.SummaryRevenue, error) {
	rows, err := r.sqlDB.Query(querySelectPopulateRevenue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summs []entity.SummaryRevenue

	for rows.Next() {
		var s entity.SummaryRevenue

		if err := rows.Scan(&s.CreatedAt, &s.Subject, &s.Status, &s.Total, &s.Revenue); err != nil {
			return nil, err
		}
		summs = append(summs, s)
	}

	if err = rows.Err(); err != nil {
		return &summs, err
	}

	return &summs, nil
}

func (r *SummaryTotalDailyRepository) Count(date time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SummaryTotalDaily{}).Where("DATE(created_at) = DATE(?)", date).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SummaryTotalDailyRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var summaries []*entity.SummaryTotalDaily
	err := r.db.Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Scopes(PaginateSummary(summaries, p, r.db)).Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	p.Rows = summaries
	return p, nil
}

func (r *SummaryTotalDailyRepository) GetAllByRange(p *model.RangeDateRequest) ([]*entity.SummaryTotalDaily, error) {
	var summaries []*entity.SummaryTotalDaily
	err := r.db.Where("subject IN('FIRSTPUSH', 'RENEWAL') AND status = 'SUCCESS' AND DATE(created_at) BETWEEN DATE(?) AND DATE(?)", p.GetStartDate(), p.GetEndDate()).Order("DATE(created_at) DESC").Find(&summaries).Error
	if err != nil {
		return nil, err
	}
	return summaries, nil
}

func (r *SummaryTotalDailyRepository) Get(date time.Time) (*entity.SummaryTotalDaily, error) {
	var c entity.SummaryTotalDaily
	err := r.db.Where("DATE(created_at) = DATE(?)", date).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SummaryTotalDailyRepository) Save(c *entity.SummaryTotalDaily) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryTotalDailyRepository) Update(c *entity.SummaryTotalDaily) error {
	err := r.db.Where("DATE(created_at) = DATE(?)", c.CreatedAt).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SummaryTotalDailyRepository) Delete(c *entity.SummaryTotalDaily) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil

}

func (r *SummaryTotalDailyRepository) SelectTotalDaily() (*[]entity.SummaryTotalDaily, error) {
	rows, err := r.sqlDB.Query(querySelectPopulateTotalDaily)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summs []entity.SummaryTotalDaily

	for rows.Next() {
		var s entity.SummaryTotalDaily
		if err := rows.Scan(&s.CreatedAt, &s.TotalSub, &s.TotalUnsub, &s.TotalRenewal, &s.TotalRevenue, &s.TotalActiveSub); err != nil {
			return nil, err
		}
		summs = append(summs, s)
	}

	if err = rows.Err(); err != nil {
		return &summs, err
	}

	return &summs, nil
}
