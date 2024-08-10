package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SeasonService struct {
	seasonRepo repository.ISeasonRepository
}

func NewSeasonService(seasonRepo repository.ISeasonRepository) *SeasonService {
	return &SeasonService{
		seasonRepo: seasonRepo,
	}
}

type ISeasonService interface {
	IsSeason(string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Season, error)
	Save(*entity.Season) (*entity.Season, error)
	Update(*entity.Season) (*entity.Season, error)
	Delete(*entity.Season) error
}

func (s *SeasonService) IsSeason(slug string) bool {
	count, _ := s.seasonRepo.Count(slug)
	return count > 0
}

func (s *SeasonService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.seasonRepo.GetAllPaginate(pagination)
}

func (s *SeasonService) Get(slug string) (*entity.Season, error) {
	return s.seasonRepo.Get(slug)
}

func (s *SeasonService) Save(a *entity.Season) (*entity.Season, error) {
	return s.seasonRepo.Save(a)
}

func (s *SeasonService) Update(a *entity.Season) (*entity.Season, error) {
	return s.seasonRepo.Update(a)
}

func (s *SeasonService) Delete(a *entity.Season) error {
	return s.seasonRepo.Delete(a)
}
