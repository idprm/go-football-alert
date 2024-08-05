package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SeasonService struct {
	seasonRepo repository.ISeasonRepository
}

func NewSeasonService(seasonRepo repository.ISeasonRepository) *SeasonService {
	return &SeasonService{
		seasonRepo: seasonRepo,
	}
}

type ISeasonService interface {
	IsHome(string) bool
	GetAll() (*[]entity.Season, error)
	GetAllPaginate(int, int) (*[]entity.Season, error)
	GetById(int) (*entity.Season, error)
	GetBySlug(string) (*entity.Season, error)
	Save(*entity.Season) (*entity.Season, error)
	Update(*entity.Season) (*entity.Season, error)
	Delete(*entity.Season) error
}
