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
	IsHome(string) bool
	GetAll() (*[]entity.News, error)
	GetAllPaginate(int, int) (*[]entity.News, error)
	GetById(int) (*entity.News, error)
	GetBySlug(string) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
}
