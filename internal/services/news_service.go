package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type NewsService struct {
	newsRepo repository.INewsRepository
}

func NewNewsService(newsRepo repository.INewsRepository) *NewsService {
	return &NewsService{
		newsRepo: newsRepo,
	}
}

type INewsService interface {
	IsNews(int, int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
}

func (s *NewsService) IsNews(fixtureId, teamId int) bool {
	count, _ := s.newsRepo.Count(fixtureId, teamId)
	return count > 0
}

func (s *NewsService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.newsRepo.GetAllPaginate(pagination)
}

func (s *NewsService) Get(fixtureId, teamId int) (*entity.News, error) {
	return s.newsRepo.Get(fixtureId, teamId)
}

func (s *NewsService) Save(a *entity.News) (*entity.News, error) {
	return s.newsRepo.Save(a)
}

func (s *NewsService) Update(a *entity.News) (*entity.News, error) {
	return s.newsRepo.Update(a)
}

func (s *NewsService) Delete(a *entity.News) error {
	return s.newsRepo.Delete(a)
}
