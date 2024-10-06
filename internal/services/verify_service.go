package services

import (
	"strings"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type VerifyService struct {
	verifyRepo repository.IVerifyRepository
}

type IVerifyService interface {
	SetPIN(*entity.Verify) error
	SetCategory(*entity.Verify) error
	GetPIN(string) (*entity.Verify, error)
	GetCategory(string) (*entity.Verify, error)
}

func NewVerifyService(verifyRepo repository.IVerifyRepository) *VerifyService {
	return &VerifyService{
		verifyRepo: verifyRepo,
	}
}

func (s *VerifyService) SetPIN(t *entity.Verify) error {
	return s.verifyRepo.SetPIN(t)
}

func (s *VerifyService) SetCategory(t *entity.Verify) error {
	return s.verifyRepo.SetCategory(t)
}

func (s *VerifyService) GetPIN(v string) (*entity.Verify, error) {
	return s.verifyRepo.GetPIN(strings.ToLower(v))
}

func (s *VerifyService) GetCategory(v string) (*entity.Verify, error) {
	return s.verifyRepo.GetCategory(strings.ToLower(v))
}
