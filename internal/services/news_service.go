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
	IsNews(string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllUSSD(int) ([]*entity.News, error)
	GetByTeamUSSD(string, int) (*entity.News, error)
	Get(string, string) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
}

func (s *NewsService) IsNews(slug, pubAt string) bool {
	count, _ := s.newsRepo.Count(slug, pubAt)
	return count > 0
}

func (s *NewsService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.newsRepo.GetAllPaginate(pagination)
}

func (s *NewsService) GetAllUSSD(page int) ([]*entity.News, error) {
	return s.newsRepo.GetAllUSSD(page)
}

func (s *NewsService) GetByTeamUSSD(pubAt string, teamId int) (*entity.News, error) {
	return s.newsRepo.GetByTeamUSSD(pubAt, teamId)
}

func (s *NewsService) Get(slug, pubAt string) (*entity.News, error) {
	return s.newsRepo.Get(slug, pubAt)
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
