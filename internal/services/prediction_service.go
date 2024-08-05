package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type PredictionService struct {
	predictionRepo repository.IPredictionRepository
}

func NewPredictionService(predictionRepo repository.IPredictionRepository) *PredictionService {
	return &PredictionService{
		predictionRepo: predictionRepo,
	}
}

type IPredictionService interface {
	IsHome(string) bool
	GetAll() (*[]entity.Prediction, error)
	GetAllPaginate(int, int) (*[]entity.Prediction, error)
	GetById(int) (*entity.Prediction, error)
	GetBySlug(string) (*entity.Prediction, error)
	Save(*entity.Prediction) (*entity.Prediction, error)
	Update(*entity.Prediction) (*entity.Prediction, error)
	Delete(*entity.Prediction) error
}
