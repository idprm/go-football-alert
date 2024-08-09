package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

type IScheduleRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.Schedule, error)
	Save(*entity.Schedule) (*entity.Schedule, error)
	Update(*entity.Schedule) (*entity.Schedule, error)
	Delete(*entity.Schedule) error
}

func (r *ScheduleRepository) Count(fixtureId, subscriptionId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Schedule{}).Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subscriptionId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScheduleRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var schedules []*entity.Schedule
	err := r.db.Scopes(Paginate(schedules, pagination, r.db)).Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = schedules
	return pagination, nil
}

func (r *ScheduleRepository) Get(fixtureId, subscriptionId int) (*entity.Schedule, error) {
	var c entity.Schedule
	err := r.db.Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subscriptionId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ScheduleRepository) Save(c *entity.Schedule) (*entity.Schedule, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ScheduleRepository) Update(c *entity.Schedule) (*entity.Schedule, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ScheduleRepository) Delete(c *entity.Schedule) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
