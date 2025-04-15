package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type IUserRepository interface {
	CountByEmail(string) (int64, error)
	Count(string, string) (int64, error)
	Get(string, string) (*entity.User, error)
	Save(*entity.User) error
	Update(*entity.User) error
	Delete(*entity.User) error
}

func (r *UserRepository) CountByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *UserRepository) Count(email, pass string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("email = ? AND pass = ?", email, pass).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *UserRepository) Get(email, pass string) (*entity.User, error) {
	var c entity.User
	err := r.db.Where("email = ? AND pass = ?", email, pass).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *UserRepository) Save(c *entity.User) error {
	err := r.db.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(c *entity.User) error {
	err := r.db.Where("email = ?", c.Email).Updates(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(c *entity.User) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
