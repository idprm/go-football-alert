package services

import (
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionFollowTeamService struct {
	subFollowTeamRepo repository.ISubscriptionFollowTeamRepository
}

func NewSubscriptionFollowTeamService(
	subFollowTeamRepo repository.ISubscriptionFollowTeamRepository,
) *SubscriptionFollowTeamService {
	return &SubscriptionFollowTeamService{
		subFollowTeamRepo: subFollowTeamRepo,
	}
}

type ISubscriptionFollowTeamService interface {
	IsSub(int64, int64) bool
	CountSub(int64, int64) int64
	IsTeam(int64) bool
	IsLimit(int64, int64) bool
	IsUpdated(int64, int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int64, int64) (*entity.SubscriptionFollowTeam, error)
	Save(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Update(*entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error)
	Sent(*entity.SubscriptionFollowTeam) error
	Disable(*entity.SubscriptionFollowTeam) error
	Delete(*entity.SubscriptionFollowTeam) error
	GetAllSubByTeam(int64) *[]entity.SubscriptionFollowTeam
}

func (s *SubscriptionFollowTeamService) IsSub(subId, teamId int64) bool {
	count, _ := s.subFollowTeamRepo.Count(subId, teamId)
	return count > 0
}

func (s *SubscriptionFollowTeamService) CountSub(subId, teamId int64) int64 {
	count, _ := s.subFollowTeamRepo.Count(subId, teamId)
	return count
}

func (s *SubscriptionFollowTeamService) IsTeam(teamId int64) bool {
	count, _ := s.subFollowTeamRepo.CountByTeam(teamId)
	return count > 0
}

func (s *SubscriptionFollowTeamService) IsLimit(subId, teamId int64) bool {
	count, _ := s.subFollowTeamRepo.CountByLimit(subId, teamId)
	return count > 0
}

func (s *SubscriptionFollowTeamService) IsUpdated(subId, teamId int64) bool {
	count, _ := s.subFollowTeamRepo.CountByUpdated(subId, teamId)
	return count > 0
}

func (s *SubscriptionFollowTeamService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subFollowTeamRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionFollowTeamService) Get(subId, teamId int64) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Get(subId, teamId)
}

func (s *SubscriptionFollowTeamService) Save(a *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Save(a)
}

func (s *SubscriptionFollowTeamService) Update(a *entity.SubscriptionFollowTeam) (*entity.SubscriptionFollowTeam, error) {
	return s.subFollowTeamRepo.Update(a)
}

func (s *SubscriptionFollowTeamService) Sent(a *entity.SubscriptionFollowTeam) error {
	if s.IsUpdated(a.SubscriptionID, a.TeamID) {
		if s.IsLimit(a.SubscriptionID, a.TeamID) {
			sl, err := s.Get(a.SubscriptionID, a.TeamID)
			if err != nil {
				return err
			}
			s.subFollowTeamRepo.Update(
				&entity.SubscriptionFollowTeam{
					SubscriptionID: a.SubscriptionID,
					TeamID:         a.TeamID,
					Sent:           sl.Sent + 1,
				},
			)
		}
	} else {
		// reset
		s.subFollowTeamRepo.Update(
			&entity.SubscriptionFollowTeam{
				SubscriptionID: a.SubscriptionID,
				TeamID:         a.TeamID,
				Sent:           1,
			},
		)
	}
	return nil
}

func (s *SubscriptionFollowTeamService) Disable(a *entity.SubscriptionFollowTeam) error {
	return s.subFollowTeamRepo.Disable(a)
}

func (s *SubscriptionFollowTeamService) Delete(a *entity.SubscriptionFollowTeam) error {
	return s.subFollowTeamRepo.Delete(a)
}

func (s *SubscriptionFollowTeamService) GetAllSubByTeam(teamId int64) *[]entity.SubscriptionFollowTeam {
	subs, err := s.subFollowTeamRepo.GetAllSubByTeam(teamId)
	if err != nil {
		log.Println(err)
	}
	return subs
}
