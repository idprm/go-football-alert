package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SMSAlerteService struct {
	smsAlerteRepo repository.ISMSAlerteRespository
}

func NewSMSAlerteService(smsAlerteRepo repository.ISMSAlerteRespository) *SMSAlerteService {
	return &SMSAlerteService{
		smsAlerteRepo: smsAlerteRepo,
	}
}

type ISMSAlerteService interface {
	ISMSAlerte(int64, int64) bool
	Save(*entity.SMSAlerte) (*entity.SMSAlerte, error)
}

func (s *SMSAlerteService) ISMSAlerte(subId, newsId int64) bool {
	count, _ := s.smsAlerteRepo.Count(subId, newsId)
	return count > 0
}

func (s *SMSAlerteService) Save(a *entity.SMSAlerte) (*entity.SMSAlerte, error) {
	return s.smsAlerteRepo.Save(a)
}
