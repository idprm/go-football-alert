package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type LiveScoreService struct {
	livescoreRepo repository.ILiveScoreRepository
}

func NewLiveScoreService(
	livescoreRepo repository.ILiveScoreRepository,
) *LiveScoreService {
	return &LiveScoreService{
		livescoreRepo: livescoreRepo,
	}
}

type ILiveScoreService interface {
	IsLiveScore(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Livescore, error)
	Save(*entity.Livescore) (*entity.Livescore, error)
	Update(*entity.Livescore) (*entity.Livescore, error)
	Delete(*entity.Livescore) error
}

func (s *LiveScoreService) IsLiveScore(fixtureId int) bool {
	count, _ := s.livescoreRepo.Count(fixtureId)
	return count > 0
}

func (s *LiveScoreService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.livescoreRepo.GetAllPaginate(pagination)
}

func (s *LiveScoreService) Get(fixtureId int) (*entity.Livescore, error) {
	return s.livescoreRepo.Get(fixtureId)
}

func (s *LiveScoreService) Save(a *entity.Livescore) (*entity.Livescore, error) {
	return s.livescoreRepo.Save(a)
}

func (s *LiveScoreService) Update(a *entity.Livescore) (*entity.Livescore, error) {
	return s.livescoreRepo.Update(a)
}

func (s *LiveScoreService) Delete(a *entity.Livescore) error {
	return s.livescoreRepo.Delete(a)
}
