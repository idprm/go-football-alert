package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type FollowLeagueService struct {
	FollowLeagueRepo repository.IFollowLeagueRepository
}

func NewFollowLeagueService(
	FollowLeagueRepo repository.IFollowLeagueRepository,
) *FollowLeagueService {
	return &FollowLeagueService{
		FollowLeagueRepo: FollowLeagueRepo,
	}
}

type IFollowLeagueService interface {
	Save(*entity.FollowLeague) (*entity.FollowLeague, error)
	Update(*entity.FollowLeague) (*entity.FollowLeague, error)
	Delete(*entity.FollowLeague) error
}

func (s *FollowLeagueService) Save(a *entity.FollowLeague) (*entity.FollowLeague, error) {
	return s.FollowLeagueRepo.Save(a)
}

func (s *FollowLeagueService) Update(a *entity.FollowLeague) (*entity.FollowLeague, error) {
	return s.FollowLeagueRepo.Update(a)
}

func (s *FollowLeagueService) Delete(a *entity.FollowLeague) error {
	return s.FollowLeagueRepo.Delete(a)
}
