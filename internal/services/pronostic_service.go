package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type PronosticService struct {
	pronosticRepo repository.IPronosticRepository
}

func NewPronosticService(pronosticRepo repository.IPronosticRepository) *PronosticService {
	return &PronosticService{
		pronosticRepo: pronosticRepo,
	}
}

type IPronosticService interface {
	IsPronosticByFixtureId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Pronostic, error)
	Save(*entity.Pronostic) (*entity.Pronostic, error)
	Update(*entity.Pronostic) (*entity.Pronostic, error)
	Delete(*entity.Pronostic) error
}

func (s *PronosticService) IsPronosticByFixtureId(fixtureId int) bool {
	count, _ := s.pronosticRepo.Count(fixtureId)
	return count > 0
}

func (s *PronosticService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.pronosticRepo.GetAllPaginate(pagination)
}

func (s *PronosticService) Get(fixtureId int) (*entity.Pronostic, error) {
	return s.pronosticRepo.Get(fixtureId)
}

func (s *PronosticService) Save(a *entity.Pronostic) (*entity.Pronostic, error) {
	return s.pronosticRepo.Save(a)
}

func (s *PronosticService) Update(a *entity.Pronostic) (*entity.Pronostic, error) {
	return s.pronosticRepo.Update(a)
}

func (s *PronosticService) Delete(a *entity.Pronostic) error {
	return s.pronosticRepo.Delete(a)
}
