package services

import (
	"strconv"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/utils"
)

type ContentService struct {
	contentRepo repository.IContentRepository
}

func NewContentService(contentRepo repository.IContentRepository) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
	}
}

type IContentService interface {
	IsContent(string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetFollowCompetition(string, *entity.Service, *entity.League) (*entity.Content, error)
	GetFollowTeam(string, *entity.Service, *entity.Team) (*entity.Content, error)
	Get(string) (*entity.Content, error)
	Save(*entity.Content) (*entity.Content, error)
	Update(*entity.Content) (*entity.Content, error)
	Delete(*entity.Content) error
}

func (s *ContentService) IsContent(name string) bool {
	count, _ := s.contentRepo.Count(name)
	return count > 0
}

func (s *ContentService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.contentRepo.GetAllPaginate(pagination)
}

func (s *ContentService) GetFollowCompetition(name string, service *entity.Service, league *entity.League) (*entity.Content, error) {
	c, err := s.contentRepo.Get(name)
	if err != nil {
		return nil, err
	}
	c.SetValueSubFollowCompetition(league.GetName(), strconv.Itoa(time.Now().Day()), utils.FormatFROnlyMonth(time.Now()), service.GetPriceToString(), service.GetCurrency())
	return c, nil
}

func (s *ContentService) GetFollowTeam(name string, service *entity.Service, team *entity.Team) (*entity.Content, error) {
	c, err := s.contentRepo.Get(name)
	if err != nil {
		return nil, err
	}
	c.SetValueSubFollowTeam(team.GetName(), strconv.Itoa(time.Now().Day()), utils.FormatFROnlyMonth(time.Now()), service.GetPriceToString(), service.GetCurrency())
	return c, nil
}

func (s *ContentService) Get(name string) (*entity.Content, error) {
	return s.contentRepo.Get(name)
}

func (s *ContentService) Save(a *entity.Content) (*entity.Content, error) {
	return s.contentRepo.Save(a)
}

func (s *ContentService) Update(a *entity.Content) (*entity.Content, error) {
	return s.contentRepo.Update(a)
}

func (s *ContentService) Delete(a *entity.Content) error {
	return s.contentRepo.Delete(a)
}
