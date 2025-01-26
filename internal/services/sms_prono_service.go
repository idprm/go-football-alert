package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SMSPronoService struct {
	smsPronoRepo repository.ISMSPronoRespository
}

func NewSMSPronoService(smsPronoRepo repository.ISMSPronoRespository) *SMSPronoService {
	return &SMSPronoService{
		smsPronoRepo: smsPronoRepo,
	}
}

type ISMSPronoService interface {
	ISMSProno(int64, int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.SMSProno) (*entity.SMSProno, error)
}

func (s *SMSPronoService) ISMSProno(subId, pronoId int64) bool {
	count, _ := s.smsPronoRepo.Count(subId, pronoId)
	return count > 0
}

func (s *SMSPronoService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.smsPronoRepo.GetAllPaginate(pagination)
}

func (s *SMSPronoService) Save(a *entity.SMSProno) (*entity.SMSProno, error) {
	return s.smsPronoRepo.Save(a)
}
