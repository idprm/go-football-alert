package services

import "github.com/idprm/go-football-alert/internal/domain/repository"

type BettingService struct {
	bettingRepo repository.IBettingRepository
}

func NewBettingService(bettingRepo repository.IBettingRepository) *BettingService {
	return &BettingService{
		bettingRepo: bettingRepo,
	}
}

type IBettingService interface {
}
