package services

import (
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.ISubscriptionRepository
}

func NewSubscriptionService(
	subscriptionRepo repository.ISubscriptionRepository,
) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

type ISubscriptionService interface {
	IsSubscription(int, string) bool
	IsActiveSubscription(int, string) bool
	IsActiveSubscriptionByCategory(string, string) bool
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetByCategory(string, string) (*entity.Subscription, error)
	Get(int, string) (*entity.Subscription, error)
	Save(*entity.Subscription) (*entity.Subscription, error)
	Update(*entity.Subscription) (*entity.Subscription, error)
	Delete(*entity.Subscription) error
	IsNotActive(*entity.Subscription) (*entity.Subscription, error)
	IsNotRetry(*entity.Subscription) (*entity.Subscription, error)
	IsNotFollowTeam(*entity.Subscription) (*entity.Subscription, error)
	IsNotFollowLeague(*entity.Subscription) (*entity.Subscription, error)
	IsNotPrediction(*entity.Subscription) (*entity.Subscription, error)
	FollowLeague() *[]entity.Subscription
	FollowTeam() *[]entity.Subscription
	Prediction() *[]entity.Subscription
	CreditGoal() *[]entity.Subscription
	Renewal() *[]entity.Subscription
	Retry() *[]entity.Subscription
}

func (s *SubscriptionService) IsSubscription(serviceId int, msisdn string) bool {
	count, _ := s.subscriptionRepo.Count(serviceId, msisdn)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscription(serviceId int, msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActive(serviceId, msisdn)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscriptionByCategory(category string, msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActiveByCategory(category, msisdn)
	return count > 0
}

func (s *SubscriptionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subscriptionRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionService) GetByCategory(category, msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetByCategory(category, msisdn)
}

func (s *SubscriptionService) Get(serviceId int, msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.Get(serviceId, msisdn)
}

func (s *SubscriptionService) Save(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.Save(a)
}

func (s *SubscriptionService) Update(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.Update(a)
}

func (s *SubscriptionService) Delete(a *entity.Subscription) error {
	return s.subscriptionRepo.Delete(a)
}

func (s *SubscriptionService) IsNotActive(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotActive(a)
}

func (s *SubscriptionService) IsNotRetry(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotRetry(a)
}

func (s *SubscriptionService) IsNotFollowTeam(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotFollowTeam(a)
}

func (s *SubscriptionService) IsNotFollowLeague(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotFollowLeague(a)
}

func (s *SubscriptionService) IsNotPrediction(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotPrediction(a)
}

func (s *SubscriptionService) CreditGoal() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.CreditGoal()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Prediction() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Prediction()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) FollowTeam() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.FollowTeam()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) FollowLeague() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.FollowLeague()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Renewal() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Renewal()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Retry() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Retry()
	if err != nil {
		log.Println(err)
	}
	return subs
}
