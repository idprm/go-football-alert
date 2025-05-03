package services

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SummaryDashboardService struct {
	summaryDashboardRepo repository.ISummaryDashboardRepository
}

type SummaryRevenueService struct {
	summaryRevenueRepo repository.ISummaryRevenueRepository
}

type SummaryTotalDailyService struct {
	summaryTotalDailyRepo repository.ISummaryTotalDailyRepository
}

func NewSummaryDashboardService(
	summaryDashboardRepo repository.ISummaryDashboardRepository,
) *SummaryDashboardService {
	return &SummaryDashboardService{
		summaryDashboardRepo: summaryDashboardRepo,
	}
}

func NewSummaryRevenueService(
	summaryRevenueRepo repository.ISummaryRevenueRepository,
) *SummaryRevenueService {
	return &SummaryRevenueService{
		summaryRevenueRepo: summaryRevenueRepo,
	}
}

func NewSummaryTotalDailyService(
	summaryTotalDailyRepo repository.ISummaryTotalDailyRepository,
) *SummaryTotalDailyService {
	return &SummaryTotalDailyService{
		summaryTotalDailyRepo: summaryTotalDailyRepo,
	}
}

type ISummaryDashboardService interface {
	IsSummary(time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(time.Time) (*entity.SummaryDashboard, error)
	Save(*entity.SummaryDashboard) error
	Update(*entity.SummaryDashboard) error
	Delete(*entity.SummaryDashboard) error
	GetTotalActiveSub() (int, error)
	GetTotalRevenue() (float64, error)
}

type ISummaryRevenueService interface {
	IsSummary(time.Time, string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByRange(*model.RangeDateRequest) ([]*entity.SummaryRevenue, error)
	Get(time.Time, string, string) (*entity.SummaryRevenue, error)
	Save(*entity.SummaryRevenue) error
	Update(*entity.SummaryRevenue) error
	Delete(*entity.SummaryRevenue) error
	SelectRevenue() (*[]entity.SummaryRevenue, error)
}

type ISummaryTotalDailyService interface {
	IsSummary(time.Time) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByRange(*model.RangeDateRequest) ([]*entity.SummaryTotalDaily, error)
	Get(time.Time) (*entity.SummaryTotalDaily, error)
	Save(*entity.SummaryTotalDaily) error
	Update(*entity.SummaryTotalDaily) error
	Delete(*entity.SummaryTotalDaily) error
	SelectTotalDaily() (*[]entity.SummaryTotalDaily, error)
}

func (s *SummaryDashboardService) IsSummary(date time.Time) bool {
	count, _ := s.summaryDashboardRepo.Count(date)
	return count > 0
}

func (s *SummaryDashboardService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.summaryDashboardRepo.GetAllPaginate(p)
}

func (s *SummaryDashboardService) Get(date time.Time) (*entity.SummaryDashboard, error) {
	return s.summaryDashboardRepo.Get(date)
}

func (s *SummaryDashboardService) Save(a *entity.SummaryDashboard) error {
	if s.IsSummary(a.CreatedAt) {
		return s.Update(a)
	}
	return s.summaryDashboardRepo.Save(a)
}

func (s *SummaryDashboardService) Update(a *entity.SummaryDashboard) error {
	return s.summaryDashboardRepo.Update(a)
}

func (s *SummaryDashboardService) Delete(a *entity.SummaryDashboard) error {
	return s.summaryDashboardRepo.Delete(a)
}

func (s *SummaryDashboardService) GetTotalActiveSub() (int, error) {
	return s.summaryDashboardRepo.GetTotalActiveSub()
}

func (s *SummaryDashboardService) GetTotalRevenue() (float64, error) {
	return s.summaryDashboardRepo.GetTotalRevenue()
}

func (s *SummaryRevenueService) IsSummary(date time.Time, subject, status string) bool {
	count, _ := s.summaryRevenueRepo.Count(date, subject, status)
	return count > 0
}

func (s *SummaryRevenueService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.summaryRevenueRepo.GetAllPaginate(p)
}

func (s *SummaryRevenueService) GetAllByRange(p *model.RangeDateRequest) ([]*entity.SummaryRevenue, error) {
	return s.summaryRevenueRepo.GetAllByRange(p)
}

func (s *SummaryRevenueService) Get(date time.Time, subject, status string) (*entity.SummaryRevenue, error) {
	return s.summaryRevenueRepo.Get(date, subject, status)
}

func (s *SummaryRevenueService) Save(a *entity.SummaryRevenue) error {
	if s.IsSummary(a.CreatedAt, a.Subject, a.Status) {
		return s.Update(a)
	}
	return s.summaryRevenueRepo.Save(a)
}

func (s *SummaryRevenueService) Update(a *entity.SummaryRevenue) error {
	return s.summaryRevenueRepo.Update(a)
}

func (s *SummaryRevenueService) Delete(a *entity.SummaryRevenue) error {
	return s.summaryRevenueRepo.Delete(a)
}

func (s *SummaryRevenueService) SelectRevenue() (*[]entity.SummaryRevenue, error) {
	return s.summaryRevenueRepo.SelectRevenue()
}

func (s *SummaryTotalDailyService) IsSummary(date time.Time) bool {
	count, _ := s.summaryTotalDailyRepo.Count(date)
	return count > 0
}

func (s *SummaryTotalDailyService) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	return s.summaryTotalDailyRepo.GetAllPaginate(p)
}

func (s *SummaryTotalDailyService) GetAllByRange(p *model.RangeDateRequest) ([]*entity.SummaryTotalDaily, error) {
	return s.summaryTotalDailyRepo.GetAllByRange(p)
}

func (s *SummaryTotalDailyService) Get(date time.Time) (*entity.SummaryTotalDaily, error) {
	return s.summaryTotalDailyRepo.Get(date)
}

func (s *SummaryTotalDailyService) Save(a *entity.SummaryTotalDaily) error {
	if s.IsSummary(a.CreatedAt) {
		return s.Update(a)
	}
	return s.summaryTotalDailyRepo.Save(a)
}

func (s *SummaryTotalDailyService) Update(a *entity.SummaryTotalDaily) error {
	return s.summaryTotalDailyRepo.Update(a)
}
func (s *SummaryTotalDailyService) Delete(a *entity.SummaryTotalDaily) error {
	return s.summaryTotalDailyRepo.Delete(a)
}

func (s *SummaryTotalDailyService) SelectTotalDaily() (*[]entity.SummaryTotalDaily, error) {
	return s.summaryTotalDailyRepo.SelectTotalDaily()
}
