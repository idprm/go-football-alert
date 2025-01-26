package services

import (
	"time"

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
	IsPronosticByStartAt(time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetById(int64) (*entity.Pronostic, error)
	Save(*entity.Pronostic) error
	Update(*entity.Pronostic) error
	Delete(*entity.Pronostic) error
}

func (s *PronosticService) IsPronosticByStartAt(startAt time.Time) bool {
	count, _ := s.pronosticRepo.CountByStartAt(startAt)
	return count > 0
}

func (s *PronosticService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.pronosticRepo.GetAllPaginate(p)
}

func (s *PronosticService) GetById(id int64) (*entity.Pronostic, error) {
	return s.pronosticRepo.GetById(id)
}

func (s *PronosticService) Save(a *entity.Pronostic) error {
	return s.pronosticRepo.Save(a)
}

func (s *PronosticService) Update(a *entity.Pronostic) error {
	return s.pronosticRepo.Update(a)
}

func (s *PronosticService) Delete(a *entity.Pronostic) error {
	return s.pronosticRepo.Delete(a)
}
