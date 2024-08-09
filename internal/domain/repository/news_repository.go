package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{
		db: db,
	}
}

type INewsRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
}

func (r *NewsRepository) Count(fixtureId, teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.News{}).Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *NewsRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var news []*entity.News
	err := r.db.Scopes(Paginate(news, pagination, r.db)).Find(&news).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = news
	return pagination, nil
}

func (r *NewsRepository) Get(fixtureId, teamId int) (*entity.News, error) {
	var c entity.News
	err := r.db.Where("fixture_id = ?", fixtureId).Where("team_id = ?", teamId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *NewsRepository) Save(c *entity.News) (*entity.News, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) Update(c *entity.News) (*entity.News, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) Delete(c *entity.News) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
