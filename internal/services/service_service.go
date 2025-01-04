package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type ServiceService struct {
	serviceRepo repository.IServiceRepository
}

func NewServiceService(serviceRepo repository.IServiceRepository) *ServiceService {
	return &ServiceService{
		serviceRepo: serviceRepo,
	}
}

type IServiceService interface {
	IsService(string) bool
	IsServiceById(int) bool
	IsServiceByPackage(string, string) bool
	GetAll() ([]*entity.Service, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByCategory(string) ([]*entity.Service, error)
	Get(string) (*entity.Service, error)
	GetById(int) (*entity.Service, error)
	GetByPackage(string, string) (*entity.Service, error)
	Save(*entity.Service) (*entity.Service, error)
	Update(*entity.Service) (*entity.Service, error)
	Delete(*entity.Service) error
}

func (s *ServiceService) IsService(code string) bool {
	count, _ := s.serviceRepo.Count(code)
	return count > 0
}

func (s *ServiceService) IsServiceById(id int) bool {
	count, _ := s.serviceRepo.CountById(id)
	return count > 0
}

func (s *ServiceService) IsServiceByPackage(category, pkg string) bool {
	count, _ := s.serviceRepo.CountByPackage(category, pkg)
	return count > 0
}

func (s *ServiceService) GetAll() ([]*entity.Service, error) {
	return s.serviceRepo.GetAll()
}

func (s *ServiceService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.serviceRepo.GetAllPaginate(pagination)
}

func (s *ServiceService) GetAllByCategory(category string) ([]*entity.Service, error) {
	return s.serviceRepo.GetAllByCategory(category)
}

func (s *ServiceService) Get(code string) (*entity.Service, error) {
	return s.serviceRepo.Get(code)
}

func (s *ServiceService) GetById(id int) (*entity.Service, error) {
	return s.serviceRepo.GetById(id)
}

func (s *ServiceService) GetByPackage(category, pkg string) (*entity.Service, error) {
	return s.serviceRepo.GetByPackage(category, pkg)
}

func (s *ServiceService) Save(a *entity.Service) (*entity.Service, error) {
	return s.serviceRepo.Save(a)
}

func (s *ServiceService) Update(a *entity.Service) (*entity.Service, error) {
	return s.serviceRepo.Update(a)
}

func (s *ServiceService) Delete(a *entity.Service) error {
	return s.serviceRepo.Delete(a)
}
