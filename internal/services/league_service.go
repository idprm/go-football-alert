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
	IsHome(string) bool
	GetAll() (*[]entity.League, error)
	GetAllPaginate(int, int) (*[]entity.League, error)
	GetById(int) (*entity.League, error)
	GetBySlug(string) (*entity.League, error)
	Save(*entity.League) (*entity.League, error)
	Update(*entity.League) (*entity.League, error)
	Delete(*entity.League) error
}
