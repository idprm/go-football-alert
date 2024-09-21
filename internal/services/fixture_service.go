package services

import (
	"time"

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
	IsFixture(int) bool
	IsFixtureByPrimaryId(int) bool
	IsFixtureByDate(time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAll() ([]*entity.Fixture, error)
	GetAllByFixtureDate(time.Time) ([]*entity.Fixture, error)
	Get(int) (*entity.Fixture, error)
	Save(*entity.Fixture) (*entity.Fixture, error)
	Update(*entity.Fixture) (*entity.Fixture, error)
	UpdateByPrimaryId(*entity.Fixture) (*entity.Fixture, error)
	Delete(*entity.Fixture) error
}

func (s *FixtureService) IsFixture(id int) bool {
	count, _ := s.fixtureRepo.Count(id)
	return count > 0
}

func (s *FixtureService) IsFixtureByPrimaryId(primaryId int) bool {
	count, _ := s.fixtureRepo.CountByPrimaryId(primaryId)
	return count > 0
}

func (s *FixtureService) IsFixtureByDate(fixDate time.Time) bool {
	count, _ := s.fixtureRepo.CountByFixtureDate(fixDate)
	return count > 0
}

func (s *FixtureService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.fixtureRepo.GetAllPaginate(pagination)
}

func (s *FixtureService) GetAll() ([]*entity.Fixture, error) {
	return s.fixtureRepo.GetAll()
}

func (s *FixtureService) GetAllByFixtureDate(fixDate time.Time) ([]*entity.Fixture, error) {
	return s.fixtureRepo.GetAllByFixtureDate(fixDate)
}

func (s *FixtureService) Get(id int) (*entity.Fixture, error) {
	return s.fixtureRepo.Get(id)
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
