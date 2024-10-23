package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionFollowLeagueRepository struct {
	db *gorm.DB
}

func NewSubscriptionFollowLeagueRepository(db *gorm.DB) *SubscriptionFollowLeagueRepository {
	return &SubscriptionFollowLeagueRepository{
		db: db,
	}
}

type ISubscriptionFollowLeagueRepository interface {
	CountBySub(int64) (int64, error)
	CountByLeague(int64) (int64, error)
	CountByLimit(int64) (int64, error)
	CountByUpdated(subId int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetBySub(int64) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Disable(*entity.SubscriptionFollowLeague) error
	Delete(*entity.SubscriptionFollowLeague) error
	GetAllSubByLeague(int64) (*[]entity.SubscriptionFollowLeague, error)
}

func (r *SubscriptionFollowLeagueRepository) CountBySub(subId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowLeague{}).Where(&entity.SubscriptionFollowLeague{SubscriptionID: subId, IsActive: true}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowLeagueRepository) CountByLimit(subId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowLeague{}).Where(
		&entity.SubscriptionFollowLeague{SubscriptionID: subId, IsActive: true}).
		Where("limit_per_day >= sent").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowLeagueRepository) CountByUpdated(subId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowLeague{}).Where(
		&entity.SubscriptionFollowLeague{SubscriptionID: subId, IsActive: true}).
		Where("DATE(updated_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowLeagueRepository) CountByLeague(leagueId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowLeague{}).Where(&entity.SubscriptionFollowLeague{LeagueID: leagueId}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowLeagueRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionFollowLeague
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionFollowLeagueRepository) GetBySub(subId int64) (*entity.SubscriptionFollowLeague, error) {
	var c entity.SubscriptionFollowLeague
	err := r.db.Where(&entity.SubscriptionFollowLeague{SubscriptionID: subId, IsActive: true}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionFollowLeagueRepository) Save(c *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowLeagueRepository) Update(c *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	err := r.db.Where("subscription_id = ?", c.SubscriptionID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowLeagueRepository) Disable(c *entity.SubscriptionFollowLeague) error {
	err := r.db.Model(c).Where("subscription_id = ?", c.SubscriptionID).Update("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionFollowLeagueRepository) Delete(c *entity.SubscriptionFollowLeague) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionFollowLeagueRepository) GetAllSubByLeague(leagueId int64) (*[]entity.SubscriptionFollowLeague, error) {
	var sub []entity.SubscriptionFollowLeague
	err := r.db.Where(&entity.SubscriptionFollowLeague{LeagueID: leagueId, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}
