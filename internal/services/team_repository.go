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
	IsHome(string) bool
	GetAll() (*[]entity.Team, error)
	GetAllPaginate(int, int) (*[]entity.Team, error)
	GetById(int) (*entity.Team, error)
	GetBySlug(string) (*entity.Team, error)
	Save(*entity.Team) (*entity.Team, error)
	Update(*entity.Team) (*entity.Team, error)
	Delete(*entity.Team) error
}
