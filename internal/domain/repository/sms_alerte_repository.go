package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SMSAlerteRespository struct {
	db *gorm.DB
}

func NewSMSAlerteRespository(db *gorm.DB) *SMSAlerteRespository {
	return &SMSAlerteRespository{
		db: db,
	}
}

type ISMSAlerteRespository interface {
	Count(int64, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.SMSAlerte) (*entity.SMSAlerte, error)
}

func (r *SMSAlerteRespository) Count(subscriptionId, newsId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SMSAlerte{}).Where(&entity.SMSAlerte{SubscriptionID: subscriptionId, NewsID: newsId}).Where("DATE(created_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SMSAlerteRespository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var alerts []*entity.SMSAlerte
	err := r.db.Scopes(Paginate(alerts, pagination, r.db)).Preload("Subscription").Preload("News").Find(&alerts).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = alerts
	return pagination, nil
}

func (r *SMSAlerteRespository) Save(c *entity.SMSAlerte) (*entity.SMSAlerte, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
