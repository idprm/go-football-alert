package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionPredictService struct {
	subPredictRepo repository.ISubscriptionPredictRepository
}

func NewSubscriptionPredictService(subPredictRepo repository.ISubscriptionPredictRepository) *SubscriptionPredictService {
	return &SubscriptionPredictService{
		subPredictRepo: subPredictRepo,
	}
}

type ISubscriptionPredictService interface {
	IsSubPredict(int, int, int) bool
	IsSubPredictBySubId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionPredict, error)
	Save(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Update(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Delete(*entity.SubscriptionPredict) error
}

func (s *SubscriptionPredictService) IsSubPredict(subId, fixtureId, teamId int) bool {
	count, _ := s.subPredictRepo.Count(subId, fixtureId, teamId)
	return count > 0
}

func (s *SubscriptionPredictService) IsSubPredictBySubId(subId int) bool {
	count, _ := s.subPredictRepo.CountBySubId(subId)
	return count > 0
}

func (s *SubscriptionPredictService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subPredictRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionPredictService) Get(subId, fixtureId, teamId int) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Get(subId, fixtureId, teamId)
}

func (s *SubscriptionPredictService) Save(a *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Save(a)
}

func (s *SubscriptionPredictService) Update(a *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Update(a)
}

func (s *SubscriptionPredictService) Delete(a *entity.SubscriptionPredict) error {
	return s.subPredictRepo.Delete(a)
}
