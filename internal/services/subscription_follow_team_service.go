package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionFollowTeamService struct {
	subFollowTeamRepo repository.ISubscriptionFollowTeamRepository
}

func NewSubscriptionFollowTeamService(
	subFollowTeamRepo repository.ISubscriptionFollowTeamRepository,
) *SubscriptionFollowTeamService {
	return &SubscriptionFollowTeamService{
		subFollowTeamRepo: subFollowTeamRepo,
	}
}

type ISubscriptionFollowTeamService interface {
	IsSubFollowTeam(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionFollowTeam, error)
	Save(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Update(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Delete(*entity.SubscriptionFollowTeam) error
}

func (s *SubscriptionFollowTeamService) IsSubFollowTeam(subId, teamId int) bool {
	count, _ := s.subFollowTeamRepo.Count(subId, teamId)
	return count > 0
}

func (s *SubscriptionFollowTeamService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowTeamRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowTeamService) Get(subId, teamId int) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Get(subId, teamId)
}

func (s *SubscriptionFollowTeamService) Save(a *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Save(a)
}

func (s *SubscriptionFollowTeamService) Update(a *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Update(a)
}

func (s *SubscriptionFollowTeamService) Delete(a *entity.SubscriptionFollowTeam) error {
	return s.subFollowTeamRepo.Delete(a)
}
