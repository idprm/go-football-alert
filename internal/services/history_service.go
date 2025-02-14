package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type HistoryService struct {
	historyRepo repository.IHistoryRepository
}

func NewHistoryService(historyRepo repository.IHistoryRepository) *HistoryService {
	return &HistoryService{
		historyRepo: historyRepo,
	}
}

type IHistoryService interface {
	IsHistory(int, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string) (*entity.History, error)
	Save(*entity.History) (*entity.History, error)
	Update(*entity.History) (*entity.History, error)
	Delete(*entity.History) error
}

func (s *HistoryService) IsHistory(serviceId int, msisdn string) bool {
	count, _ := s.historyRepo.Count(serviceId, msisdn)
	return count > 0
}

func (s *HistoryService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.historyRepo.GetAllPaginate(pagination)
}

func (s *HistoryService) Get(serviceId int, msisdn string) (*entity.History, error) {
	return s.historyRepo.Get(serviceId, msisdn)
}

func (s *HistoryService) Save(a *entity.History) (*entity.History, error) {
	return s.historyRepo.Save(a)
}

func (s *HistoryService) Update(a *entity.History) (*entity.History, error) {
	return s.historyRepo.Update(a)
}

func (s *HistoryService) Delete(a *entity.History) error {
	return s.historyRepo.Delete(a)
}
