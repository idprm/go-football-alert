package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(transactionRepo repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
	IsTransaction(int, string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string, string) (*entity.Transaction, error)
	Save(*entity.Transaction) (*entity.Transaction, error)
	Update(*entity.Transaction) (*entity.Transaction, error)
	Delete(*entity.Transaction) error
}

func (s *TransactionService) IsTransaction(serviceId int, msisdn, date string) bool {
	count, _ := s.transactionRepo.Count(serviceId, msisdn, date)
	return count > 0
}

func (s *TransactionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.transactionRepo.GetAllPaginate(pagination)
}

func (s *TransactionService) Get(serviceId int, msisdn, date string) (*entity.Transaction, error) {
	return s.transactionRepo.Get(serviceId, msisdn, date)
}

func (s *TransactionService) Save(a *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepo.Save(a)
}

func (s *TransactionService) Update(a *entity.Transaction) (*entity.Transaction, error) {
	return s.transactionRepo.Update(a)
}

func (s *TransactionService) Delete(a *entity.Transaction) error {
	return s.transactionRepo.Delete(a)
}
