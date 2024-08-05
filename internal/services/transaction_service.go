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
	IsHome(string) bool
	GetAll() (*[]entity.Transaction, error)
	GetAllPaginate(int, int) (*[]entity.Transaction, error)
	GetById(int) (*entity.Transaction, error)
	GetBySlug(string) (*entity.Transaction, error)
	Save(*entity.Transaction) (*entity.Transaction, error)
	Update(*entity.Transaction) (*entity.Transaction, error)
	Delete(*entity.Transaction) error
}
