package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type UssdRepository struct {
	db *gorm.DB
}

func NewUssdRepository(db *gorm.DB) *UssdRepository {
	return &UssdRepository{
		db: db,
	}
}

type IUssdRepository interface {
	GetAll() ([]*entity.Ussd, error)
	Save(*entity.Ussd) (*entity.Ussd, error)
	Update(*entity.Ussd) (*entity.Ussd, error)
	Delete(*entity.Ussd) error
}

func (r *UssdRepository) GetAll() ([]*entity.Ussd, error) {
	var ussds []*entity.Ussd
	err := r.db.Where("is_active = ?", true).Order("parent, child ASC").Find(&ussds).Error
	if err != nil {
		return nil, err
	}
	return ussds, nil
}

func (r *UssdRepository) Save(e *entity.Ussd) (*entity.Ussd, error) {
	err := r.db.Create(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *UssdRepository) Update(e *entity.Ussd) (*entity.Ussd, error) {
	err := r.db.Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *UssdRepository) Delete(e *entity.Ussd) error {
	err := r.db.Delete(&e, e.ID).Error
	if err != nil {
		return err
	}
	return nil
}
