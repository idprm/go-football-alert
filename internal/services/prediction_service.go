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
	IsPrediction(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Prediction, error)
	Save(*entity.Prediction) (*entity.Prediction, error)
	Update(*entity.Prediction) (*entity.Prediction, error)
	Delete(*entity.Prediction) error
}

func (s *PredictionService) IsPrediction(fixtureId int) bool {
	count, _ := s.predictionRepo.Count(fixtureId)
	return count > 0
}

func (s *PredictionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.predictionRepo.GetAllPaginate(pagination)
}

func (s *PredictionService) Get(fixtureId int) (*entity.Prediction, error) {
	return s.predictionRepo.Get(fixtureId)
}

func (s *PredictionService) Save(a *entity.Prediction) (*entity.Prediction, error) {
	return s.predictionRepo.Save(a)
}

func (s *PredictionService) Update(a *entity.Prediction) (*entity.Prediction, error) {
	return s.predictionRepo.Update(a)
}

func (s *PredictionService) Delete(a *entity.Prediction) error {
	return s.predictionRepo.Delete(a)
}
