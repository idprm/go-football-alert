package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type MTService struct {
	mtRepo repository.IMTRepository
}

func NewMTService(mtRepo repository.IMTRepository) *MTService {
	return &MTService{
		mtRepo: mtRepo,
	}
}

type IMTService interface {
	Save(*entity.MT) (*entity.MT, error)
}

func (s *MTService) Save(a *entity.MT) (*entity.MT, error) {
	return s.mtRepo.Save(a)
}
