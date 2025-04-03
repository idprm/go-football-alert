package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type MOService struct {
	moRepo repository.IMORepository
}

func NewMOService(moRepo repository.IMORepository) *MOService {
	return &MOService{
		moRepo: moRepo,
	}
}

type IMOService interface {
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.MO) (*entity.MO, error)
}

func (s *MOService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.moRepo.GetAllPaginate(pagination)
}

func (s *MOService) Save(a *entity.MO) (*entity.MO, error) {
	return s.moRepo.Save(a)
}
