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
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	Get(int, string) (*entity.Subscription, error)
	Save(*entity.Subscription) (*entity.Subscription, error)
	Update(*entity.Subscription) (*entity.Subscription, error)
	Delete(*entity.Subscription) error
	IsNotActive(*entity.Subscription) (*entity.Subscription, error)
	IsNotRetry(*entity.Subscription) (*entity.Subscription, error)
	IsNotFollow(*entity.Subscription) (*entity.Subscription, error)
	IsNotPrediction(*entity.Subscription) (*entity.Subscription, error)
	FollowCompetition() *[]entity.Subscription
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

func (s *SubscriptionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subscriptionRepo.GetAllPaginate(pagination)
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

func (s *SubscriptionService) IsNotFollow(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotFollow(a)
}

func (s *SubscriptionService) IsNotPrediction(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.IsNotPrediction(a)
}

func (s *SubscriptionService) FollowCompetition() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.FollowCompetition()
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

func (s *SubscriptionService) Prediction() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Prediction()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) CreditGoal() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.CreditGoal()
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
