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
	Count(string, string) (int64, error)
	CountNewsLeague(int64) (int64, error)
	CountNewsTeam(int64) (int64, error)
	CountById(int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllUSSD(int) ([]*entity.News, error)
	GetByTeamUSSD(int) (*entity.News, error)
	GetById(int64) (*entity.News, error)
	GetBySlug(string) (*entity.News, error)
	Get(string, string) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
	GetAllNewsLeague(int64) ([]*entity.NewsLeagues, error)
	GetAllNewsTeam(int64) ([]*entity.NewsTeams, error)
	SaveNewsLeague(*entity.NewsLeagues) (*entity.NewsLeagues, error)
	UpdateNewsLeague(*entity.NewsLeagues) (*entity.NewsLeagues, error)
	SaveNewsTeam(*entity.NewsTeams) (*entity.NewsTeams, error)
	UpdateNewsTeam(*entity.NewsTeams) (*entity.NewsTeams, error)
}

func (r *NewsRepository) Count(slug, pubAt string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.News{}).Where("slug = ?", slug).Where("DATE(publish_at) = DATE(?)", pubAt).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *NewsRepository) CountNewsLeague(leagueId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.NewsLeagues{}).Where("league_id = ?", leagueId).Where("DATE(created_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *NewsRepository) CountNewsTeam(teamId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.NewsTeams{}).Where("team_id = ?", teamId).Where("DATE(created_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *NewsRepository) CountById(id int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.News{}).Where("id = ?", id).Count(&count).Error
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

func (r *NewsRepository) GetAllUSSD(page int) ([]*entity.News, error) {
	var c []*entity.News
	err := r.db.Where("DATE(publish_at) BETWEEN DATE(NOW() - INTERVAL 1 DAY) AND DATE(NOW())").Order("DATE(publish_at) DESC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) GetByTeamUSSD(teamId int) (*entity.News, error) {
	var c entity.News
	err := r.db.Where("DATE(publish_at) <= DATE(NOW())").Where("team_id = ?", teamId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *NewsRepository) GetById(id int64) (*entity.News, error) {
	var c entity.News
	err := r.db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *NewsRepository) GetBySlug(slug string) (*entity.News, error) {
	var c entity.News
	err := r.db.Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *NewsRepository) Get(slug, pubAt string) (*entity.News, error) {
	var c entity.News
	err := r.db.Where("slug = ?", slug).Where("DATE(publish_at) = DATE(NOW())", pubAt).Take(&c).Error
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

func (r *NewsRepository) GetAllNewsLeague(leagueId int64) ([]*entity.NewsLeagues, error) {
	var c []*entity.NewsLeagues
	err := r.db.Where("league_id = ?", leagueId).Where("DATE(created_at) BETWEEN DATE(NOW() - INTERVAL 1 DAY) AND DATE(NOW())").Preload("News").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) GetAllNewsTeam(teamId int64) ([]*entity.NewsTeams, error) {
	var c []*entity.NewsTeams
	err := r.db.Where("team_id = ?", teamId).Where("DATE(created_at) BETWEEN DATE(NOW() - INTERVAL 1 DAY) AND DATE(NOW())").Preload("News").Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) SaveNewsLeague(c *entity.NewsLeagues) (*entity.NewsLeagues, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) UpdateNewsLeague(c *entity.NewsLeagues) (*entity.NewsLeagues, error) {
	err := r.db.Where("news_id = ?", c.NewsID).Where("DATE(created_at) = DATE(NOW())").Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) SaveNewsTeam(c *entity.NewsTeams) (*entity.NewsTeams, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *NewsRepository) UpdateNewsTeam(c *entity.NewsTeams) (*entity.NewsTeams, error) {
	err := r.db.Where("news_id = ?", c.NewsID).Where("DATE(created_at) = DATE(NOW())").Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
