package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type LiveMatchRepository struct {
	db *gorm.DB
}

func NewLiveMatchRepository(db *gorm.DB) *LiveMatchRepository {
	return &LiveMatchRepository{
		db: db,
	}
}

type ILiveMatchRepository interface {
	Count(int) (int64, error)
	CountByFixtureDate(time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllLiveMatchUSSD(int) ([]*entity.LiveMatch, error)
	Get(int) (*entity.LiveMatch, error)
	Save(*entity.LiveMatch) (*entity.LiveMatch, error)
	Update(*entity.LiveMatch) (*entity.LiveMatch, error)
	Delete(*entity.LiveMatch) error
}

func (r *LiveMatchRepository) Count(fixtureId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.LiveMatch{}).Where("fixture_id = ?", fixtureId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LiveMatchRepository) CountByFixtureDate(fixDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.LiveMatch{}).Where("DATE(fixture_date) = DATE(?)", fixDate).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LiveMatchRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var livematches []*entity.LiveMatch
	err := r.db.Scopes(Paginate(livematches, p, r.db)).Preload("Fixture.Home").Preload("Fixture.Away").Find(&livematches).Error
	if err != nil {
		return nil, err
	}
	p.Rows = livematches
	return p, nil
}

func (r *LiveMatchRepository) GetAllLiveMatchUSSD(page int) ([]*entity.LiveMatch, error) {
	var c []*entity.LiveMatch
	err := r.db.Where("DATE(fixture_date) = DATE(NOW())").Preload("Fixture.Home").Preload("Fixture.Away").Order("DATE(fixture_date) ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LiveMatchRepository) Get(id int) (*entity.LiveMatch, error) {
	var c entity.LiveMatch
	err := r.db.Where("id = ?", id).Preload("Home").Preload("Away").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LiveMatchRepository) Save(c *entity.LiveMatch) (*entity.LiveMatch, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LiveMatchRepository) Update(c *entity.LiveMatch) (*entity.LiveMatch, error) {
	err := r.db.Where("fixture_id = ?", c.FixtureID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LiveMatchRepository) Delete(c *entity.LiveMatch) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
