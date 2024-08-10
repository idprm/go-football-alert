package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type ScheduleService struct {
	scheduleRepo repository.IScheduleRepository
}

func NewScheduleService(scheduleRepo repository.IScheduleRepository) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

type IScheduleService interface {
	IsUnlocked(string, string) bool
	IsLocked(string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Schedule, error)
	Save(*entity.Schedule) (*entity.Schedule, error)
	Update(*entity.Schedule) error
}

func (s *ScheduleService) IsUnlocked(key, hour string) bool {
	count, _ := s.scheduleRepo.CountUnlocked(key, hour)
	return count > 0
}

func (s *ScheduleService) IsLocked(key, hour string) bool {
	count, _ := s.scheduleRepo.CountLocked(key, hour)
	return count > 0
}

func (s *ScheduleService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.scheduleRepo.GetAllPaginate(pagination)
}

func (s *ScheduleService) Get(key, hour string) (*entity.Schedule, error) {
	return s.scheduleRepo.Get(key, hour)
}

func (s *ScheduleService) Save(a *entity.Schedule) (*entity.Schedule, error) {
	return s.scheduleRepo.Save(a)
}

func (s *ScheduleService) Update(a *entity.Schedule) error {
	return s.scheduleRepo.Update(a)
}
