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
	IsSub(int64) bool
	IsLeague(int64) bool
	IsLimit(int64) bool
	IsUpdated(int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetBySub(int64) (*entity.SubscriptionFollowLeague, error)
	Save(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Update(*entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error)
	Disable(*entity.SubscriptionFollowLeague) error
	Delete(*entity.SubscriptionFollowLeague) error
	GetAllSubByLeague(int64) *[]entity.SubscriptionFollowLeague
}

func (s *SubscriptionFollowLeagueService) IsSub(subId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountBySub(subId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsLeague(leagueId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByLeague(leagueId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsLimit(subId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByLimit(subId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) IsUpdated(subId int64) bool {
	count, _ := s.subFollowLeagueRepo.CountByUpdated(subId)
	return count > 0
}

func (s *SubscriptionFollowLeagueService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowLeagueRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowLeagueService) GetBySub(subId int64) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.GetBySub(subId)
}

func (s *SubscriptionFollowLeagueService) Save(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Save(a)
}

func (s *SubscriptionFollowLeagueService) Update(a *entity.SubscriptionFollowLeague) (*entity.SubscriptionFollowLeague, error) {
	return s.subFollowLeagueRepo.Update(a)
}

func (s *SubscriptionFollowLeagueService) Sent(a *entity.SubscriptionFollowLeague) error {
	if s.IsUpdated(a.SubscriptionID) {
		if s.IsLimit(a.SubscriptionID) {
			sl, err := s.GetBySub(a.SubscriptionID)
			if err != nil {
				return err
			}
			s.subFollowLeagueRepo.Update(
				&entity.SubscriptionFollowLeague{
					SubscriptionID: a.SubscriptionID,
					Sent:           sl.Sent + 1,
				},
			)
		}
	}

	if !s.IsUpdated(a.SubscriptionID) {
		// reset
		s.subFollowLeagueRepo.Update(
			&entity.SubscriptionFollowLeague{
				SubscriptionID: a.SubscriptionID,
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
