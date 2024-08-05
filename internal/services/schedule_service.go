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
	IsHome(string) bool
	GetAll() (*[]entity.Schedule, error)
	GetAllPaginate(int, int) (*[]entity.Schedule, error)
	GetById(int) (*entity.Schedule, error)
	GetBySlug(string) (*entity.Schedule, error)
	Save(*entity.Schedule) (*entity.Schedule, error)
	Update(*entity.Schedule) (*entity.Schedule, error)
	Delete(*entity.Schedule) error
}
