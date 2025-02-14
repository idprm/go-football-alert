package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type BettingService struct {
	bettingRepo repository.IBettingRepository
}

func NewBettingService(bettingRepo repository.IBettingRepository) *BettingService {
	return &BettingService{
		bettingRepo: bettingRepo,
	}
}

type IBettingService interface {
	IsBetting(int64, int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.Betting, error)
	Save(*entity.Betting) (*entity.Betting, error)
	Update(*entity.Betting) (*entity.Betting, error)
	Delete(*entity.Betting) error
}

func (s *BettingService) IsBetting(fixtureId, subId int64) bool {
	count, _ := s.bettingRepo.Count(fixtureId, subId)
	return count > 0
}

func (s *BettingService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.bettingRepo.GetAllPaginate(pagination)
}

func (s *BettingService) Get(fixtureId, subId int64) (*entity.Betting, error) {
	return s.bettingRepo.Get(fixtureId, subId)
}

func (s *BettingService) Save(a *entity.Betting) (*entity.Betting, error) {
	return s.bettingRepo.Save(a)
}

func (s *BettingService) Update(a *entity.Betting) (*entity.Betting, error) {
	return s.bettingRepo.Update(a)
}

func (s *BettingService) Delete(a *entity.Betting) error {
	return s.bettingRepo.Delete(a)
}
