package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type SubscriptionFollowTeamRepository struct {
	db *gorm.DB
}

func NewSubscriptionFollowTeamRepository(db *gorm.DB) *SubscriptionFollowTeamRepository {
	return &SubscriptionFollowTeamRepository{
		db: db,
	}
}

type ISubscriptionFollowTeamRepository interface {
	Count(int64, int64) (int64, error)
	CountByTeam(int64) (int64, error)
	CountByLimit(int64, int64) (int64, error)
	CountByUpdated(int64, int64) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.SubscriptionFollowTeam, error)
	Save(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Update(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Disable(*entity.SubscriptionFollowTeam) error
	Delete(*entity.SubscriptionFollowTeam) error
	GetAllSubByTeam(int64) (*[]entity.SubscriptionFollowTeam, error)
	Renewal() (*[]entity.SubscriptionFollowTeam, error)
	Retry() (*[]entity.SubscriptionFollowTeam, error)
}

func (r *SubscriptionFollowTeamRepository) Count(subId, teamId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowTeam{}).Where(
		&entity.SubscriptionFollowTeam{SubscriptionID: subId, TeamID: teamId, IsActive: true}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowTeamRepository) CountByTeam(teamId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowTeam{}).Where(&entity.SubscriptionFollowTeam{TeamID: teamId}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowTeamRepository) CountByLimit(subId, teamId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowTeam{}).Where(
		&entity.SubscriptionFollowTeam{SubscriptionID: subId, TeamID: teamId, IsActive: true}).
		Where("sent <= limit_per_day AND DATE(updated_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowTeamRepository) CountByUpdated(subId, teamId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.SubscriptionFollowTeam{}).Where(
		&entity.SubscriptionFollowTeam{SubscriptionID: subId, TeamID: teamId, IsActive: true}).
		Where("DATE(updated_at) = DATE(NOW())").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionFollowTeamRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var creditgoals []*entity.SubscriptionFollowTeam
	err := r.db.Scopes(Paginate(creditgoals, pagination, r.db)).Find(&creditgoals).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = creditgoals
	return pagination, nil
}

func (r *SubscriptionFollowTeamRepository) Get(subId, teamId int64) (*entity.SubscriptionFollowTeam, error) {
	var c entity.SubscriptionFollowTeam
	err := r.db.Where(&entity.SubscriptionFollowTeam{SubscriptionID: subId, TeamID: teamId, IsActive: true}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SubscriptionFollowTeamRepository) Save(c *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowTeamRepository) Update(c *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	err := r.db.Where("subscription_id = ? AND team_id = ?", c.SubscriptionID, c.TeamID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *SubscriptionFollowTeamRepository) Disable(c *entity.SubscriptionFollowTeam) error {
	err := r.db.Model(c).Where("subscription_id = ? AND team_id = ?", c.SubscriptionID, c.TeamID).Update("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionFollowTeamRepository) Delete(c *entity.SubscriptionFollowTeam) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionFollowTeamRepository) GetAllSubByTeam(teamId int64) (*[]entity.SubscriptionFollowTeam, error) {
	var sub []entity.SubscriptionFollowTeam
	err := r.db.Where(&entity.SubscriptionFollowTeam{TeamID: teamId, IsActive: true}).Find(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

// SELECT (UNIX_TIMESTAMP("2017-06-10 18:30:10")-UNIX_TIMESTAMP("2017-06-10 18:40:10"))/3600 hour_diff
func (r *SubscriptionFollowTeamRepository) Renewal() (*[]entity.SubscriptionFollowTeam, error) {
	var sub []entity.SubscriptionFollowTeam
	err := r.db.Where("is_active = true AND (UNIX_TIMESTAMP(NOW()) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

// SELECT (UNIX_TIMESTAMP("2017-06-10 18:30:10" + INTERVAL 1 DAY)-UNIX_TIMESTAMP("2017-06-10 18:40:10"))/3600 hour_diff (tommorow)
func (r *SubscriptionFollowTeamRepository) Retry() (*[]entity.SubscriptionFollowTeam, error) {
	var sub []entity.SubscriptionFollowTeam
	err := r.db.Where("is_active = true AND is_retry = true AND (UNIX_TIMESTAMP(NOW() + INTERVAL 1 DAY) - UNIX_TIMESTAMP(renewal_at)) / 3600 > 0").Order("DATE(created_at) DESC").Find(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}
