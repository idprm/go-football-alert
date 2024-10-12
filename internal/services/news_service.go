package services

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type NewsService struct {
	newsRepo repository.INewsRepository
}

func NewNewsService(newsRepo repository.INewsRepository) *NewsService {
	return &NewsService{
		newsRepo: newsRepo,
	}
}

type INewsService interface {
	IsNews(string, string) bool
	IsNewsLeague(leagueId int64) bool
	IsNewsTeam(teamId int64) bool
	IsNewsById(int64) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetAllUSSD(int) ([]*entity.News, error)
	GetByTeamUSSD(int) (*entity.News, error)
	Get(string, string) (*entity.News, error)
	GetById(int64) (*entity.News, error)
	GetBySlug(string) (*entity.News, error)
	Save(*entity.News) (*entity.News, error)
	Update(*entity.News) (*entity.News, error)
	Delete(*entity.News) error
	GetAllNewsLeague(leagueId int64) ([]*entity.NewsLeagues, error)
	GetAllNewsTeam(teamId int64) ([]*entity.NewsTeams, error)
	SaveNewsLeague(*entity.NewsLeagues) (*entity.NewsLeagues, error)
	UpdateNewsLeague(*entity.NewsLeagues) (*entity.NewsLeagues, error)
	SaveNewsTeam(*entity.NewsTeams) (*entity.NewsTeams, error)
	UpdateNewsTeam(*entity.NewsTeams) (*entity.NewsTeams, error)
}

func (s *NewsService) IsNews(slug, pubAt string) bool {
	count, _ := s.newsRepo.Count(slug, pubAt)
	return count > 0
}

func (s *NewsService) IsNewsLeague(leagueId int64) bool {
	count, _ := s.newsRepo.CountNewsLeague(leagueId)
	return count > 0
}

func (s *NewsService) IsNewsTeam(teamId int64) bool {
	count, _ := s.newsRepo.CountNewsTeam(teamId)
	return count > 0
}

func (s *NewsService) IsNewsById(id int64) bool {
	count, _ := s.newsRepo.CountById(id)
	return count > 0
}

func (s *NewsService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.newsRepo.GetAllPaginate(pagination)
}

func (s *NewsService) GetAllUSSD(page int) ([]*entity.News, error) {
	return s.newsRepo.GetAllUSSD(page)
}

func (s *NewsService) GetByTeamUSSD(teamId int) (*entity.News, error) {
	return s.newsRepo.GetByTeamUSSD(teamId)
}

func (s *NewsService) Get(slug, pubAt string) (*entity.News, error) {
	return s.newsRepo.Get(slug, pubAt)
}

func (s *NewsService) GetById(id int64) (*entity.News, error) {
	return s.newsRepo.GetById(id)
}

func (s *NewsService) GetBySlug(slug string) (*entity.News, error) {
	return s.newsRepo.GetBySlug(slug)
}

func (s *NewsService) Save(a *entity.News) (*entity.News, error) {
	return s.newsRepo.Save(a)
}

func (s *NewsService) Update(a *entity.News) (*entity.News, error) {
	return s.newsRepo.Update(a)
}

func (s *NewsService) Delete(a *entity.News) error {
	return s.newsRepo.Delete(a)
}

func (s *NewsService) GetAllNewsLeague(leagueId int64) ([]*entity.NewsLeagues, error) {
	return s.newsRepo.GetAllNewsLeague(leagueId)
}

func (s *NewsService) GetAllNewsTeam(teamId int64) ([]*entity.NewsTeams, error) {
	return s.newsRepo.GetAllNewsTeam(teamId)
}

func (s *NewsService) SaveNewsLeague(a *entity.NewsLeagues) (*entity.NewsLeagues, error) {
	return s.newsRepo.SaveNewsLeague(a)
}

func (s *NewsService) UpdateNewsLeague(a *entity.NewsLeagues) (*entity.NewsLeagues, error) {
	return s.newsRepo.UpdateNewsLeague(a)
}

func (s *NewsService) SaveNewsTeam(a *entity.NewsTeams) (*entity.NewsTeams, error) {
	return s.newsRepo.SaveNewsTeam(a)
}

func (s *NewsService) UpdateNewsTeam(a *entity.NewsTeams) (*entity.NewsTeams, error) {
	return s.newsRepo.UpdateNewsTeam(a)
}
