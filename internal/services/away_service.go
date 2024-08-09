package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type AwayService struct {
	awayRepo repository.IAwayRepository
}

func NewAwayService(awayRepo repository.IAwayRepository) *AwayService {
	return &AwayService{
		awayRepo: awayRepo,
	}
}

type IAwayService interface {
	IsAwayTeamId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetByTeamId(int) (*entity.Away, error)
	Save(*entity.Away) (*entity.Away, error)
	Update(*entity.Away) (*entity.Away, error)
	Delete(*entity.Away) error
}

func (s *AwayService) IsPostBySlug(teamId int) bool {
	count, _ := s.awayRepo.CountByTeamId(teamId)
	return count > 0
}

func (s *AwayService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.awayRepo.GetAllPaginate(pagination)
}

func (s *AwayService) GetByTeamId(teamId int) (*entity.Away, error) {
	return s.awayRepo.GetByTeamId(teamId)
}

func (s *AwayService) Save(a *entity.Away) (*entity.Away, error) {
	return s.awayRepo.Save(a)
}

func (s *AwayService) Update(a *entity.Away) (*entity.Away, error) {
	return s.awayRepo.Update(a)
}

func (s *AwayService) Delete(a *entity.Away) error {
	return s.awayRepo.Delete(a)
}
