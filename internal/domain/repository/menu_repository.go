package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewMenuRepository(
	db *gorm.DB,
	rds *redis.Client,
) *MenuRepository {
	return &MenuRepository{
		db:  db,
		rds: rds,
	}
}

type IMenuRepository interface {
	CountBySlug(string) (int64, error)
	CountByKeyPress(string) (int64, error)
	GetAll() ([]*entity.Menu, error)
	GetBySlug(string) (*entity.Menu, error)
	GetByKeyPress(string) (*entity.Menu, error)
	Save(*entity.Menu) (*entity.Menu, error)
	Update(*entity.Menu) (*entity.Menu, error)
	Delete(*entity.Menu) error
}

func (r *MenuRepository) CountBySlug(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Menu{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MenuRepository) CountByKeyPress(keyPress string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Menu{}).Where("key_press = ?", keyPress).Count(&count).Error
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

func (r *MenuRepository) GetByKeyPress(keyPress string) (*entity.Menu, error) {
	var menu *entity.Menu
	err := r.db.Where(&entity.Menu{IsActive: true}).Take(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *MenuRepository) GetBySlug(slug string) (*entity.Menu, error) {
	var menu *entity.Menu
	err := r.db.Where(&entity.Menu{Slug: slug, IsActive: true}).Take(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
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
