package entity

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID                   int64     `gorm:"primaryKey" json:"id"`
	CountryID            int       `json:"country_id"`
	Country              *Country  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	ServiceID            int       `json:"service_id"`
	Service              *Service  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Category             string    `gorm:"size:20" json:"category"`
	Msisdn               string    `gorm:"size:15;not null" json:"msisdn"`
	Channel              string    `gorm:"size:15" json:"channel,omitempty"`
	LatestTrxId          string    `gorm:"size:100" json:"trx_id,omitempty"`
	LatestKeyword        string    `gorm:"size:50" json:"latest_keyword,omitempty"`
	LatestSubject        string    `gorm:"size:25" json:"latest_subject,omitempty"`
	LatestStatus         string    `gorm:"size:25" json:"latest_status,omitempty"`
	LatestPayload        string    `gorm:"type:text" json:"latest_payload,omitempty"`
	RenewalAt            time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"renewal_at,omitempty"`
	UnsubAt              time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"unsub_at,omitempty"`
	ChargeAt             time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"charge_at,omitempty"`
	RetryAt              time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"retry_at,omitempty"`
	TrialAt              time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"trial_at,omitempty"`
	FollowAt             time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"follow_at,omitempty"`
	PredictionAt         time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"prediction_at,omitempty"`
	CreditGoalAt         time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"credit_goal_at,omitempty"`
	FirstSuccessAt       time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"first_success_at,omitempty"`
	TotalSuccess         int       `gorm:"default:0" json:"total_success,omitempty"`
	TotalFailed          int       `gorm:"default:0" json:"total_failed,omitempty"`
	TotalAmount          float64   `gorm:"default:0" json:"total_amount,omitempty"`
	TotalFirstpush       int       `gorm:"default:0" json:"total_firstpush,omitempty"`
	TotalRenewal         int       `gorm:"default:0" json:"total_renewal,omitempty"`
	TotalSub             int       `gorm:"default:0" json:"total_sub,omitempty"`
	TotalUnsub           int       `gorm:"default:0" json:"total_unsub,omitempty"`
	TotalAmountFirstpush float64   `gorm:"default:0" json:"total_amount_firstpush,omitempty"`
	TotalAmountRenewal   float64   `gorm:"default:0" json:"total_amount_renewal,omitempty"`
	BeforeBalance        float64   `gorm:"default:0" json:"before_balance,omitempty"`
	AfterBalance         float64   `gorm:"default:0" json:"after_balance,omitempty"`
	ExpireAt             time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"expire_at,omitempty"`
	IpAddress            string    `gorm:"size:25" json:"ip_address,omitempty"`
	IsFollow             bool      `gorm:"type:boolean;column:is_follow" json:"is_follow,omitempty"`
	IsPrediction         bool      `gorm:"type:boolean;column:is_prediction" json:"is_prediction,omitempty"`
	IsCreditGoal         bool      `gorm:"type:boolean;column:is_credit_goal" json:"is_credit_goal,omitempty"`
	IsRetry              bool      `gorm:"type:boolean;column:is_retry" json:"is_retry,omitempty"`
	IsTrial              bool      `gorm:"type:boolean;column:is_trial" json:"is_trial,omitempty"`
	IsActive             bool      `gorm:"type:boolean;column:is_active" json:"is_active,omitempty"`
	gorm.Model           `json:"-"`
}

func (e *Subscription) GetId() int64 {
	return e.ID
}

func (e *Subscription) GetServiceId() int {
	return e.ServiceID
}

func (e *Subscription) GetMsisdn() string {
	return e.Msisdn
}

func (s *Subscription) GetLatestTrxId() string {
	return s.LatestTrxId
}

func (s *Subscription) GetLatestKeyword() string {
	return s.LatestKeyword
}

func (s *Subscription) GetLatestSubject() string {
	return strings.ToUpper(s.LatestSubject)
}

func (s *Subscription) GetLatestStatus() string {
	return s.LatestStatus
}

func (s *Subscription) GetIpAddress() string {
	return s.IpAddress
}

func (s *Subscription) SetLatestSubject(v string) {
	s.LatestSubject = v
}

func (s *Subscription) SetLatestStatus(v string) {
	s.LatestStatus = v
}

func (s *Subscription) SetLatestPayload(v string) {
	s.LatestPayload = v
}

func (s *Subscription) SetIsPrediction(v bool) {
	s.IsPrediction = v
}

func (s *Subscription) SetIsFollow(v bool) {
	s.IsFollow = v
}

func (s *Subscription) SetIsRetry(v bool) {
	s.IsRetry = v
}

func (s *Subscription) SetIsActive(v bool) {
	s.IsActive = v
}

func (s *Subscription) SetRenewalAt(v time.Time) {
	s.RenewalAt = v
}

func (s *Subscription) SetRetryAt(v time.Time) {
	s.RetryAt = v
}

func (s *Subscription) SetChargeAt(v time.Time) {
	s.ChargeAt = v
}

func (s *Subscription) SetUnsubAt(v time.Time) {
	s.UnsubAt = v
}

func (s *Subscription) IsCreatedAtToday() bool {
	return s.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02")
}

func (s *Subscription) IsRetryAtToday() bool {
	return s.RetryAt.Format("2006-01-02") == time.Now().Format("2006-01-02")
}

func (s *Subscription) IsFirstpush() bool {
	return s.GetLatestSubject() == "FIRSTPUSH"
}

func (s *Subscription) IsRenewal() bool {
	return s.GetLatestSubject() == "RENEWAL"
}
