package services

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type LiveMatchService struct {
	liveMatchRepo repository.ILiveMatchRepository
}

func NewLiveMatchService(liveMatchRepo repository.ILiveMatchRepository) *LiveMatchService {
	return &LiveMatchService{
		liveMatchRepo: liveMatchRepo,
	}
}

type ILiveMatchService interface {
	IsLiveMatch(int) bool
	IsLiveMatchByDate(time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllLiveMatchUSSD(int) ([]*entity.LiveMatch, error)
	Get(int) (*entity.LiveMatch, error)
	Save(*entity.LiveMatch) (*entity.LiveMatch, error)
	Update(*entity.LiveMatch) (*entity.LiveMatch, error)
	Delete(*entity.LiveMatch) error
}

func (s *LiveMatchService) IsLiveMatch(fixtureId int) bool {
	count, _ := s.liveMatchRepo.Count(fixtureId)
	return count > 0
}

func (s *LiveMatchService) IsLiveMatchByDate(fixDate time.Time) bool {
	count, _ := s.liveMatchRepo.CountByFixtureDate(fixDate)
	return count > 0
}

func (s *LiveMatchService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.liveMatchRepo.GetAllPaginate(pagination)
}

func (s *LiveMatchService) GetAllLiveMatchUSSD(page int) ([]*entity.LiveMatch, error) {
	return s.liveMatchRepo.GetAllLiveMatchUSSD(page)
}

func (s *LiveMatchService) Get(id int) (*entity.LiveMatch, error) {
	return s.liveMatchRepo.Get(id)
}

func (s *LiveMatchService) Save(a *entity.LiveMatch) (*entity.LiveMatch, error) {
	return s.liveMatchRepo.Save(a)
}

func (s *LiveMatchService) Update(a *entity.LiveMatch) (*entity.LiveMatch, error) {
	return s.liveMatchRepo.Update(a)
}

func (s *LiveMatchService) Delete(a *entity.LiveMatch) error {
	return s.liveMatchRepo.Delete(a)
}
