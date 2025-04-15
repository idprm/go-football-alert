package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(
	userRepo repository.IUserRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

type IUserService interface {
	IsEmail(string) bool
	IsValid(string, string) bool
	Get(string, string) (*entity.User, error)
	Save(*entity.User) error
	Update(*entity.User) error
	Delete(*entity.User) error
}

func (s *UserService) IsEmail(email string) bool {
	count, _ := s.userRepo.CountByEmail(email)
	return count > 0
}

func (s *UserService) IsValid(email, password string) bool {
	count, _ := s.userRepo.CountByEmail(email)
	return count > 0
}

func (s *UserService) Get(email, password string) (*entity.User, error) {
	return s.userRepo.Get(email, password)
}

func (s *UserService) Save(a *entity.User) error {
	return s.userRepo.Save(a)
}

func (s *UserService) Update(a *entity.User) error {
	return s.userRepo.Update(a)
}

func (s *UserService) Delete(a *entity.User) error {
	return s.userRepo.Delete(a)
}
