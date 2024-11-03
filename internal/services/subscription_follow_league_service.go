package services

import (
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionFollowLeagueService struct {
	subFollowLeagueRepo repository.ISubscriptionFollowLeagueRepository
}

func NewSubscriptionFollowLeagueService(
	subFollowLeagueRepo repository.ISubscriptionFollowLeagueRepository,
) *SubscriptionFollowLeagueService {
	return &SubscriptionFollowLeagueService{
		subFollowLeagueRepo: subFollowLeagueRepo,
	}
}

type ISubscriptionFollowLeagueService interface {
	IsSub(int64, int64) bool
	CountSub(int64, int64) int64
	IsLeague(int64) bool
	IsLimit(int64, int64) bool
	IsUpdated(int64, int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Sent(*entity.SubscriptionFollowLeague) error
	Disable(*entity.SubscriptionFollowLeague) error
	Delete(*entity.SubscriptionFollowLeague) error
	GetAllSubByLeague(int64) *[]entity.SubscriptionFollowLeague
}

func (s *SubscriptionFollowLeagueService) IsSub(subId, leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.Count(subId, leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) CountSub(subId, teamId int64) int64 {
	count, _ := s.subFollowLeagueRepo.Count(subId, teamId)
	return count
}

func (s *SubscriptionFollowLeagueService) IsLeague(leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByLeague(leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsLimit(subId, leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByLimit(subId, leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsUpdated(subId, leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByUpdated(subId, leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowLeagueRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowLeagueService) Get(subId, leagueId int64) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Get(subId, leagueId)
}

func (s *SubscriptionFollowLeagueService) Save(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Save(a)
}

func (s *SubscriptionFollowLeagueService) Update(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Update(a)
}

func (s *SubscriptionFollowLeagueService) Sent(a *entity.SubscriptionFollowLeague) error {
	if s.IsUpdated(a.SubscriptionID, a.LeagueID) {
		if s.IsLimit(a.SubscriptionID, a.LeagueID) {
			sl, err := s.Get(a.SubscriptionID, a.LeagueID)
			if err != nil {
				return err
			}
			s.subFollowLeagueRepo.Update(
				&entity.SubscriptionFollowLeague{
					SubscriptionID: a.SubscriptionID,
					LeagueID:       a.LeagueID,
					Sent:           sl.Sent + 1,
				},
			)
		}
	} else {
		// reset
		s.subFollowLeagueRepo.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: a.SubscriptionID,
				LeagueID:       a.LeagueID,
				Sent:           1,
			},
		)
	}
	return nil
}

func (s *SubscriptionFollowLeagueService) Disable(a *entity.SubscriptionFollowLeague) error {
	return s.subFollowLeagueRepo.Disable(a)
}

func (s *SubscriptionFollowLeagueService) Delete(a *entity.SubscriptionFollowLeague) error {
	return s.subFollowLeagueRepo.Delete(a)
}

func (s *SubscriptionFollowLeagueService) GetAllSubByLeague(leagueId int64) *[]entity.SubscriptionFollowLeague {
	subs, err := s.subFollowLeagueRepo.GetAllSubByLeague(leagueId)
	if err != nil {
		log.Println(err)
	}
	return subs
}
