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
	IsReward(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Reward, error)
	Save(*entity.Reward) (*entity.Reward, error)
	Update(*entity.Reward) (*entity.Reward, error)
	Delete(*entity.Reward) error
}

func (s *RewardService) IsReward(fixtureId, subscriptionId int) bool {
	count, _ := s.rewardRepo.Count(fixtureId, subscriptionId)
	return count > 0
}

func (s *RewardService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.rewardRepo.GetAllPaginate(pagination)
}

func (s *RewardService) Get(fixtureId, subscriptionId int) (*entity.Reward, error) {
	return s.rewardRepo.Get(fixtureId, subscriptionId)
}

func (s *RewardService) Save(a *entity.Reward) (*entity.Reward, error) {
	return s.rewardRepo.Save(a)
}

func (s *RewardService) Update(a *entity.Reward) (*entity.Reward, error) {
	return s.rewardRepo.Update(a)
}

func (s *RewardService) Delete(a *entity.Reward) error {
	return s.rewardRepo.Delete(a)
}
