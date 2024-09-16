package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type UssdService struct {
	ussdRepo repository.IUssdRepository
}

type IUssdService interface {
	GetAll() ([]*entity.Ussd, error)
	Save(*entity.Ussd) (*entity.Ussd, error)
	Update(*entity.Ussd) (*entity.Ussd, error)
	Delete(*entity.Ussd) error
	Set(*entity.Ussd) error
	Get(string) (*entity.Ussd, error)
}

func NewUssdService(
	ussdRepo repository.IUssdRepository,
) *UssdService {
	return &UssdService{
		ussdRepo: ussdRepo,
	}
}

func (s *UssdService) GetAll() ([]*entity.Ussd, error) {
	return s.ussdRepo.GetAll()
}

func (s *UssdService) Save(e *entity.Ussd) (*entity.Ussd, error) {
	return s.ussdRepo.Save(e)
}

func (s *UssdService) Update(e *entity.Ussd) (*entity.Ussd, error) {
	return s.ussdRepo.Update(e)
}

func (s *UssdService) Delete(e *entity.Ussd) error {
	return s.ussdRepo.Delete(e)
}

func (s *UssdService) Set(e *entity.Ussd) error {
	return s.ussdRepo.Set(e)
}

func (s *UssdService) Get(msisdn string) (*entity.Ussd, error) {
	return s.ussdRepo.Get(msisdn)
}
