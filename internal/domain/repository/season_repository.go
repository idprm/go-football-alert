package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{
		db: db,
	}
}

type ISeasonRepository interface {
	Count(string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Season, error)
	Save(*entity.Season) (*entity.Season, error)
	Update(*entity.Season) (*entity.Season, error)
	Delete(*entity.Season) error
}

func (r *SeasonRepository) Count(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Season{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SeasonRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var seasons []*entity.Season
	err := r.db.Scopes(Paginate(seasons, pagination, r.db)).Find(&seasons).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = seasons
	return pagination, nil
}

func (r *SeasonRepository) Get(slug string) (*entity.Season, error) {
	var c entity.Season
	err := r.db.Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SeasonRepository) Save(c *entity.Season) (*entity.Season, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SeasonRepository) Update(c *entity.Season) (*entity.Season, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SeasonRepository) Delete(c *entity.Season) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
