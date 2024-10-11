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
	Count(int64, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Delete(*entity.SubscriptionFollowLeague) error
}

func (r *SubscriptionFollowLeagueRepository) Count(subId, leagueId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowLeague{}).Where(&entity.SubscriptionFollowLeague{SubscriptionID: int64(subId), LeagueID: int64(leagueId)}).Count(&count).Error
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

func (r *SubscriptionFollowLeagueRepository) Get(subId, leagueId int64) (*entity.SubscriptionFollowLeague, error) {
	var c entity.SubscriptionFollowLeague
	err := r.db.Where(&entity.SubscriptionFollowLeague{SubscriptionID: subId, LeagueID: leagueId}).Take(&c).Error
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
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowLeagueRepository) Delete(c *entity.SubscriptionFollowLeague) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
