package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionCreditGoalService struct {
	subCreditGoalRepo repository.ISubscriptionCreditGoalRepository
}

func NewSubscriptionCreditGoalService(subCreditGoalRepo repository.ISubscriptionCreditGoalRepository) *SubscriptionCreditGoalService {
	return &SubscriptionCreditGoalService{
		subCreditGoalRepo: subCreditGoalRepo,
	}
}

type ISubscriptionCreditGoalService interface {
	IsSubCreditGoal(int, int, int) bool
	IsSubCreditGoalBySubId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionCreditGoal, error)
	Save(*entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error)
	Update(*entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error)
	Delete(*entity.SubscriptionCreditGoal) error
}

func (s *SubscriptionCreditGoalService) IsSubCreditGoal(subId, fixtureId, teamId int) bool {
	count, _ := s.subCreditGoalRepo.Count(subId, fixtureId, teamId)
	return count > 0
}

func (s *SubscriptionCreditGoalService) IsSubCreditGoalBySubId(subId int) bool {
	count, _ := s.subCreditGoalRepo.CountBySubId(subId)
	return count > 0
}

func (s *SubscriptionCreditGoalService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subCreditGoalRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionCreditGoalService) Get(subId, fixtureId, teamId int) (*entity.SubscriptionCreditGoal, error) {
	return s.subCreditGoalRepo.Get(subId, fixtureId, teamId)
}

func (s *SubscriptionCreditGoalService) Save(a *entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error) {
	return s.subCreditGoalRepo.Save(a)
}

func (s *SubscriptionCreditGoalService) Update(a *entity.SubscriptionCreditGoal) (*entity.SubscriptionCreditGoal, error) {
	return s.subCreditGoalRepo.Update(a)
}

func (s *SubscriptionCreditGoalService) Delete(a *entity.SubscriptionCreditGoal) error {
	return s.subCreditGoalRepo.Delete(a)
}
