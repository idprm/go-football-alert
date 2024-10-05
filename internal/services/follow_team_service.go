package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type FollowTeamService struct {
	followTeamRepo repository.IFollowTeamRepository
}

func NewFollowTeamService(followTeamRepo repository.IFollowTeamRepository) *FollowTeamService {
	return &FollowTeamService{
		followTeamRepo: followTeamRepo,
	}
}

type IFollowTeamService interface {
	Save(*entity.FollowTeam) (*entity.FollowTeam, error)
	Update(*entity.FollowTeam) (*entity.FollowTeam, error)
	Delete(*entity.FollowTeam) error
}

func (s *FollowTeamService) Save(a *entity.FollowTeam) (*entity.FollowTeam, error) {
	return s.followTeamRepo.Save(a)
}

func (s *FollowTeamService) Update(a *entity.FollowTeam) (*entity.FollowTeam, error) {
	return s.followTeamRepo.Update(a)
}

func (s *FollowTeamService) Delete(a *entity.FollowTeam) error {
	return s.followTeamRepo.Delete(a)
}
