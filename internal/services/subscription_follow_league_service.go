package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionFollowLeagueService struct {
	subFollowLeagueRepo repository.ISubscriptionFollowLeagueRepository
}

func NewSubscriptionFollowLeagueService(
	subFollowLeagueRepo repository.ISubscriptionFollowLeagueRepository,
) *SubscriptionFollowLeagueService {
	return &SubscriptionFollowLeagueService{
		subFollowLeagueRepo: subFollowLeagueRepo,
	}
}

type ISubscriptionFollowLeagueService interface {
	IsSubFollowLeague(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Delete(*entity.SubscriptionFollowLeague) error
}

func (s *SubscriptionFollowLeagueService) IsSubFollowLeague(subId, leagueId int) bool {
	count, _ := s.subFollowLeagueRepo.Count(subId, leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowLeagueRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowLeagueService) Get(subId, leagueId int) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Get(subId, leagueId)
}

func (s *SubscriptionFollowLeagueService) Save(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Save(a)
}

func (s *SubscriptionFollowLeagueService) Update(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Update(a)
}

func (s *SubscriptionFollowLeagueService) Delete(a *entity.SubscriptionFollowLeague) error {
	return s.subFollowLeagueRepo.Delete(a)
}
