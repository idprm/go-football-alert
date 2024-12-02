package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type LeagueRepository struct {
	db *gorm.DB
}

func NewLeagueRepository(db *gorm.DB) *LeagueRepository {
	return &LeagueRepository{
		db: db,
	}
}

type ILeagueRepository interface {
	Count(string) (int64, error)
	CountByCode(string) (int64, error)
	CountByPrimaryId(int) (int64, error)
	CountByName(string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllByActive() ([]*entity.League, error)
	GetOnlyWorldByActive() ([]*entity.League, error)
	GetAllUSSD(int) ([]*entity.League, error)
	GetAllEuropeUSSD(int) ([]*entity.League, error)
	GetAllAfriqueUSSD(int) ([]*entity.League, error)
	GetAllWorldUSSD(int) ([]*entity.League, error)
	GetAllInternationalUSSD(int) ([]*entity.League, error)
	GetAllTopLeagueUSSD(page int) ([]*entity.League, error)
	Get(string) (*entity.League, error)
	GetByCode(string) (*entity.League, error)
	GetByPrimaryId(int) (*entity.League, error)
	GetByName(string) (*entity.League, error)
	Save(*entity.League) (*entity.League, error)
	Update(*entity.League) (*entity.League, error)
	UpdateByPrimaryId(*entity.League) (*entity.League, error)
	Delete(*entity.League) error
}

func (r *LeagueRepository) Count(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) CountByCode(code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("code = ? AND is_active = true", code).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) CountByName(name string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.League{}).Where("is_active = ?", true).Where("UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+name+"%", "%"+name+"%", "%"+name+"%").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *LeagueRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var leagues []*entity.League
	err := r.db.Where("is_active = ?", true).Where("UPPER(name) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateLeagues(leagues, p, r.db)).Find(&leagues).Error
	if err != nil {
		return nil, err
	}
	p.Rows = leagues
	return p, nil
}

func (r *LeagueRepository) GetAllByActive() ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where(&entity.League{IsActive: true}).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetOnlyWorldByActive() ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where(&entity.League{IsActive: true, Country: "World", Code: "UEFA"}).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = ?", true).Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllEuropeUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = ?", true).Where("country IN('England', 'Belgium', 'Portugal', 'France', 'Italy', 'Spain', 'Germany', 'World')").Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllAfriqueUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = ?", true).Where("country IN('Mali')").Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllWorldUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = ?", true).Where("country IN('World')").Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllTopLeagueUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = true AND primary_id IN (39, 94, 61, 135, 140, 78)").Order("sort ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) GetAllInternationalUSSD(page int) ([]*entity.League, error) {
	var c []*entity.League
	err := r.db.Where("is_active = ?", true).Where("primary_id IN(34, 30, 36, 5)").Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) Get(slug string) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) GetByCode(slug string) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("code = ? AND is_active = true", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) GetByPrimaryId(primaryId int) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("primary_id = ?", primaryId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) GetByName(name string) (*entity.League, error) {
	var c entity.League
	err := r.db.Where("is_active = ?", true).Where("UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?)", "%"+name+"%", "%"+name+"%", "%"+name+"%").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *LeagueRepository) Save(c *entity.League) (*entity.League, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) Update(c *entity.League) (*entity.League, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) UpdateByPrimaryId(c *entity.League) (*entity.League, error) {
	err := r.db.Where("primary_id = ?", c.PrimaryID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *LeagueRepository) Delete(c *entity.League) error {
	err := r.db.Delete(&c, c.PrimaryID).Error
	if err != nil {
		return err
	}
	return nil
}
