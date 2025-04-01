package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SMSActuRespository struct {
	db *gorm.DB
}

func NewSMSActuRespository(db *gorm.DB) *SMSActuRespository {
	return &SMSActuRespository{
		db: db,
	}
}

type ISMSActuRespository interface {
	Count(string, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.SMSActu) (*entity.SMSActu, error)
}

func (r *SMSActuRespository) Count(msisdn string, newsId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SMSActu{}).Where(&entity.SMSActu{Msisdn: msisdn, NewsID: newsId}).Where("DATE(created_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SMSActuRespository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var alerts []*entity.SMSAlerte
	// Where("Subscription.msisdn LIKE UPPER(?) OR News.title LIKE UPPER(?)", p.GetSearch(), p.GetSearch())
	err := r.db.Scopes(Paginate(alerts, p, r.db)).Where("msisdn LIKE UPPER(?)", p.GetSearch()).Joins("News", r.db.Where("UPPER(title) LIKE UPPER(?)", p.GetSearch())).Find(&alerts).Error
	if err != nil {
		return nil, err
	}
	p.Rows = alerts
	return p, nil
}

func (r *SMSActuRespository) Save(c *entity.SMSActu) (*entity.SMSActu, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
