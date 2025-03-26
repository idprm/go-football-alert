package entity

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	CATEGORY_LIVEMATCH             string = "LIVEMATCH"
	CATEGORY_FLASHNEWS             string = "FLASHNEWS"
	CATEGORY_SMSALERTE_COMPETITION string = "SMSALERTE_COMPETITION"
	CATEGORY_SMSALERTE_EQUIPE      string = "SMSALERTE_EQUIPE"
	CATEGORY_PRONOSTIC_SAFE        string = "PRONOSTIC_SAFE"
	CATEGORY_PRONOSTIC_COMBINED    string = "PRONOSTIC_COMBINED"
	CATEGORY_PRONOSTIC_VIP         string = "PRONOSTIC_VIP"
	CATEGORY_CREDIT_GOAL           string = "CREDIT_GOAL"
	CATEGORY_PREDICT_WIN           string = "PREDICT_WIN"
)

type Subscription struct {
	ID                   int64     `gorm:"primaryKey" json:"id"`
	ServiceID            int       `json:"service_id"`
	Service              *Service  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Category             string    `gorm:"size:30" json:"category"`
	Msisdn               string    `gorm:"size:15;not null" json:"msisdn"`
	Code                 string    `gorm:"size:25;not null" json:"code"`
	Channel              string    `gorm:"size:15" json:"channel,omitempty"`
	LatestTrxId          string    `gorm:"size:100" json:"trx_id,omitempty"`
	LatestKeyword        string    `gorm:"size:100" json:"latest_keyword,omitempty"`
	LatestSubject        string    `gorm:"size:25" json:"latest_subject,omitempty"`
	LatestStatus         string    `gorm:"size:25" json:"latest_status,omitempty"`
	LatestPayload        string    `gorm:"type:text" json:"latest_payload,omitempty"`
	LatestNote           string    `gorm:"type:text" json:"latest_note,omitempty"`
	RenewalAt            time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"renewal_at,omitempty"`
	UnsubAt              time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"unsub_at,omitempty"`
	ChargeAt             time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"charge_at,omitempty"`
	RetryAt              time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"retry_at,omitempty"`
	FreeAt               time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"free_at,omitempty"`
	FollowAt             time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"follow_at,omitempty"`
	PredictionAt         time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"prediction_at,omitempty"`
	CreditGoalAt         time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"credit_goal_at,omitempty"`
	FirstSuccessAt       time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"first_success_at,omitempty"`
	TotalSuccess         int       `gorm:"default:0" json:"total_success"`
	TotalFailed          int       `gorm:"default:0" json:"total_failed"`
	TotalAmount          float64   `gorm:"default:0" json:"total_amount"`
	TotalFirstpush       int       `gorm:"default:0" json:"total_firstpush"`
	TotalRenewal         int       `gorm:"default:0" json:"total_renewal"`
	TotalSub             int       `gorm:"default:0" json:"total_sub"`
	TotalUnsub           int       `gorm:"default:0" json:"total_unsub"`
	TotalAmountFirstpush float64   `gorm:"default:0" json:"total_amount_firstpush"`
	TotalAmountRenewal   float64   `gorm:"default:0" json:"total_amount_renewal"`
	BeforeBalance        float64   `gorm:"default:0" json:"before_balance,omitempty"`
	AfterBalance         float64   `gorm:"default:0" json:"after_balance,omitempty"`
	IpAddress            string    `gorm:"size:25" json:"ip_address"`
	AttempSuccess        int       `gorm:"size:3;default:1" json:"attemp_success"`
	IsFollowTeam         bool      `gorm:"type:boolean;column:is_follow_team" json:"is_follow_team,omitempty"`
	IsFollowLeague       bool      `gorm:"type:boolean;column:is_follow_competition" json:"is_follow_competition,omitempty"`
	IsPredictWin         bool      `gorm:"type:boolean;column:is_predict_win" json:"is_predict_win,omitempty"`
	IsCreditGoal         bool      `gorm:"type:boolean;column:is_credit_goal" json:"is_credit_goal,omitempty"`
	IsPronostic          bool      `gorm:"type:boolean;column:is_prono" json:"is_prono,omitempty"`
	IsRetry              bool      `gorm:"type:boolean;column:is_retry" json:"is_retry,omitempty"`
	IsFree               bool      `gorm:"type:boolean;column:is_free" json:"is_free,omitempty"`
	IsActive             bool      `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model
}

func (e *Subscription) GetId() int64 {
	return e.ID
}

func (e *Subscription) GetIdToString() string {
	return strconv.Itoa(int(e.ID))
}

func (e *Subscription) GetServiceId() int {
	return e.ServiceID
}

func (e *Subscription) GetMsisdn() string {
	return e.Msisdn
}

func (e *Subscription) GetCode() string {
	return e.Code
}

func (e *Subscription) GetChannel() string {
	return e.Channel
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

func (s *Subscription) GetAttempSuccess() int {
	return s.AttempSuccess
}

func (s *Subscription) SetLatestTrxId(v string) {
	s.LatestTrxId = v
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

func (s *Subscription) SetIsPredictWin(v bool) {
	s.IsPredictWin = v
}

func (s *Subscription) SetIsFollowTeam(v bool) {
	s.IsFollowLeague = v
}

func (s *Subscription) SetIsFollowLeague(v bool) {
	s.IsFollowLeague = v
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

func (e *Subscription) SetTotalSuccess(v int) {
	e.TotalSuccess = v
}

func (e *Subscription) SetTotalFailed(v int) {
	e.TotalFailed = v
}

func (e *Subscription) SetTotalAmount(v float64) {
	e.TotalAmount = v
}

func (e *Subscription) SetTotalSub(v int) {
	e.TotalSub = v
}

func (e *Subscription) SetTotalUnsub(v int) {
	e.TotalUnsub = v
}

func (e *Subscription) SetTotalAmountFirstpush(v float64) {
	e.TotalAmountFirstpush = v
}

func (e *Subscription) SetTotalAmountRenewal(v float64) {
	e.TotalAmountRenewal = v
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

func (e *Subscription) IsFirstFreeDay() bool {
	t := e.CreatedAt.Add(time.Hour * 24)
	diff := t.Sub(e.CreatedAt)
	return diff <= 24*time.Hour
}

func (e *Subscription) ISMSAlerte() bool {
	return e.Category == CATEGORY_SMSALERTE_COMPETITION || e.Category == CATEGORY_SMSALERTE_EQUIPE
}

func (e *Subscription) IsCompetition() bool {
	return e.Category == CATEGORY_SMSALERTE_COMPETITION
}

func (e *Subscription) IsEquipe() bool {
	return e.Category == CATEGORY_SMSALERTE_EQUIPE
}
