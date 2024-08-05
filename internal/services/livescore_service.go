package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type LiveScoreService struct {
	livescoreRepo repository.ILiveScoreRepository
}

func NewLiveScoreService(livescoreRepo repository.ILiveScoreRepository) *LiveScoreService {
	return &LiveScoreService{
		livescoreRepo: livescoreRepo,
	}
}

type ILiveScoreService interface {
	IsHome(string) bool
	GetAll() (*[]entity.Livescore, error)
	GetAllPaginate(int, int) (*[]entity.Livescore, error)
	GetById(int) (*entity.Livescore, error)
	GetBySlug(string) (*entity.Livescore, error)
	Save(*entity.Livescore) (*entity.Livescore, error)
	Update(*entity.Livescore) (*entity.Livescore, error)
	Delete(*entity.Livescore) error
}
