package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}

type IMenuRepository interface {
	CountByKeyPress(string) (int64, error)
	CountByAction(string) (int64, error)
	GetAll() ([]*entity.Menu, error)
	GetMenuByKeyPress(string) (*entity.Menu, error)
	GetMenuByParentId(int) ([]*entity.Menu, error)
	Save(*entity.Menu) (*entity.Menu, error)
	Update(*entity.Menu) (*entity.Menu, error)
	Delete(*entity.Menu) error
}

func (r *MenuRepository) CountByKeyPress(keyPress string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Menu{}).Where("key_press = ?", keyPress).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MenuRepository) CountByAction(keyPress string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Menu{}).Where("key_press = ?", keyPress).Where("action != ?", "").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MenuRepository) GetAll() ([]*entity.Menu, error) {
	var menus []*entity.Menu
	err := r.db.Where(&entity.Menu{IsActive: true}).Where("parent_id = 0").Order("parent_id ASC").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepository) GetMenuByKeyPress(keyPress string) (*entity.Menu, error) {
	var menu *entity.Menu
	err := r.db.Where(&entity.Menu{KeyPress: keyPress, IsActive: true}).Take(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *MenuRepository) GetMenuByParentId(parentId int) ([]*entity.Menu, error) {
	var menus []*entity.Menu
	err := r.db.Where(&entity.Menu{ParentID: parentId, IsActive: true}).Order("child ASC").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepository) Save(e *entity.Menu) (*entity.Menu, error) {
	err := r.db.Create(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *MenuRepository) Update(e *entity.Menu) (*entity.Menu, error) {
	err := r.db.Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *MenuRepository) Delete(e *entity.Menu) error {
	err := r.db.Delete(&e, e.ID).Error
	if err != nil {
		return err
	}
	return nil
}
