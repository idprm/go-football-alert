package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.ISubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

type ISubscriptionService interface {
	IsHome(string) bool
	GetAll() (*[]entity.Service, error)
	GetAllPaginate(int, int) (*[]entity.Service, error)
	GetById(int) (*entity.Service, error)
	GetBySlug(string) (*entity.Service, error)
	Save(*entity.Service) (*entity.Service, error)
	Update(*entity.Service) (*entity.Service, error)
	Delete(*entity.Service) error
}
