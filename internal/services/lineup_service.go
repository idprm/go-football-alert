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
	IsLineup(int) bool
	IsLineupByFixtureId(int) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllUSSD() ([]*entity.Lineup, error)
	Get(int) (*entity.Lineup, error)
	GetByFixtureId(int) (*entity.Lineup, error)
	Save(*entity.Lineup) (*entity.Lineup, error)
	Update(*entity.Lineup) (*entity.Lineup, error)
	UpdateByFixtureId(*entity.Lineup) (*entity.Lineup, error)
	Delete(*entity.Lineup) error
}

func (s *LineupService) IsLineup(fixtureId int) bool {
	count, _ := s.lineupRepo.Count(fixtureId)
	return count > 0
}

func (s *LineupService) IsLineupByFixtureId(primaryId int) bool {
	count, _ := s.lineupRepo.CountByFixtureId(primaryId)
	return count > 0
}

func (s *LineupService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.lineupRepo.GetAllPaginate(pagination)
}

func (s *LineupService) GetAllUSSD() ([]*entity.Lineup, error) {
	return s.lineupRepo.GetAllUSSD()
}

func (s *LineupService) Get(fixtureId int) (*entity.Lineup, error) {
	return s.lineupRepo.Get(fixtureId)
}

func (s *LineupService) GetByFixtureId(fixtureId int) (*entity.Lineup, error) {
	return s.lineupRepo.GetByFixtureId(fixtureId)
}

func (s *LineupService) Save(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.Save(a)
}

func (s *LineupService) Update(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.Update(a)
}

func (s *LineupService) UpdateByFixtureId(a *entity.Lineup) (*entity.Lineup, error) {
	return s.lineupRepo.UpdateByFixtureId(a)
}

func (s *LineupService) Delete(a *entity.Lineup) error {
	return s.lineupRepo.Delete(a)
}
