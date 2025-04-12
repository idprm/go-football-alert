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
	IsSubscription(int, string, string) bool
	IsActiveSubscription(int, string, string) bool
	IsActiveSubscriptionByCategory(string, string, string) bool
	IsActiveSubscriptionByNonSMSAlerte(string, string) bool
	IsActiveSubscriptionBySMSAlerteMsisdn(string) bool
	IsActiveSubscriptionBySubId(int64) bool
	IsActiveAllByMsisdn(string) bool
	IsAfter24Hour(int, string, string) bool
	IsRenewal(int, string, string) bool
	IsRetry(int, string, string) bool
	IsRetryUnderpayment(int, string, string) bool
	GetTotalActiveSubscription() int
	GetAllPaginate(*entity.Pagination) (*entity.Pagination, error)
	GetByCategory(string, string, string) (*entity.Subscription, error)
	GetActiveByCategory(string, string, string) (*entity.Subscription, error)
	GetActiveByNonSMSAlerte(string, string) (*entity.Subscription, error)
	GetActiveBySMSAlerteMsisdn(string) (*entity.Subscription, error)
	GetByNonSMSAlerte(string, string) (*entity.Subscription, error)
	GetBySubId(int64) (*entity.Subscription, error)
	GetActiveAllByMsisdnUSSD(string, int) (*[]entity.Subscription, error)
	Get(int, string, string) (*entity.Subscription, error)
	GetLongRetry(int, string, string) (int, error)
	Save(*entity.Subscription) (*entity.Subscription, error)
	Update(*entity.Subscription) (*entity.Subscription, error)
	Delete(*entity.Subscription) error
	UpdateNotActive(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFree(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotRetry(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotUnderpayment(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFollowTeam(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotFollowLeague(*entity.Subscription) (*entity.Subscription, error)
	UpdateNotPredictWin(*entity.Subscription) (*entity.Subscription, error)
	PredictWin() *[]entity.Subscription
	CreditGoal() *[]entity.Subscription
	Follow() *[]entity.Subscription
	Prono() *[]entity.Subscription
	Renewal() *[]entity.Subscription
	Retry() *[]entity.Subscription
	RetryUnderpayment() *[]entity.Subscription
	Reminder48HBeforeCharging() *[]entity.Subscription
	ReminderAfterTrialEnds() *[]entity.Subscription
	CountActiveSub(int) (int, error)
	GetAllSubBySMSAlerte() (*[]entity.Subscription, error)
}

func (s *SubscriptionService) IsSubscription(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.Count(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscription(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountActive(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscriptionByCategory(category, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountActiveByCategory(category, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscriptionByNonSMSAlerte(category, msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActiveByNonSMSAlerte(category, msisdn)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscriptionBySMSAlerteMsisdn(msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActiveBySMSAlerteMsisdn(msisdn)
	return count > 0
}

func (s *SubscriptionService) IsActiveSubscriptionBySubId(subId int64) bool {
	count, _ := s.subscriptionRepo.CountActiveBySubId(subId)
	return count > 0
}

func (s *SubscriptionService) IsActiveAllByMsisdn(msisdn string) bool {
	count, _ := s.subscriptionRepo.CountActiveAllByMsisdn(msisdn)
	return count > 0
}

func (s *SubscriptionService) IsAfter24Hour(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountAfter24Hour(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsRenewal(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountRenewal(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsRetry(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountRetry(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) IsRetryUnderpayment(serviceId int, msisdn, code string) bool {
	count, _ := s.subscriptionRepo.CountRetryUnderpayment(serviceId, msisdn, code)
	return count > 0
}

func (s *SubscriptionService) GetTotalActiveSubscription() int {
	count, _ := s.subscriptionRepo.CountTotalActiveSub()
	return int(count)
}

func (s *SubscriptionService) GetAllPaginate(pagination *entity.Pagination) (*entity.Pagination, error) {
	return s.subscriptionRepo.GetAllPaginate(pagination)
}

func (s *SubscriptionService) GetByCategory(category, msisdn, code string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetByCategory(category, msisdn, code)
}

func (s *SubscriptionService) GetActiveByCategory(category, msisdn, code string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetActiveByCategory(category, msisdn, code)
}

func (s *SubscriptionService) GetByNonSMSAlerte(category, msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetByNonSMSAlerte(category, msisdn)
}

func (s *SubscriptionService) GetActiveByNonSMSAlerte(category, msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetActiveByNonSMSAlerte(category, msisdn)
}

func (s *SubscriptionService) GetActiveBySMSAlerteMsisdn(msisdn string) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetActiveBySMSAlerteMsisdn(msisdn)
}

func (s *SubscriptionService) GetBySubId(subId int64) (*entity.Subscription, error) {
	return s.subscriptionRepo.GetBySubId(subId)
}

func (s *SubscriptionService) GetActiveAllByMsisdnUSSD(msisdn string, page int) (*[]entity.Subscription, error) {
	return s.subscriptionRepo.GetActiveAllByMsisdnUSSD(msisdn, page)
}

func (s *SubscriptionService) Get(serviceId int, msisdn, code string) (*entity.Subscription, error) {
	return s.subscriptionRepo.Get(serviceId, msisdn, code)
}

func (s *SubscriptionService) GetLongRetry(serviceId int, msisdn, code string) (int, error) {
	return s.subscriptionRepo.GetLongRetry(serviceId, msisdn, code)
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

func (s *SubscriptionService) UpdateNotActive(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotActive(a)
}

func (s *SubscriptionService) UpdateNotFree(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotFree(a)
}

func (s *SubscriptionService) UpdateNotRetry(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotRetry(a)
}

func (s *SubscriptionService) UpdateNotUnderpayment(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotUnderpayment(a)
}

func (s *SubscriptionService) UpdateNotFollowTeam(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotFollowTeam(a)
}

func (s *SubscriptionService) UpdateNotFollowLeague(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotFollowLeague(a)
}

func (s *SubscriptionService) UpdateNotPredictWin(a *entity.Subscription) (*entity.Subscription, error) {
	return s.subscriptionRepo.UpdateNotPredictWin(a)
}

func (s *SubscriptionService) CreditGoal() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.CreditGoal()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) PredictWin() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.PredictWin()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Follow() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Follow()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Prono() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Prono()
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

func (s *SubscriptionService) RetryUnderpayment() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.RetryUnderpayment()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) Reminder48HBeforeCharging() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.Reminder48HBeforeCharging()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) ReminderAfterTrialEnds() *[]entity.Subscription {
	subs, err := s.subscriptionRepo.ReminderAfterTrialEnds()
	if err != nil {
		log.Println(err)
	}
	return subs
}

func (s *SubscriptionService) CountActiveSub(serviceId int) (int, error) {
	r, err := s.subscriptionRepo.CountActiveSub(serviceId)
	if err != nil {
		return 0, err
	}
	return int(r), nil
}

func (s *SubscriptionService) GetAllSubBySMSAlerte() (*[]entity.Subscription, error) {
	subs, err := s.subscriptionRepo.GetAllSubBySMSAlerte()
	if err != nil {
		return nil, err
	}
	return subs, nil
}
