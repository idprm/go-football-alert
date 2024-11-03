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
	IsTeamByCode(string) bool
	IsTeamByPrimaryId(int) bool
	IsTeamByName(string) bool
	IsLeagueTeam(*entity.LeagueTeam) bool
	IsLeagueByTeam(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllTeamUSSD(int, int) ([]*entity.LeagueTeam, error)
	Get(string) (*entity.Team, error)
	GetByCode(string) (*entity.Team, error)
	GetByPrimaryId(int) (*entity.Team, error)
	GetByName(string) (*entity.Team, error)
	Save(*entity.Team) (*entity.Team, error)
	Update(*entity.Team) (*entity.Team, error)
	UpdateByPrimaryId(*entity.Team) (*entity.Team, error)
	Delete(*entity.Team) error
	GetLeagueByTeam(int) (*entity.LeagueTeam, error)
	SaveLeagueTeam(*entity.LeagueTeam) (*entity.LeagueTeam, error)
}

func (s *TeamService) IsTeam(slug string) bool {
	count, _ := s.teamRepo.Count(slug)
	return count > 0
}

func (s *TeamService) IsTeamByCode(code string) bool {
	count, _ := s.teamRepo.CountByCode(code)
	return count > 0
}

func (s *TeamService) IsTeamByPrimaryId(primaryId int) bool {
	count, _ := s.teamRepo.CountByPrimaryId(primaryId)
	return count > 0
}

func (s *TeamService) IsTeamByName(name string) bool {
	count, _ := s.teamRepo.CountByName(name)
	return count > 0
}

func (s *TeamService) IsLeagueTeam(v *entity.LeagueTeam) bool {
	count, _ := s.teamRepo.CountByLeagueTeam(v)
	return count > 0
}

func (s *TeamService) IsLeagueByTeam(teamId int) bool {
	count, _ := s.teamRepo.CountLeagueByTeam(teamId)
	return count > 0
}

func (s *TeamService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.teamRepo.GetAllPaginate(pagination)
}

func (s *TeamService) GetAllTeamUSSD(leagueId, page int) ([]*entity.LeagueTeam, error) {
	return s.teamRepo.GetAllTeamUSSD(leagueId, page)
}

func (s *TeamService) Get(slug string) (*entity.Team, error) {
	return s.teamRepo.Get(slug)
}

func (s *TeamService) GetByCode(code string) (*entity.Team, error) {
	return s.teamRepo.GetByCode(code)
}

func (s *TeamService) GetByPrimaryId(primaryId int) (*entity.Team, error) {
	return s.teamRepo.GetByPrimaryId(primaryId)
}

func (s *TeamService) GetByName(name string) (*entity.Team, error) {
	return s.teamRepo.GetByName(name)
}

func (s *TeamService) Save(a *entity.Team) (*entity.Team, error) {
	return s.teamRepo.Save(a)
}

func (s *TeamService) Update(a *entity.Team) (*entity.Team, error) {
	return s.teamRepo.Update(a)
}

func (s *TeamService) UpdateByPrimaryId(a *entity.Team) (*entity.Team, error) {
	return s.teamRepo.UpdateByPrimaryId(a)
}

func (s *TeamService) Delete(a *entity.Team) error {
	return s.teamRepo.Delete(a)
}

func (s *TeamService) GetLeagueByTeam(teamId int) (*entity.LeagueTeam, error) {
	return s.teamRepo.GetLeagueByTeam(teamId)
}

func (s *TeamService) SaveLeagueTeam(a *entity.LeagueTeam) (*entity.LeagueTeam, error) {
	return s.teamRepo.SaveLeagueTeam(a)
}
