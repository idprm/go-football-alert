package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type FixtureRepository struct {
	db *gorm.DB
}

func NewFixtureRepository(db *gorm.DB) *FixtureRepository {
	return &FixtureRepository{
		db: db,
	}
}

type IFixtureRepository interface {
	Count(int) (int64, error)
	CountByPrimaryId(int) (int64, error)
	CountByFixtureDate(time.Time) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllCurrent() ([]*entity.Fixture, error)
	GetAllLiveMatch() ([]*entity.Fixture, error)
	GetAllLiveMatchTodayUSSD(int) ([]*entity.Fixture, error)
	GetAllLiveMatchLaterUSSD(int) ([]*entity.Fixture, error)
	GetAllScheduleUSSD(int) ([]*entity.Fixture, error)
	GetAllByLeagueIdUSSD(int, int) ([]*entity.Fixture, error)
	GetAllByFixtureDate(time.Time) ([]*entity.Fixture, error)
	Get(int) (*entity.Fixture, error)
	GetByPrimaryId(int) (*entity.Fixture, error)
	Save(*entity.Fixture) (*entity.Fixture, error)
	Update(*entity.Fixture) (*entity.Fixture, error)
	UpdateByPrimaryId(*entity.Fixture) (*entity.Fixture, error)
	Delete(*entity.Fixture) error
}

func (r *FixtureRepository) Count(id int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Fixture{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *FixtureRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Fixture{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *FixtureRepository) CountByFixtureDate(fixDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Fixture{}).Where("DATE(fixture_date) >= DATE(?)", fixDate).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *FixtureRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var fixtures []*entity.Fixture
	err := r.db.Scopes(Paginate(fixtures, p, r.db)).Preload("Home").Preload("Away").Find(&fixtures).Error
	if err != nil {
		return nil, err
	}
	p.Rows = fixtures
	return p, nil
}

func (r *FixtureRepository) GetAllCurrent() ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) >= DATE(NOW())").Order("DATE(fixture_date) ASC").Preload("Home").Preload("Away").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllLiveMatch() ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) = DATE(NOW())").Order("DATE(fixture_date) ASC").Preload("Home").Preload("Away").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllLiveMatchTodayUSSD(page int) ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) = DATE(NOW())").Preload("Home").Preload("Away").Order("DATE(fixture_date) ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllLiveMatchLaterUSSD(page int) ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) > DATE(NOW())").Preload("Home").Preload("Away").Order("DATE(fixture_date) ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllScheduleUSSD(page int) ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) BETWEEN DATE(NOW()) AND DATE(NOW() + INTERVAL 30 DAY)").Preload("Home").Preload("Away").Order("DATE(fixture_date) ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllByLeagueIdUSSD(leagueId, page int) ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("league_id = ?", leagueId).Where("DATE(fixture_date) BETWEEN DATE(NOW()) AND DATE(NOW() + INTERVAL 30 DAY)").Preload("Home").Preload("Away").Order("DATE(fixture_date) ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) GetAllByFixtureDate(fixDate time.Time) ([]*entity.Fixture, error) {
	var c []*entity.Fixture
	err := r.db.Where("DATE(fixture_date) BETWEEN DATE(?) AND DATE(? + INTERVAL 5 DAY)", fixDate, fixDate).Preload("Home").Preload("Away").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) Get(id int) (*entity.Fixture, error) {
	var c entity.Fixture
	err := r.db.Where("id = ?", id).Preload("Home").Preload("Away").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FixtureRepository) GetByPrimaryId(primaryId int) (*entity.Fixture, error) {
	var c entity.Fixture
	err := r.db.Where("primary_id = ?", primaryId).Preload("Home").Preload("Away").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *FixtureRepository) Save(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) Update(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) UpdateByPrimaryId(c *entity.Fixture) (*entity.Fixture, error) {
	err := r.db.Where("primary_id = ?", c.PrimaryID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *FixtureRepository) Delete(c *entity.Fixture) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
