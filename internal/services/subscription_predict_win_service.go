package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionPredictWinService struct {
	subPredictRepo repository.ISubscriptionPredictWinRepository
}

func NewSubscriptionPredictWinService(
	subPredictRepo repository.ISubscriptionPredictWinRepository,
) *SubscriptionPredictWinService {
	return &SubscriptionPredictWinService{
		subPredictRepo: subPredictRepo,
	}
}

type ISubscriptionPredictWinService interface {
	IsSubPredict(int, int, int) bool
	IsSubPredictBySubId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int, int) (*entity.SubscriptionPredict, error)
	Save(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Update(*entity.SubscriptionPredict) (*entity.SubscriptionPredict, error)
	Delete(*entity.SubscriptionPredict) error
}

func (s *SubscriptionPredictWinService) IsSubPredict(subId, fixtureId, teamId int) bool {
	count, _ := s.subPredictRepo.Count(subId, fixtureId, teamId)
	return count > 0
}

func (s *SubscriptionPredictWinService) IsSubPredictBySubId(subId int) bool {
	count, _ := s.subPredictRepo.CountBySubId(subId)
	return count > 0
}

func (s *SubscriptionPredictWinService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subPredictRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionPredictWinService) Get(subId, fixtureId, teamId int) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Get(subId, fixtureId, teamId)
}

func (s *SubscriptionPredictWinService) Save(a *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Save(a)
}

func (s *SubscriptionPredictWinService) Update(a *entity.SubscriptionPredict) (*entity.SubscriptionPredict, error) {
	return s.subPredictRepo.Update(a)
}

func (s *SubscriptionPredictWinService) Delete(a *entity.SubscriptionPredict) error {
	return s.subPredictRepo.Delete(a)
}
