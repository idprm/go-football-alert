package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SMSActuService struct {
	smsActuRepo repository.ISMSActuRespository
}

func NewSMSActuService(smsActuRepo repository.ISMSActuRespository) *SMSActuService {
	return &SMSActuService{
		smsActuRepo: smsActuRepo,
	}
}

type ISMSActuService interface {
	ISMSActu(string, int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.SMSActu) (*entity.SMSActu, error)
}

func (s *SMSActuService) ISMSActu(msisdn string, newsId int64) bool {
	count, _ := s.smsActuRepo.Count(msisdn, newsId)
	return count > 0
}

func (s *SMSActuService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.smsActuRepo.GetAllPaginate(p)
}

func (s *SMSActuService) Save(a *entity.SMSActu) (*entity.SMSActu, error) {
	return s.smsActuRepo.Save(a)
}
