package services

import "github.com/idprm/go-football-alert/internal/domain/repository"

type FollowTeamService struct {
	followTeamRepo repository.IFollowTeamRepository
}

func NewFollowTeamService(followTeamRepo repository.IFollowTeamRepository) *FollowTeamService {
	return &FollowTeamService{
		followTeamRepo: followTeamRepo,
	}
}

type IFollowTeamService interface {
}
