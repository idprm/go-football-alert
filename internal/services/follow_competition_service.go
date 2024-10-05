package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

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
	Save(*entity.FollowCompetition) (*entity.FollowCompetition, error)
	Update(*entity.FollowCompetition) (*entity.FollowCompetition, error)
	Delete(*entity.FollowCompetition) error
}

func (s *FollowCompetitionService) Save(a *entity.FollowCompetition) (*entity.FollowCompetition, error) {
	return s.followCompetitionRepo.Save(a)
}

func (s *FollowCompetitionService) Update(a *entity.FollowCompetition) (*entity.FollowCompetition, error) {
	return s.followCompetitionRepo.Update(a)
}

func (s *FollowCompetitionService) Delete(a *entity.FollowCompetition) error {
	return s.followCompetitionRepo.Delete(a)
}
