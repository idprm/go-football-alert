package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type StandingService struct {
	standingRepo repository.IStandingRepository
}

func NewStandingService(standingRepo repository.IStandingRepository) *StandingService {
	return &StandingService{
		standingRepo: standingRepo,
	}
}

type IStandingService interface {
	IsStanding(int) bool
	IsRank(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Standing, error)
	GetByRank(int, int) (*entity.Standing, error)
	Save(*entity.Standing) (*entity.Standing, error)
	Update(*entity.Standing) (*entity.Standing, error)
	UpdateByRank(*entity.Standing) (*entity.Standing, error)
	Delete(*entity.Standing) error
}

func (s *StandingService) IsStanding(id int) bool {
	count, _ := s.standingRepo.Count(id)
	return count > 0
}

func (s *StandingService) IsRank(leagueId, rank int) bool {
	count, _ := s.standingRepo.CountByRank(leagueId, rank)
	return count > 0
}

func (s *StandingService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.standingRepo.GetAllPaginate(pagination)
}

func (s *StandingService) Get(id int) (*entity.Standing, error) {
	return s.standingRepo.Get(id)
}

func (s *StandingService) GetByRank(leagueId, rank int) (*entity.Standing, error) {
	return s.standingRepo.GetByRank(leagueId, rank)
}

func (s *StandingService) Save(a *entity.Standing) (*entity.Standing, error) {
	return s.standingRepo.Save(a)
}

func (s *StandingService) Update(a *entity.Standing) (*entity.Standing, error) {
	return s.standingRepo.Update(a)
}

func (s *StandingService) UpdateByRank(a *entity.Standing) (*entity.Standing, error) {
	return s.standingRepo.UpdateByRank(a)
}

func (s *StandingService) Delete(a *entity.Standing) error {
	return s.standingRepo.Delete(a)
}
