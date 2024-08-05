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
	IsHome(string) bool
	GetAll() (*[]entity.Home, error)
	GetAllPaginate(int, int) (*[]entity.Home, error)
	GetById(int) (*entity.Home, error)
	GetBySlug(string) (*entity.Home, error)
	Save(*entity.Home) (*entity.Home, error)
	Update(*entity.Home) (*entity.Home, error)
	Delete(*entity.Home) error
}
