package repository

import (
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
	Get(int, string, string, string) (*entity.Transaction, error)
	Save(*entity.Transaction) error
	Update(*entity.Transaction) error
	Delete(*entity.Transaction) error
	CountSubByDay(int) (int64, error)
	CountUnSubByDay(int) (int64, error)
	CountRenewalByDay(int) (int64, error)
	CountSuccessByDay(int) (int64, error)
	CountFailedByDay(int) (int64, error)
	TotalRevenueByDay(int) (float64, error)
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

func (r *TransactionRepository) CountSubByDay(serviceId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND DATE(created_at) = DATE(NOW()) AND subject = 'FREE'", serviceId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) CountUnSubByDay(serviceId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND DATE(created_at) = DATE(NOW()) AND subject = 'UNSUB'", serviceId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) CountRenewalByDay(serviceId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND DATE(created_at) = DATE(NOW()) AND subject = 'RENEWAL'", serviceId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) CountSuccessByDay(serviceId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND DATE(created_at) = DATE(NOW()) AND status = 'SUCCESS'", serviceId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) CountFailedByDay(serviceId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ? AND DATE(created_at) = DATE(NOW()) AND status = 'FAILED'", serviceId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) TotalRevenueByDay(serviceId int) (float64, error) {
	var c entity.Transaction
	err := r.db.Table("transactions").Select("SUM(amount) as amount").Where("service_id = ? AND DATE(created_at) = DATE(NOW())", serviceId).Scan(&c).Error
	if err != nil {
		return 0, err
	}
	return c.Amount, nil
}
