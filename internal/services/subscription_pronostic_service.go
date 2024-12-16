package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionPronosticService struct {
	subPronosticRepo repository.ISubscriptionPronosticRepository
}

func NewSubscriptionPronosticService(
	subPronosticRepo repository.ISubscriptionPronosticRepository,
) *SubscriptionPronosticService {
	return &SubscriptionPronosticService{
		subPronosticRepo: subPronosticRepo,
	}
}

type ISubscriptionPronosticService interface {
	IsSubPronostic(int, int) bool
	IsSubPronosticBySubId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionPronostic, error)
	Save(*entity.SubscriptionPronostic) error
	Update(*entity.SubscriptionPronostic) error
	Delete(*entity.SubscriptionPronostic) error
}

func (s *SubscriptionPronosticService) IsSubPronostic(subId, pronosticId int) bool {
	count, _ := s.subPronosticRepo.Count(subId, pronosticId)
	return count > 0
}

func (s *SubscriptionPronosticService) IsSubPronosticBySubId(subId int) bool {
	count, _ := s.subPronosticRepo.CountBySubId(subId)
	return count > 0
}

func (s *SubscriptionPronosticService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.subPronosticRepo.GetAllPaginate(p)
}

func (s *SubscriptionPronosticService) Get(subId, pronosticId int) (*entity.SubscriptionPronostic, error) {
	return s.subPronosticRepo.Get(subId, pronosticId)
}

func (s *SubscriptionPronosticService) Save(a *entity.SubscriptionPronostic) error {
	return s.subPronosticRepo.Save(a)
}

func (s *SubscriptionPronosticService) Update(a *entity.SubscriptionPronostic) error {
	return s.subPronosticRepo.Update(a)
}

func (s *SubscriptionPronosticService) Delete(a *entity.SubscriptionPronostic) error {
	return s.subPronosticRepo.Delete(a)
}
