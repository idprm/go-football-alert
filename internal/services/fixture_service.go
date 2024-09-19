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
	IsFixture(int, int) bool
	IsFixtureByPrimaryId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Fixture, error)
	Save(*entity.Fixture) (*entity.Fixture, error)
	Update(*entity.Fixture) (*entity.Fixture, error)
	UpdateByPrimaryId(*entity.Fixture) (*entity.Fixture, error)
	Delete(*entity.Fixture) error
}

func (s *FixtureService) IsFixture(homeId, awayId int) bool {
	count, _ := s.fixtureRepo.Count(homeId, awayId)
	return count > 0
}

func (s *FixtureService) IsFixtureByPrimaryId(primaryId int) bool {
	count, _ := s.fixtureRepo.CountByPrimaryId(primaryId)
	return count > 0
}

func (s *FixtureService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.fixtureRepo.GetAllPaginate(pagination)
}

func (s *FixtureService) Get(homeId, awayId int) (*entity.Fixture, error) {
	return s.fixtureRepo.Get(homeId, awayId)
}

func (s *FixtureService) Save(a *entity.Fixture) (*entity.Fixture, error) {
	return s.fixtureRepo.Save(a)
}

func (s *FixtureService) Update(a *entity.Fixture) (*entity.Fixture, error) {
	return s.fixtureRepo.Update(a)
}

func (s *FixtureService) UpdateByPrimaryId(a *entity.Fixture) (*entity.Fixture, error) {
	return s.fixtureRepo.UpdateByPrimaryId(a)
}

func (s *FixtureService) Delete(a *entity.Fixture) error {
	return s.fixtureRepo.Delete(a)
}
