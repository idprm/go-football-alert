package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type PredictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) *PredictionRepository {
	return &PredictionRepository{
		db: db,
	}
}

type IPredictionRepository interface {
	Count(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int) (*entity.Prediction, error)
	Save(*entity.Prediction) (*entity.Prediction, error)
	Update(*entity.Prediction) (*entity.Prediction, error)
	UpdateByFixtureId(*entity.Prediction) (*entity.Prediction, error)
	Delete(*entity.Prediction) error
}

func (r *PredictionRepository) Count(fixtureId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Prediction{}).Where("fixture_id = ?", fixtureId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *PredictionRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var predictions []*entity.Prediction
	err := r.db.Scopes(Paginate(predictions, pagination, r.db)).Find(&predictions).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = predictions
	return pagination, nil
}

func (r *PredictionRepository) Get(fixtureId int) (*entity.Prediction, error) {
	var c entity.Prediction
	err := r.db.Where("fixture_id = ?", fixtureId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PredictionRepository) Save(c *entity.Prediction) (*entity.Prediction, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PredictionRepository) Update(c *entity.Prediction) (*entity.Prediction, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PredictionRepository) UpdateByFixtureId(c *entity.Prediction) (*entity.Prediction, error) {
	err := r.db.Where("fixture_id = ?", c.FixtureID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *PredictionRepository) Delete(c *entity.Prediction) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
