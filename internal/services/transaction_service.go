package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(
	transactionRepo repository.ITransactionRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

type ITransactionService interface {
	IsTransaction(int, string, string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string, string, string) (*entity.Transaction, error)
	Save(*entity.Transaction) error
	Update(*entity.Transaction) error
	Delete(*entity.Transaction) error
}

func (s *TransactionService) IsTransaction(serviceId int, msisdn, code, date string) bool {
	count, _ := s.transactionRepo.Count(serviceId, msisdn, code, date)
	return count > 0
}

func (s *TransactionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.transactionRepo.GetAllPaginate(pagination)
}

func (s *TransactionService) Get(serviceId int, msisdn, code, date string) (*entity.Transaction, error) {
	return s.transactionRepo.Get(serviceId, msisdn, code, date)
}

func (s *TransactionService) Save(a *entity.Transaction) error {
	return s.transactionRepo.Save(a)
}

func (s *TransactionService) Update(a *entity.Transaction) error {
	d := &entity.Transaction{
		ServiceID: a.ServiceID,
		Msisdn:    a.Msisdn,
		Code:      a.Code,
		Subject:   a.Subject,
		Status:    "FAILED",
	}
	errD := s.transactionRepo.Delete(d)
	if errD != nil {
		return errD
	}

	errS := s.transactionRepo.Save(a)
	if errS != nil {
		return errS
	}

	return nil

}

func (s *TransactionService) Delete(a *entity.Transaction) error {
	return s.transactionRepo.Delete(a)
}
