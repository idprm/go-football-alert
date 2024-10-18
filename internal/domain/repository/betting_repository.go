package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type BettingRepository struct {
	db *gorm.DB
}

func NewBettingRepository(db *gorm.DB) *BettingRepository {
	return &BettingRepository{
		db: db,
	}
}

type IBettingRepository interface {
	Count(int64, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.Betting, error)
	Save(*entity.Betting) (*entity.Betting, error)
	Update(*entity.Betting) (*entity.Betting, error)
	Delete(*entity.Betting) error
}

func (r *BettingRepository) Count(fixtureId, subId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Betting{}).Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *BettingRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var contents []*entity.Content
	err := r.db.Where("UPPER(name) LIKE UPPER(?)", "%"+p.GetSearch()+"%").Scopes(Paginate(contents, p, r.db)).Find(&contents).Error
	if err != nil {
		return nil, err
	}
	p.Rows = contents
	return p, nil
}

func (r *BettingRepository) Get(fixtureId, subId int64) (*entity.Betting, error) {
	var content entity.Betting
	err := r.db.Where("fixture_id = ?", fixtureId).Where("subscription_id = ?", subId).Take(&content).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *BettingRepository) Save(c *entity.Betting) (*entity.Betting, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *BettingRepository) Update(c *entity.Betting) (*entity.Betting, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *BettingRepository) Delete(c *entity.Betting) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
