package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type FixtureService struct {
	fixtureRepo repository.IFixtureRepository
}

func NewFixtureService(fixtureRepo repository.IFixtureRepository) *FixtureService {
	return &FixtureService{
		fixtureRepo: fixtureRepo,
	}
}

type IFixtureService interface {
	IsAway(string) bool
	GetAll() (*[]entity.Fixture, error)
	GetAllPaginate(int, int) (*[]entity.Fixture, error)
	GetById(int) (*entity.Fixture, error)
	GetBySlug(string) (*entity.Fixture, error)
	Save(*entity.Fixture) (*entity.Fixture, error)
	Update(*entity.Fixture) (*entity.Fixture, error)
	Delete(*entity.Fixture) error
}
