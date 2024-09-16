package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UssdRepository struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewUssdRepository(
	db *gorm.DB,
	rds *redis.Client,
) *UssdRepository {
	return &UssdRepository{
		db:  db,
		rds: rds,
	}
}

type IUssdRepository interface {
	GetAll() ([]*entity.Ussd, error)
	Save(*entity.Ussd) (*entity.Ussd, error)
	Update(*entity.Ussd) (*entity.Ussd, error)
	Delete(*entity.Ussd) error
	Set(*entity.Ussd) error
	Get(string) (*entity.Ussd, error)
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

func (r *UssdRepository) Set(e *entity.Ussd) error {
	jsonData, _ := json.Marshal(e)
	err := r.rds.Set(context.TODO(), "ussd_"+e.GetMsisdn(), string(jsonData), 48*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UssdRepository) Get(msisdn string) (*entity.Ussd, error) {
	val, err := r.rds.Get(context.TODO(), "ussd_"+msisdn).Result()
	if err != nil {
		return nil, err
	}
	var e *entity.Ussd
	json.Unmarshal([]byte(val), &e)
	return e, nil
}
