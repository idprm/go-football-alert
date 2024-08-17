package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type HomeService struct {
	homeRepo repository.IHomeRepository
}

func NewHomeService(homeRepo repository.IHomeRepository) *HomeService {
	return &HomeService{
		homeRepo: homeRepo,
	}
}

type IHomeService interface {
	IsHome(int, int) bool
	IsHomeByPrimaryId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Home, error)
	GetByPrimaryId(int) (*entity.Home, error)
	Save(*entity.Home) (*entity.Home, error)
	Update(*entity.Home) (*entity.Home, error)
	Delete(*entity.Home) error
}

func (s *HomeService) IsHome(fixtureId, teamId int) bool {
	count, _ := s.homeRepo.Count(fixtureId, teamId)
	return count > 0
}

func (s *HomeService) IsHomeByPrimaryId(primaryId int) bool {
	count, _ := s.homeRepo.CountByPrimaryId(primaryId)
	return count > 0
}

func (s *HomeService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.homeRepo.GetAllPaginate(pagination)
}

func (s *HomeService) Get(fixtureId, teamId int) (*entity.Home, error) {
	return s.homeRepo.Get(fixtureId, teamId)
}

func (s *HomeService) GetByPrimaryId(primaryId int) (*entity.Home, error) {
	return s.homeRepo.GetByPrimaryId(primaryId)
}

func (s *HomeService) Save(a *entity.Home) (*entity.Home, error) {
	return s.homeRepo.Save(a)
}

func (s *HomeService) Update(a *entity.Home) (*entity.Home, error) {
	return s.homeRepo.Update(a)
}

func (s *HomeService) Delete(a *entity.Home) error {
	return s.homeRepo.Delete(a)
}
