package repository

import (
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
	Count(int, string, string) (int64, error)
	CountActive(int, string, string) (int64, error)
	CountActiveByCategory(string, string, string) (int64, error)
	CountActiveByNonSMSAlerte(string, string) (int64, error)
	CountActiveBySubId(int64) (int64, error)
	CountActiveBySMSAlerteMsisdn(string) (int64, error)
	CountActiveAllByMsisdn(string) (int64, error)
	CountAfter24Hour(int, string, string) (int64, error)
	CountRenewal(int, string, string) (int64, error)
	CountRetry(int, string, string) (int64, error)
	CountRetryUnderpayment(int, string, string) (int64, error)
	CountTotalActiveSub() (int64, error)
	CountPopulateRenewal() (int64, error)
	CountPopulateRetry() (int64, error)
	CountPopulateRetryUnderpayment() (int64, error)
	CountPopulateReminder48HBeforeCharging() (int64, error)
	CountPopulateReminderAfterTrialEnds() (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetByCategory(string, string, string) (*entity.Subscription, error)
	GetActiveByCategory(string, string, string) (*entity.Subscription, error)
	GetByNonSMSAlerte(string, string) (*entity.Subscription, error)
	GetActiveByNonSMSAlerte(string, string) (*entity.Subscription, error)
	GetActiveBySMSAlerteMsisdn(string) (*entity.Subscription, error)
	GetBySubId(int64) (*entity.Subscription, error)
	GetActiveAllByMsisdnUSSD(string, int) (*[]entity.Subscription, error)
	Get(int, string, string) (*entity.Subscription, error)
	GetLongRetry(int, string, string) (int, error)
	Save(*entity.Subscription) (*entity.Subscription, error)
	Update(*entity.Subscription) (*entity.Subscription, error)
	Delete(*entity.Subscription) error
	UpdateNotActive(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFree(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotRetry(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotUnderpayment(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFollowTeam(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFollowLeague(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotPredictWin(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotCreditGoal(*entity.Subscription) (*entity.Subscription, error)
	CreditGoal() (*[]entity.Subscription, error)
	PredictWin() (*[]entity.Subscription, error)
	Follow() (*[]entity.Subscription, error)
	Prono() (*[]entity.Subscription, error)
	Renewal() (*[]entity.Subscription, error)
	Retry() (*[]entity.Subscription, error)
	RetryUnderpayment() (*[]entity.Subscription, error)
	Reminder48HBeforeCharging() (*[]entity.Subscription, error)
	ReminderAfterTrialEnds() (*[]entity.Subscription, error)
	GetAllSubBySMSAlerte() (*[]entity.Subscription, error)
}

func (r *SubscriptionRepository) Count(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActive(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("is_active = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveByCategory(category, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("category = ? AND msisdn = ? AND code = ?", category, msisdn, code).Where("is_active = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveByNonSMSAlerte(category, msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("category = ? AND msisdn = ?", category, msisdn).Where("is_active = ?", true).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveBySMSAlerteMsisdn(msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("category IN('SMSALERTE_COMPETITION', 'SMSALERTE_EQUIPE') AND msisdn = ?", msisdn).Where("is_retry = false AND is_active = true").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveBySubId(subId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("id = ?", subId).Where("is_retry = false AND is_active = true").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActiveAllByMsisdn(msisdn string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("msisdn = ? AND is_active = true", msisdn).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountAfter24Hour(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ? AND HOUR(TIMEDIFF(NOW(), created_at)) > 24", serviceId, msisdn, code).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountRenewal(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("is_active = true AND (UNIX_TIMESTAMP(NOW()) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountRetry(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("is_active = true AND is_retry = true AND (UNIX_TIMESTAMP(NOW() + INTERVAL 1 DAY) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountRetryUnderpayment(serviceId int, msisdn, code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Where("is_active = true AND is_retry = false AND is_underpayment = true AND total_underpayment > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountTotalActiveSub() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountPopulateRenewal() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true AND (UNIX_TIMESTAMP(NOW()) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountPopulateRetry() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true AND is_retry = true AND (UNIX_TIMESTAMP(NOW() + INTERVAL 1 DAY) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountPopulateRetryUnderpayment() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true AND is_retry = false AND is_underpayment = true AND total_underpayment > 0").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountPopulateReminder48HBeforeCharging() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true AND HOUR(TIMEDIFF(NOW(), renewal_at)) = 48 AND is_free = false AND service_id IN(2, 3, 5, 6, 8, 9, 11, 12, 14, 15, 17, 18, 20, 21)").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountPopulateReminderAfterTrialEnds() (int64, error) {
	var count int64
	err := r.db.Model(&entity.Subscription{}).Where("is_active = true AND HOUR(TIMEDIFF(NOW(), free_at)) BETWEEN 23 AND 24 AND is_free = true AND service_id IN(2, 3, 5, 6, 8, 9, 11, 12, 14, 15, 17, 18, 20, 21)").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var subscriptions []*entity.Subscription
	err := r.db.Where("UPPER(msisdn) LIKE UPPER(?) OR UPPER(latest_keyword) LIKE UPPER(?)", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateSubscriptions(subscriptions, p, r.db)).Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	p.Rows = subscriptions
	return p, nil
}

func (r *SubscriptionRepository) GetByCategory(category, msisdn, code string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("category = ? AND msisdn = ? AND code = ?", category, msisdn, code).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetActiveByCategory(category, msisdn, code string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("category = ? AND msisdn = ? AND code = ? AND is_active = true", category, msisdn, code).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetByNonSMSAlerte(category, msisdn string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("category = ? AND msisdn = ?", category, msisdn).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetActiveByNonSMSAlerte(category, msisdn string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("category = ? AND msisdn = ? AND is_active = true", category, msisdn).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetActiveBySMSAlerteMsisdn(msisdn string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("category IN('SMSALERTE_COMPETITION', 'SMSALERTE_EQUIPE') AND msisdn = ? AND is_active = true", msisdn).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetBySubId(subId int64) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("id = ?", subId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetActiveAllByMsisdnUSSD(msisdn string, page int) (*[]entity.Subscription, error) {
	var c []entity.Subscription
	err := r.db.Where("msisdn = ? AND is_active = true", msisdn).Preload("Service").Order("id ASC").Offset((page - 1) * 5).Limit(5).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) Get(serviceId int, msisdn, code string) (*entity.Subscription, error) {
	var c entity.Subscription
	err := r.db.Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionRepository) GetLongRetry(serviceId int, msisdn, code string) (int, error) {
	var c entity.Subscription
	err := r.db.Select("ROUND((UNIX_TIMESTAMP(NOW() + INTERVAL 1 DAY) - UNIX_TIMESTAMP(renewal_at)) / 3600) as long_retry").Where("service_id = ? AND msisdn = ? AND code = ?", serviceId, msisdn, code).Take(&c).Error
	if err != nil {
		return 0, err
	}
	return c.LongRetry, nil
}

func (r *SubscriptionRepository) Save(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) Update(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Updates(&c).Error
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

func (r *SubscriptionRepository) UpdateNotActive(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_active", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotRetry(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_retry", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotUnderpayment(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_underpayment", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotFree(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_free", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotFollowTeam(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_follow_team", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotFollowLeague(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_follow_competition", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotPredictWin(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_predict_win", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) UpdateNotCreditGoal(c *entity.Subscription) (*entity.Subscription, error) {
	err := r.db.Model(c).Where("service_id = ? AND msisdn = ? AND code = ?", c.ServiceID, c.Msisdn, c.Code).Update("is_credit_goal", false).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionRepository) Follow() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND (is_follow_team = true OR is_follow_competition = true)").Find(&sub).Error
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

func (r *SubscriptionRepository) PredictWin() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where(&entity.Subscription{IsPredictWin: true, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriptionRepository) Prono() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("category = 'PRONOSTIC' AND is_retry = false AND is_active = true").Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

// SELECT (UNIX_TIMESTAMP("2017-06-10 18:30:10")-UNIX_TIMESTAMP("2017-06-10 18:40:10"))/3600 hour_diff
func (r *SubscriptionRepository) Renewal() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND (UNIX_TIMESTAMP(NOW()) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// SELECT (UNIX_TIMESTAMP("2017-06-10 18:30:10" + INTERVAL 1 DAY)-UNIX_TIMESTAMP("2017-06-10 18:40:10"))/3600 hour_diff (tommorow)
func (r *SubscriptionRepository) Retry() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND is_retry = true AND (UNIX_TIMESTAMP(NOW() + INTERVAL 1 DAY) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) RetryUnderpayment() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND is_retry = false AND is_underpayment = true AND total_underpayment > 0").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Reminder48HBeforeCharging() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND HOUR(TIMEDIFF(NOW(), renewal_at)) = 48 AND is_free = false AND service_id IN(2, 3, 5, 6, 8, 9, 11, 12, 14, 15, 17, 18, 20, 21)").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) ReminderAfterTrialEnds() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Where("is_active = true AND HOUR(TIMEDIFF(NOW(), free_at)) BETWEEN 23 AND 24 AND is_free = true AND service_id IN(2, 3, 5, 6, 8, 9, 11, 12, 14, 15, 17, 18, 20, 21)").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) GetAllSubBySMSAlerte() (*[]entity.Subscription, error) {
	var sub []entity.Subscription
	err := r.db.Select("category, service_id, msisdn, code, channel, latest_subject").Where("category IN('SMSALERTE_COMPETITION','SMSALERTE_EQUIPE') AND is_active = true").Group("msisdn").Find(&sub).Error
	if err != nil {
		return nil, err
	}

	var subs []entity.Subscription
	if len(sub) > 0 {
		for _, a := range sub {
			sub := entity.Subscription{
				Category:      a.Category,
				ServiceID:     a.ServiceID,
				Msisdn:        a.Msisdn,
				Code:          a.Code,
				Channel:       a.Channel,
				LatestSubject: a.LatestSubject,
			}
			subs = append(subs, sub)
		}
	}

	return &subs, nil
}
