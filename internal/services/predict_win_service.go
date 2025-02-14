package services

import "github.com/idprm/go-football-alert/internal/domain/repository"

type PredictWinService struct {
	predictWinRepo repository.IPredictWinRepository
}

func NewPredictWinService(
	predictWinRepo repository.IPredictWinRepository,
) *PredictWinService {
	return &PredictWinService{
		predictWinRepo: predictWinRepo,
	}
}

type IPredictWinService interface {
}
