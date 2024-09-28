package services

import "github.com/idprm/go-football-alert/internal/domain/repository"

type FollowCompetitionService struct {
	followCompetitionRepo repository.IFollowCompetitionRepository
}

func NewFollowCompetitionService(
	followCompetitionRepo repository.IFollowCompetitionRepository,
) *FollowCompetitionService {
	return &FollowCompetitionService{
		followCompetitionRepo: followCompetitionRepo,
	}
}

type IFollowCompetitionService interface {
}
