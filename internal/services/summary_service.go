package services

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SummaryService struct {
	summaryRepo repository.ISummaryRepository
}

func NewSummaryService(
	summaryRepo repository.ISummaryRepository,
) *SummaryService {
	return &SummaryService{
		summaryRepo: summaryRepo,
	}
}

type ISummaryService interface {
	IsSummary(int, time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, time.Time) (*entity.Summary, error)
	GetActiveSub(time.Time, time.Time) (int, error)
	GetSub(time.Time, time.Time) (int, error)
	GetUnSub(time.Time, time.Time) (int, error)
	GetRenewal(time.Time, time.Time) (int, error)
	GetRevenue(time.Time, time.Time) (float64, error)
	Save(*entity.Summary) (*entity.Summary, error)
	Update(*entity.Summary) (*entity.Summary, error)
	UpdateRetry(*entity.Summary) (*entity.Summary, error)
	Delete(*entity.Summary) error
}

func (s *SummaryService) IsSummary(serviceId int, date time.Time) bool {
	count, _ := s.summaryRepo.Count(serviceId, date)
	return count > 0
}

func (s *SummaryService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.summaryRepo.GetAllPaginate(pagination)
}

func (s *SummaryService) Get(serviceId int, date time.Time) (*entity.Summary, error) {
	return s.summaryRepo.Get(serviceId, date)
}

func (s *SummaryService) GetActiveSub(start, end time.Time) (int, error) {
	return s.summaryRepo.GetActiveSub(start, end)
}

func (s *SummaryService) GetSub(start, end time.Time) (int, error) {
	return s.summaryRepo.GetSub(start, end)
}

func (s *SummaryService) GetUnSub(start, end time.Time) (int, error) {
	return s.summaryRepo.GetUnSub(start, end)
}

func (s *SummaryService) GetRenewal(start, end time.Time) (int, error) {
	return s.summaryRepo.GetRenewal(start, end)
}

func (s *SummaryService) GetRevenue(start, end time.Time) (float64, error) {
	return s.summaryRepo.GetRevenue(start, end)
}

func (s *SummaryService) Save(a *entity.Summary) (*entity.Summary, error) {
	if s.IsSummary(a.ServiceID, a.CreatedAt) {
		summ, err := s.Get(a.ServiceID, a.CreatedAt)
		if err != nil {
			return nil, err
		}
		return s.summaryRepo.Update(
			&entity.Summary{
				ServiceID:          a.ServiceID,
				CreatedAt:          a.CreatedAt,
				TotalSub:           summ.TotalSub + a.TotalSub,
				TotalUnsub:         summ.TotalUnsub + a.TotalUnsub,
				TotalRenewal:       summ.TotalRenewal + a.TotalRenewal,
				TotalChargeSuccess: summ.TotalChargeSuccess + a.TotalChargeSuccess,
				TotalChargeFailed:  summ.TotalChargeFailed + a.TotalChargeFailed,
				TotalRevenue:       summ.TotalRevenue + a.TotalRevenue,
			},
		)
	}
	return s.summaryRepo.Save(a)
}

func (s *SummaryService) Update(a *entity.Summary) (*entity.Summary, error) {
	return s.summaryRepo.Update(a)
}

func (s *SummaryService) UpdateRetry(a *entity.Summary) (*entity.Summary, error) {
	if s.IsSummary(a.ServiceID, a.CreatedAt) {
		summ, err := s.Get(a.ServiceID, a.CreatedAt)
		if err != nil {
			return nil, err
		}
		return s.summaryRepo.Update(
			&entity.Summary{
				ServiceID:          a.ServiceID,
				CreatedAt:          a.CreatedAt,
				TotalChargeSuccess: summ.TotalChargeSuccess + a.TotalChargeSuccess,
				TotalChargeFailed:  summ.TotalChargeFailed - 1,
				TotalRevenue:       summ.TotalRevenue + a.TotalRevenue,
			},
		)
	}
	return s.summaryRepo.Save(a)
}

func (s *SummaryService) Delete(a *entity.Summary) error {
	return s.summaryRepo.Delete(a)
}
