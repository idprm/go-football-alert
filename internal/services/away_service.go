package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type AwayService struct {
	awayRepo repository.IAwayRepository
}

func NewAwayService(awayRepo repository.IAwayRepository) *AwayService {
	return &AwayService{
		awayRepo: awayRepo,
	}
}

type IAwayService interface {
	IsAway(string) bool
	GetAll() (*[]entity.Away, error)
	GetAllPaginate(int, int) (*[]entity.Away, error)
	GetById(int) (*entity.Away, error)
	GetBySlug(string) (*entity.Away, error)
	Save(*entity.Away) (*entity.Away, error)
	Update(*entity.Away) (*entity.Away, error)
	Delete(*entity.Away) error
}
