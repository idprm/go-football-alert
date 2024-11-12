package repository

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

type ITeamRepository interface {
	Count(string) (int64, error)
	CountByCode(string) (int64, error)
	CountByPrimaryId(int) (int64, error)
	CountByName(string) (int64, error)
	CountByLeagueTeam(*entity.LeagueTeam) (int64, error)
	CountLeagueByTeam(int) (int64, error)
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllTeamUSSD(int, int) ([]*entity.LeagueTeam, error)
	Get(string) (*entity.Team, error)
	GetByCode(string) (*entity.Team, error)
	GetByPrimaryId(int) (*entity.Team, error)
	GetByName(string) (*entity.Team, error)
	Save(*entity.Team) (*entity.Team, error)
	Update(*entity.Team) (*entity.Team, error)
	UpdateByPrimaryId(*entity.Team) (*entity.Team, error)
	Delete(*entity.Team) error
	GetLeagueByTeam(int) (*entity.LeagueTeam, error)
	SaveLeagueTeam(*entity.LeagueTeam) (*entity.LeagueTeam, error)
}

func (r *TeamRepository) Count(slug string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Team{}).Where("slug = ?", slug).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) CountByCode(code string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Team{}).Where("code = ? AND is_active = true", code).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) CountByPrimaryId(primaryId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Team{}).Where("primary_id = ? AND is_active = true", primaryId).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) CountByName(name string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Team{}).Where("is_active = true AND (UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?))", "%"+name+"%", "%"+name+"%", "%"+name+"%").Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) CountByLeagueTeam(v *entity.LeagueTeam) (int64, error) {
	var count int64
	err := r.db.Model(&entity.LeagueTeam{}).Where(&entity.LeagueTeam{LeagueID: v.LeagueID, TeamID: v.TeamID}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) CountLeagueByTeam(teamId int) (int64, error) {
	var count int64
	err := r.db.Model(&entity.LeagueTeam{}).Where(&entity.LeagueTeam{TeamID: int64(teamId)}).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *TeamRepository) GetAllPaginate(p *entity.Pagination) (*entity.Pagination, error) {
	var teams []*entity.Team
	err := r.db.Where("is_active = true AND (UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?))", "%"+p.GetSearch()+"%", "%"+p.GetSearch()+"%").Scopes(PaginateIsActive(teams, p, r.db)).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	p.Rows = teams
	return p, nil
}

func (r *TeamRepository) GetAllTeamUSSD(leagueId, page int) ([]*entity.LeagueTeam, error) {
	var c []*entity.LeagueTeam
	err := r.db.Where("league_id = ?", leagueId).Preload("League").Preload("Team").Order("id ASC").Offset((page - 1) * 7).Limit(7).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TeamRepository) Get(slug string) (*entity.Team, error) {
	var c entity.Team
	err := r.db.Where("slug = ?", slug).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TeamRepository) GetByCode(code string) (*entity.Team, error) {
	var c entity.Team
	err := r.db.Where("code = ? AND is_active = true", code).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TeamRepository) GetByPrimaryId(primaryId int) (*entity.Team, error) {
	var c entity.Team
	err := r.db.Where("primary_id = ?", primaryId).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TeamRepository) GetByName(name string) (*entity.Team, error) {
	var c entity.Team
	err := r.db.Where("is_active = true AND (UPPER(name) LIKE UPPER(?) OR UPPER(code) LIKE UPPER(?) OR UPPER(keyword) LIKE UPPER(?))", "%"+name+"%", "%"+name+"%", "%"+name+"%").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TeamRepository) Save(c *entity.Team) (*entity.Team, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TeamRepository) Update(c *entity.Team) (*entity.Team, error) {
	err := r.db.Where("id = ?", c.ID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TeamRepository) UpdateByPrimaryId(c *entity.Team) (*entity.Team, error) {
	err := r.db.Where("primary_id = ?", c.PrimaryID).Updates(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *TeamRepository) Delete(c *entity.Team) error {
	err := r.db.Delete(&c, c.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TeamRepository) GetLeagueByTeam(teamId int) (*entity.LeagueTeam, error) {
	var c entity.LeagueTeam
	err := r.db.Where("team_id = ?", teamId).Preload("Team").Preload("League").Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *TeamRepository) SaveLeagueTeam(c *entity.LeagueTeam) (*entity.LeagueTeam, error) {
	err := r.db.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
