package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionCreditScoreService struct {
	subCreditScoreRepo repository.ISubscriptionCreditScoreRepository
}

func NewSubscriptionCreditScoreService(subCreditScoreRepo repository.ISubscriptionCreditScoreRepository) *SubscriptionCreditScoreService {
	return &SubscriptionCreditScoreService{
		subCreditScoreRepo: subCreditScoreRepo,
	}
}

type ISubscriptionCreditScoreService interface {
	IsSubCreditScore(int, int, int) bool
	IsSubCreditScoreBySubId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionCreditScore, error)
	Save(*entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error)
	Update(*entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error)
	Delete(*entity.SubscriptionCreditScore) error
}

func (s *SubscriptionCreditScoreService) IsSubCreditScore(subId, fixtureId, teamId int) bool {
	count, _ := s.subCreditScoreRepo.Count(subId, fixtureId, teamId)
	return count > 0
}

func (s *SubscriptionCreditScoreService) IsSubCreditScoreBySubId(subId int) bool {
	count, _ := s.subCreditScoreRepo.CountBySubId(subId)
	return count > 0
}

func (s *SubscriptionCreditScoreService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subCreditScoreRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionCreditScoreService) Get(subId, fixtureId, teamId int) (*entity.SubscriptionCreditScore, error) {
	return s.subCreditScoreRepo.Get(subId, fixtureId, teamId)
}

func (s *SubscriptionCreditScoreService) Save(a *entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error) {
	return s.subCreditScoreRepo.Save(a)
}

func (s *SubscriptionCreditScoreService) Update(a *entity.SubscriptionCreditScore) (*entity.SubscriptionCreditScore, error) {
	return s.subCreditScoreRepo.Update(a)
}

func (s *SubscriptionCreditScoreService) Delete(a *entity.SubscriptionCreditScore) error {
	return s.subCreditScoreRepo.Delete(a)
}
