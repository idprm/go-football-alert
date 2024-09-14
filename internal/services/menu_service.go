package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type MenuService struct {
	menuRepo repository.IMenuRepository
}

type IMenuService interface {
	IsKeyPress(string) bool
	IsAction(string) bool
	GetAll() ([]*entity.Menu, error)
	GetMenuByKeyPress(string) (*entity.Menu, error)
	GetMenuByParentId(int) ([]*entity.Menu, error)
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

func (s *MenuService) IsKeyPress(keyPress string) bool {
	count, _ := s.menuRepo.CountByKeyPress(keyPress)
	return count > 0
}

func (s *MenuService) IsAction(keyPress string) bool {
	count, _ := s.menuRepo.CountByAction(keyPress)
	return count > 0
}

func (s *MenuService) GetAll() ([]*entity.Menu, error) {
	return s.menuRepo.GetAll()
}

func (s *MenuService) GetMenuByKeyPress(keyPress string) (*entity.Menu, error) {
	return s.menuRepo.GetMenuByKeyPress(keyPress)
}

func (s *MenuService) GetMenuByParentId(parentId int) ([]*entity.Menu, error) {
	return s.menuRepo.GetMenuByParentId(parentId)
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
