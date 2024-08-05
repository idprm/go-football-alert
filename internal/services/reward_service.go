package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type RewardService struct {
	rewardRepo repository.IRewardRepository
}

func NewRewardService(rewardRepo repository.IRewardRepository) *RewardService {
	return &RewardService{
		rewardRepo: rewardRepo,
	}
}

type IRewardService interface {
	IsHome(string) bool
	GetAll() (*[]entity.Reward, error)
	GetAllPaginate(int, int) (*[]entity.Reward, error)
	GetById(int) (*entity.Reward, error)
	GetBySlug(string) (*entity.Reward, error)
	Save(*entity.Reward) (*entity.Reward, error)
	Update(*entity.Reward) (*entity.Reward, error)
	Delete(*entity.Reward) error
}
