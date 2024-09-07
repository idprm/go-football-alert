package entity

import (
	"net/url"
	"strings"
)

type Service struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	CountryID  int      `json:"country_id"`
	Country    *Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	Category   string   `gorm:"size:50;not null" json:"category"`
	Name       string   `gorm:"size:50;not null" json:"name"`
	Code       string   `gorm:"size:15;not null" json:"code"`
	Package    string   `gorm:"size:50" json:"package"`
	Price      float64  `gorm:"size:15" json:"price"`
	CreditGoal float64  `gorm:"size:15" json:"credit_goal"`
	RenewalDay int      `json:"renewal_day"`
	TrialDay   int      `json:"trial_day"`
	UrlTelco   string   `gorm:"size:350;not null" json:"url_telco"`
	UserTelco  string   `gorm:"size:100;not null" json:"user_telco"`
	PassTelco  string   `gorm:"size:100;not null" json:"pass_telco"`
	UrlMT      string   `gorm:"size:350;not null" json:"url_mt"`
	UserMT     string   `gorm:"size:100;not null" json:"user_mt"`
	PassMT     string   `gorm:"size:100;not null" json:"pass_mt"`
	ScSubMT    string   `gorm:"size:15;not null" json:"sc_sub_mt"`
	ScUnsubMT  string   `gorm:"size:15;not null" json:"sc_unsub_mt"`
	UssdCode   string   `gorm:"size:15;not null" json:"ussd_code"`
}

func (e *Service) GetId() int {
	return e.ID
}

func (e *Service) GetCountryId() int {
	return e.CountryID
}

func (e *Service) GetName() string {
	return e.Name
}

func (e *Service) GetCategory() string {
	return e.Category
}

func (e *Service) GetCode() string {
	return e.Code
}

func (s *Service) GetPackage() string {
	return s.Package
}

func (s *Service) GetPrice() float64 {
	return s.Price
}

func (s *Service) GetCreditGoal() float64 {
	return s.CreditGoal
}

func (s *Service) GetRenewalDay() int {
	return s.RenewalDay
}

func (s *Service) GetTrialDay() int {
	return s.TrialDay
}

func (e *Service) GetUrlTelco() string {
	return e.UrlTelco
}

func (e *Service) GetUserTelco() string {
	return e.UserTelco
}

func (e *Service) GetPassTelco() string {
	return e.PassTelco
}

func (e *Service) GetUrlMT() string {
	return e.UrlMT
}

func (e *Service) SetUrlMT(smsc, username, password, from, to, content string) {

	replacer := strings.NewReplacer(
		"{smsc}", url.QueryEscape(smsc),
		"{username}", url.QueryEscape(username),
		"{password}", url.QueryEscape(password),
		"{from}", url.QueryEscape(from),
		"{to}", url.QueryEscape(to),
		"{text}", url.QueryEscape(content))

	e.UrlMT = replacer.Replace(e.UrlMT)
}
