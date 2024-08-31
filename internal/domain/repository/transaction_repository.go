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
	Count(int, string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string, string) (*entity.Transaction, error)
	Save(*entity.Transaction) (*entity.Transaction, error)
	Update(*entity.Transaction) (*entity.Transaction, error)
	Delete(*entity.Transaction) error
}

func (r *TransactionRepository) Count(serviceId int, msisdn, date string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Transaction{}).Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Where("DATE(created_at) = DATE(?)", date).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TransactionRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var transactions []*entity.Transaction
	err := r.db.Scopes(Paginate(transactions, pagination, r.db)).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = transactions
	return pagination, nil
}

func (r *TransactionRepository) Get(serviceId int, msisdn, date string) (*entity.Transaction, error) {
	var c entity.Transaction
	err := r.db.Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Where("DATE(created_at) = DATE(?)", date).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TransactionRepository) Save(c *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TransactionRepository) Update(c *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TransactionRepository) Delete(c *entity.Transaction) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
