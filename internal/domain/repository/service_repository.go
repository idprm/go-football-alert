package repository

import (
	"strings"

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
	GetAllByCategory(string) ([]*entity.Service, error)
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

func (r *ServiceRepository) CountByPackage(cat, pkg string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Service{}).Where("category = ?", strings.ToUpper(cat)).Where("package = ?", strings.ToLower(pkg)).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ServiceRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var services []*entity.Service
	err := r.db.Where("UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(Paginate(services, p, r.db)).Find(&services).Error
	if err != nil {
		return nil, err
	}
	p.Rows = services
	return p, nil
}

func (r *ServiceRepository) GetAllByCategory(cat string) ([]*entity.Service, error) {
	var services []*entity.Service
	err := r.db.Where("category = ?", strings.ToUpper(cat)).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (r *ServiceRepository) Get(code string) (*entity.Service, error) {
	var c entity.Service
	err := r.db.Where("code = ?", strings.ToUpper(code)).Take(&c).Error
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

func (r *ServiceRepository) GetByPackage(cat, pkg string) (*entity.Service, error) {
	var c entity.Service
	err := r.db.Where("category = ?", strings.ToUpper(cat)).Where("package = ?", strings.ToLower(pkg)).Take(&c).Error
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
