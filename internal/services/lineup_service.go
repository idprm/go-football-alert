package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type LineupService struct {
	lineupRepo repository.ILineupRepository
}

func NewLineupService(lineupRepo repository.ILineupRepository) *LineupService {
	return &LineupService{
		lineupRepo: lineupRepo,
	}
}

type ILineupService interface {
	IsLineup(string) bool
	IsLineupByPrimaryId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(string) (*entity.Lineup, error)
	GetByPrimaryId(int) (*entity.Lineup, error)
	Save(*entity.Lineup) (*entity.Lineup, error)
	Update(*entity.Lineup) (*entity.Lineup, error)
	UpdateByPrimaryId(*entity.Lineup) (*entity.Lineup, error)
	Delete(*entity.Lineup) error
}

func (s *LineupService) IsLineup(fixtureId int) bool {
	count, _ := s.lineupRepo.Count(fixtureId)
	return count > 0
}

func (s *LineupService) IsLineupByPrimaryId(primaryId int) bool {
	count, _ := s.lineupRepo.CountByPrimaryId(primaryId)
	return count > 0
}

func (s *LineupService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.lineupRepo.GetAllPaginate(pagination)
}

func (s *LineupService) Get(fixtureId int) (*entity.Lineup, error) {
	return s.lineupRepo.Get(fixtureId)
}

func (s *LineupService) GetByPrimaryId(primaryId int) (*entity.Lineup, error) {
	return s.lineupRepo.GetByPrimaryId(primaryId)
}

func (s *LineupService) Save(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.Save(a)
}

func (s *LineupService) Update(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.Update(a)
}

func (s *LineupService) UpdateByPrimaryId(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.UpdateByPrimaryId(a)
}

func (s *LineupService) Delete(a *entity.Lineup) error {
	return s.lineupRepo.Delete(a)
}
