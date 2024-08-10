package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type LeagueService struct {
	leagueRepo repository.ILeagueRepository
}

func NewLeagueService(leagueRepo repository.ILeagueRepository) *LeagueService {
	return &LeagueService{
		leagueRepo: leagueRepo,
	}
}

type ILeagueService interface {
	IsLeague(string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.League, error)
	Save(*entity.League) (*entity.League, error)
	Update(*entity.League) (*entity.League, error)
	Delete(*entity.League) error
}

func (s *LeagueService) IsLeague(key string) bool {
	count, _ := s.leagueRepo.Count(key)
	return count > 0
}

func (s *LeagueService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.leagueRepo.GetAllPaginate(pagination)
}

func (s *LeagueService) Get(slug string) (*entity.League, error) {
	return s.leagueRepo.Get(slug)
}

func (s *LeagueService) Save(a *entity.League) (*entity.League, error) {
	return s.leagueRepo.Save(a)
}

func (s *LeagueService) Update(a *entity.League) (*entity.League, error) {
	return s.leagueRepo.Update(a)
}

func (s *LeagueService) Delete(a *entity.League) error {
	return s.leagueRepo.Delete(a)
}
