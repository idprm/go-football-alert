package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionFollowCompetitionRepository struct {
	db *gorm.DB
}

func NewSubscriptionFollowCompetitionRepository(db *gorm.DB) *SubscriptionFollowCompetitionRepository {
	return &SubscriptionFollowCompetitionRepository{
		db: db,
	}
}

type ISubscriptionFollowCompetitionRepository interface {
	Count(int, int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, int) (*entity.SubscriptionFollowCompetition, error)
	Save(*entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error)
	Update(*entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error)
	Delete(*entity.SubscriptionFollowCompetition) error
}

func (r *SubscriptionFollowCompetitionRepository) Count(subId, leagueId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowCompetition{}).Where(&entity.SubscriptionFollowCompetition{SubscriptionID: int64(subId), LeagueID: int64(leagueId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowCompetitionRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionFollowCompetition
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionFollowCompetitionRepository) Get(subId, leagueId int) (*entity.SubscriptionFollowCompetition, error) {
	var c entity.SubscriptionFollowCompetition
	err := r.db.Where(&entity.SubscriptionFollowCompetition{SubscriptionID: int64(subId), LeagueID: int64(leagueId)}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionFollowCompetitionRepository) Save(c *entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowCompetitionRepository) Update(c *entity.SubscriptionFollowCompetition) (*entity.SubscriptionFollowCompetition, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowCompetitionRepository) Delete(c *entity.SubscriptionFollowCompetition) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}
