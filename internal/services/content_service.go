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
	IsContent(int, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string) (*entity.Content, error)
	Save(*entity.Content) (*entity.Content, error)
	Update(*entity.Content) (*entity.Content, error)
	Delete(*entity.Content) error
}

func (s *ContentService) IsContent(serviceId int, key string) bool {
	count, _ := s.contentRepo.Count(serviceId, key)
	return count > 0
}

func (s *ContentService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.contentRepo.GetAllPaginate(pagination)
}

func (s *ContentService) Get(serviceId int, key string) (*entity.Content, error) {
	return s.contentRepo.Get(serviceId, key)
}

func (s *ContentService) Save(a *entity.Content) (*entity.Content, error) {
	return s.contentRepo.Save(a)
}

func (s *ContentService) Update(a *entity.Content) (*entity.Content, error) {
	return s.contentRepo.Update(a)
}

func (s *ContentService) Delete(a *entity.Content) error {
	return s.contentRepo.Delete(a)
}
