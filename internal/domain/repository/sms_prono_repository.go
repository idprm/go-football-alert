package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SMSPronoRespository struct {
	db *gorm.DB
}

func NewSMSPronoRespository(db *gorm.DB) *SMSPronoRespository {
	return &SMSPronoRespository{
		db: db,
	}
}

type ISMSPronoRespository interface {
	Count(int64, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Save(*entity.SMSProno) (*entity.SMSProno, error)
}

func (r *SMSPronoRespository) Count(subscriptionId, pronoId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SMSProno{}).Where(&entity.SMSProno{SubscriptionID: subscriptionId, PronosticID: pronoId}).Where("DATE(created_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SMSPronoRespository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var pronos []*entity.SMSProno
	// Where("Subscription.msisdn LIKE UPPER(?) OR News.title LIKE UPPER(?)", p.GetSearch(), p.GetSearch())
	err := r.db.Scopes(Paginate(pronos, p, r.db)).Joins("Subscription", r.db.Where("msisdn LIKE UPPER(?)", p.GetSearch())).Joins("Pronostic", r.db.Where("UPPER(title) LIKE UPPER(?)", p.GetSearch())).Find(&pronos).Error
	if err != nil {
		return nil, err
	}
	p.Rows = pronos
	return p, nil
}

func (r *SMSPronoRespository) Save(c *entity.SMSProno) (*entity.SMSProno, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
