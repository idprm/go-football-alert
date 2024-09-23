package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

type IServiceRepository interface {
	Count(string) (int64, error)
	CountById(int) (int64, error)
	CountByPackage(string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Service, error)
	GetById(int) (*entity.Service, error)
	GetByPackage(string, string) (*entity.Service, error)
	Save(*entity.Service) (*entity.Service, error)
	Update(*entity.Service) (*entity.Service, error)
	Delete(*entity.Service) error
}

func (r *ServiceRepository) Count(code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Service{}).Where("code = ?", code).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ServiceRepository) CountById(id int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Service{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ServiceRepository) CountByPackage(category, pkg string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Service{}).Where("category = ?", category).Where("package = ?", pkg).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ServiceRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var services []*entity.Service
	err := r.db.Scopes(Paginate(services, pagination, r.db)).Find(&services).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = services
	return pagination, nil
}

func (r *ServiceRepository) Get(code string) (*entity.Service, error) {
	var c entity.Service
	err := r.db.Where("code = ?", code).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ServiceRepository) GetById(id int) (*entity.Service, error) {
	var c entity.Service
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ServiceRepository) GetByPackage(category, pkg string) (*entity.Service, error) {
	var c entity.Service
	err := r.db.Where("category = ?", category).Where("package = ?", pkg).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ServiceRepository) Save(c *entity.Service) (*entity.Service, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ServiceRepository) Update(c *entity.Service) (*entity.Service, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ServiceRepository) Delete(c *entity.Service) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
