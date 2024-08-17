package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type AwayRepository struct {
	db *gorm.DB
}

func NewAwayRepository(db *gorm.DB) *AwayRepository {
	return &AwayRepository{
		db: db,
	}
}

type IAwayRepository interface {
	CountByTeamId(int) (int64, error)
	CountByPrimaryId(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetByTeamId(int) (*entity.Away, error)
	Save(*entity.Away) (*entity.Away, error)
	Update(*entity.Away) (*entity.Away, error)
	Delete(*entity.Away) error
}

func (r *AwayRepository) CountByTeamId(teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Away{}).Where("team_id = ?", teamId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *AwayRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Away{}).Where("primary_id = ?", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *AwayRepository) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	var aways []*entity.Away
	err := r.db.Scopes(Paginate(aways, pagination, r.db)).Preload("Team").Find(&aways).Error
	if err != nil {
		return nil, err
	}
	pagination.Rows = aways
	return pagination, nil
}

func (r *AwayRepository) GetByTeamId(teamId int) (*entity.Away, error) {
	var away entity.Away
	err := r.db.Where("team_id = ?", teamId).Preload("Team").Take(&away).Error
	if err != nil {
		return nil, err
	}
	return &away, nil
}

func (r *AwayRepository) Save(away *entity.Away) (*entity.Away, error) {
	err := r.db.Create(&away).Error
	if err != nil {
		return nil, err
	}
	return away, nil
}

func (r *AwayRepository) Update(away *entity.Away) (*entity.Away, error) {
	err := r.db.Where("id = ?", away.ID).Updates(&away).Error
	if err != nil {
		return nil, err
	}
	return away, nil
}

func (r *AwayRepository) Delete(p *entity.Away) error {
	err := r.db.Delete(&p, p.ID).Error
	if err != nil {
		return err
	}
	return nil
}
