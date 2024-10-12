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
	IsSub(int64) bool
	IsLeague(int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetBySub(int64) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Delete(*entity.SubscriptionFollowLeague) error
}

func (s *SubscriptionFollowLeagueService) IsSub(subId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountBySub(subId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsLeague(leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByLeague(leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowLeagueRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowLeagueService) GetBySub(subId int64) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.GetBySub(subId)
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
