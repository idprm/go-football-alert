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
	CountUnlocked(string, string) (int64, error)
	CountLocked(string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string, string) (*entity.Schedule, error)
	Save(*entity.Schedule) (*entity.Schedule, error)
	Update(*entity.Schedule) error
	UpdateLocked(*entity.Schedule) error
}

func (r *ScheduleRepository) CountUnlocked(key, hour string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Schedule{}).Where("name = ?", key).Where("DATE_FORMAT(publish_at, '%H:%i') = ?", hour).Where("is_unlocked = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScheduleRepository) CountLocked(key, hour string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Schedule{}).Where("name = ?", key).Where("DATE_FORMAT(unlocked_at, '%H:%i') = ?", hour).Where("is_unlocked = ?", false).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScheduleRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var schedules []*entity.Schedule
	err := r.db.Where("UPPER(name) LIKE UPPER(?)", "%"+p.GetSearch()+"%").Scopes(Paginate(schedules, p, r.db)).Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	p.Rows = schedules
	return p, nil
}

func (r *ScheduleRepository) Get(key, hour string) (*entity.Schedule, error) {
	var c entity.Schedule
	err := r.db.Where("name = ?", key).Where("DATE_FORMAT(publish_at, '%H:%i') = ?", hour).Take(&c).Error
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

func (r *ScheduleRepository) Update(c *entity.Schedule) error {
	err := r.db.Where("name = ?", c.Name).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ScheduleRepository) UpdateLocked(c *entity.Schedule) error {
	err := r.db.Where("name = ?", c.Name).Update("is_unlocked", false).Error
	if err != nil {
		return err
	}
	return nil
}
