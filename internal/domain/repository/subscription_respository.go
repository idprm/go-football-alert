package repository

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

type ISubscriptionRepository interface {
	Count(int, string) (int64, error)
	CountActive(int, string) (int64, error)
	CountActiveByCategory(string, string) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string) (*entity.Subscription, error)
	Save(*entity.Subscription) (*entity.Subscription, error)
	Update(*entity.Subscription) (*entity.Subscription, error)
	Delete(*entity.Subscription) error
	IsNotActive(*entity.Subscription) (*entity.Subscription, error)
	IsNotRetry(*entity.Subscription) (*entity.Subscription, error)
	IsNotFollowTeam(*entity.Subscription) (*entity.Subscription, error)
	IsNotFollowCompetition(*entity.Subscription) (*entity.Subscription, error)
	IsNotPrediction(*entity.Subscription) (*entity.Subscription, error)
	IsNotCreditGoal(*entity.Subscription) (*entity.Subscription, error)
	CreditGoal() (*[]entity.Subscription, error)
	Prediction() (*[]entity.Subscription, error)
	FollowTeam() (*[]entity.Subscription, error)
	FollowCompetition() (*[]entity.Subscription, error)
	Renewal() (*[]entity.Subscription, error)
	Retry() (*[]entity.Subscription, error)
}

func (r *SubscriptionRepository) Count(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActive(serviceId int, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Where("is_active = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveByCategory(category string, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("category = ?", category).Where("msisdn = ?", msisdn).Where("is_active = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var subscriptions []*entity.Subscription
	err := r.db.Scopes(Paginate(subscriptions, pagination, r.db)).Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = subscriptions
	return pagination, nil
}

func (r *SubscriptionRepository) Get(serviceId int, msisdn string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("service_id = ?", serviceId).Where("msisdn = ?", msisdn).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) Save(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) Update(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) Delete(c *entity.Subscription) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepository) IsNotActive(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_active": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) IsNotRetry(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_retry": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) IsNotFollowTeam(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_follow_team": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) IsNotFollowCompetition(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_follow_competition": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) IsNotPrediction(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_prediction": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) IsNotCreditGoal(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ?", c.ServiceID).Where("msisdn = ?", c.Msisdn).Updates(map[string]interface{}{"updated_at": time.Now(), "is_credit_goal": false}).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) FollowCompetition() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsFollowCompetition: true, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriptionRepository) FollowTeam() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsFollowTeam: true, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriptionRepository) CreditGoal() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsCreditGoal: true, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriptionRepository) Prediction() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsPrediction: true, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

// ORDER BY success DESC, DATE(created_at) DES
func (r *SubscriptionRepository) Renewal() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsActive: true}).Where("DATE(renewal_at) <= DATE(NOW())").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Retry() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsRetry: true, IsActive: true}).Where("DATE(renewal_at) = DATE(NOW() + INTERVAL 1 DAY)").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}
