package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type TeamService struct {
	teamRepo repository.ITeamRepository
}

func NewTeamService(teamRepo repository.ITeamRepository) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}

type ITeamService interface {
	IsTeam(string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Team, error)
	Save(*entity.Team) (*entity.Team, error)
	Update(*entity.Team) (*entity.Team, error)
	Delete(*entity.Team) error
}

func (s *TeamService) IsTeam(slug string) bool {
	count, _ := s.teamRepo.Count(slug)
	return count > 0
}

func (s *TeamService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.teamRepo.GetAllPaginate(pagination)
}

func (s *TeamService) Get(slug string) (*entity.Team, error) {
	return s.teamRepo.Get(slug)
}

func (s *TeamService) Save(a *entity.Team) (*entity.Team, error) {
	return s.teamRepo.Save(a)
}

func (s *TeamService) Update(a *entity.Team) (*entity.Team, error) {
	return s.teamRepo.Update(a)
}

func (s *TeamService) Delete(a *entity.Team) error {
	return s.teamRepo.Delete(a)
}
