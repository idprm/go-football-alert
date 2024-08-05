package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type ContentService struct {
	contentRepo repository.IContentRepository
}

func NewContentService(contentRepo repository.IContentRepository) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
	}
}

type IContentService interface {
	IsAway(string) bool
	GetAll() (*[]entity.Content, error)
	GetAllPaginate(int, int) (*[]entity.Content, error)
	GetById(int) (*entity.Content, error)
	GetBySlug(string) (*entity.Content, error)
	Save(*entity.Content) (*entity.Content, error)
	Update(*entity.Content) (*entity.Content, error)
	Delete(*entity.Content) error
}
