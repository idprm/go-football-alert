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
	GetSubByMonth(time.Time) (int, error)
	GetUnsubByMonth(time.Time) (int, error)
	GetRenewalByMonth(time.Time) (int, error)
	GetRevenueByMonth(time.Time) (float64, error)
	Save(*entity.Summary) (*entity.Summary, error)
	Update(*entity.Summary) (*entity.Summary, error)
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

func (s *SummaryService) GetSubByMonth(date time.Time) (int, error) {
	return s.summaryRepo.GetSubByMonth(date)
}

func (s *SummaryService) GetUnsubByMonth(date time.Time) (int, error) {
	return s.summaryRepo.GetUnsubByMonth(date)
}

func (s *SummaryService) GetRenewalByMonth(date time.Time) (int, error) {
	return s.summaryRepo.GetRenewalByMonth(date)
}

func (s *SummaryService) GetRevenueByMonth(date time.Time) (float64, error) {
	return s.summaryRepo.GetRevenueByMonth(date)
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

func (s *SummaryService) Delete(a *entity.Summary) error {
	return s.summaryRepo.Delete(a)
}
