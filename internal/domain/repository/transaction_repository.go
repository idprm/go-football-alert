package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

type ITransactionRepository interface {
	Count(int, string, string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllRevenueDailyPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string, string, string) (*entity.Transaction, error)
	Save(*entity.Transaction) error
	Update(*entity.Transaction) error
	Delete(*entity.Transaction) error
}

func (r *TransactionRepository) Count(serviceId int, msisdn, code, date string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("DATE(created_at) = DATE(?)", date).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var transactions []*entity.Transaction
	err := r.db.Where("UPPER(msisdn) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateTransactions(transactions, p, r.db)).Preload("Service").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	p.Rows = transactions
	return p, nil
}

func (r *TransactionRepository) GetAllRevenueDailyPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var transactions []*entity.Transaction
	err := r.db.Select("DATE(created_at), subject, status, COUNT(1), SUM(amount) as revenue").Where("UPPER(msisdn) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateTransactions(transactions, p, r.db)).Group("DATE(created_at), subject, status").Preload("Service").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	p.Rows = transactions
	return p, nil
}

func (r *TransactionRepository) Get(serviceId int, msisdn, code, date string) (*entity.Transaction, error) {
	var c entity.Transaction
	err := r.db.Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("DATE(created_at) = DATE(?)", date).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TransactionRepository) Save(c *entity.Transaction) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) Update(c *entity.Transaction) error {
	err := r.db.Where("service_id = ? AND msisdn = ? AND code = ? AND DATE(created_at) = DATE(NOW())", c.ServiceID, c.Msisdn, c.Code).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) Delete(c *entity.Transaction) error {
	err := r.db.Where("service_id = ? AND msisdn = ? AND code = ? AND subject = ? AND status = ? AND DATE(created_at) = DATE(NOW())", c.ServiceID, c.Msisdn, c.Code, c.Subject, c.Status).Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) SelectRevenue() (time.Time, string, string, int, float64, error) {
	type NResult struct {
		CreatedAt time.Time
		Subject   string
		Status    string
		Total     int
		Revenue   float64
	}
	var n NResult
	err := r.db.Table("transactions").Select("DATE(created_at) as created_at, subject, status, total, SUM(amount) as revenue").Where("DATE(created_at) = DATE(NOW())").Group("DATE(created_at), subject, status").Order("DATE(created_at) ASC").Scan(&n).Error
	if err != nil {
		return time.Time{}, "", "", 0, 0, err
	}
	return n.CreatedAt, n.Subject, n.Status, n.Total, n.Revenue, nil
}
