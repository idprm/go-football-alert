package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

type IContentRepository interface {
	Count(string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Content, error)
	Save(*entity.Content) (*entity.Content, error)
	Update(*entity.Content) (*entity.Content, error)
	Delete(*entity.Content) error
}

func (r *ContentRepository) Count(name string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Content{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ContentRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var contents []*entity.Content
	err := r.db.Where("UPPER(name) LIKE UPPER(?)", "%"+p.GetSearch()+"%").Scopes(Paginate(contents, p, r.db)).Find(&contents).Error
	if err != nil {
		return nil, err
	}
	p.Rows = contents
	return p, nil
}

func (r *ContentRepository) Get(name string) (*entity.Content, error) {
	var content entity.Content
	err := r.db.Where("name = ?", name).Take(&content).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *ContentRepository) Save(c *entity.Content) (*entity.Content, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ContentRepository) Update(c *entity.Content) (*entity.Content, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ContentRepository) Delete(c *entity.Content) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
