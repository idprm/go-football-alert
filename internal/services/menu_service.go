package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type MenuService struct {
	menuRepo repository.IMenuRepository
}

type IMenuService interface {
	IsSlug(string) bool
	IsKeyPress(string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAll() ([]*entity.Menu, error)
	GetBySlug(string) (*entity.Menu, error)
	GetByKeyPress(string) (*entity.Menu, error)
	Save(*entity.Menu) (*entity.Menu, error)
	Update(*entity.Menu) (*entity.Menu, error)
	Delete(*entity.Menu) error
}

func NewMenuService(
	menuRepo repository.IMenuRepository,
) *MenuService {
	return &MenuService{
		menuRepo: menuRepo,
	}
}

func (s *MenuService) IsSlug(keyPress string) bool {
	count, _ := s.menuRepo.CountBySlug(keyPress)
	return count > 0
}

func (s *MenuService) IsKeyPress(keyPress string) bool {
	count, _ := s.menuRepo.CountByKeyPress(keyPress)
	return count > 0
}

func (s *MenuService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.menuRepo.GetAllPaginate(pagination)
}

func (s *MenuService) GetAll() ([]*entity.Menu, error) {
	return s.menuRepo.GetAll()
}

func (s *MenuService) GetBySlug(slug string) (*entity.Menu, error) {
	return s.menuRepo.GetBySlug(slug)
}

func (s *MenuService) GetByKeyPress(keyPress string) (*entity.Menu, error) {
	return s.menuRepo.GetByKeyPress(keyPress)
}

func (s *MenuService) Save(e *entity.Menu) (*entity.Menu, error) {
	return s.menuRepo.Save(e)
}

func (s *MenuService) Update(e *entity.Menu) (*entity.Menu, error) {
	return s.menuRepo.Update(e)
}

func (s *MenuService) Delete(e *entity.Menu) error {
	return s.menuRepo.Delete(e)
}
