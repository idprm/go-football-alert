package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionFollowCompetitionService struct {
	subFollowCompetitionRepo repository.ISubscriptionFollowCompetitionRepository
}

func NewSubscriptionFollowCompetitionService(
	subFollowCompetitionRepo repository.ISubscriptionFollowCompetitionRepository,
) *SubscriptionFollowCompetitionService {
	return &SubscriptionFollowCompetitionService{
		subFollowCompetitionRepo: subFollowCompetitionRepo,
	}
}

type ISubscriptionFollowCompetitionService interface {
	IsSubFollowCompetition(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionFollowCompetition, error)
	Save(*entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error)
	Update(*entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error)
	Delete(*entity.SubscriptionFollowCompetition) error
}

func (s *SubscriptionFollowCompetitionService) IsSubFollowCompetition(subId, leagueId int) bool {
	count, _ := s.subFollowCompetitionRepo.Count(subId, leagueId)
	return count > 0
}

func (s *SubscriptionFollowCompetitionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowCompetitionRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowCompetitionService) Get(subId, leagueId int) (*entity.SubscriptionFollowCompetition, error) {
	return s.subFollowCompetitionRepo.Get(subId, leagueId)
}

func (s *SubscriptionFollowCompetitionService) Save(a *entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error) {
	return s.subFollowCompetitionRepo.Save(a)
}

func (s *SubscriptionFollowCompetitionService) Update(a *entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error) {
	return s.subFollowCompetitionRepo.Update(a)
}

func (s *SubscriptionFollowCompetitionService) Delete(a *entity.SubscriptionFollowCompetition) error {
	return s.subFollowCompetitionRepo.Delete(a)
}
