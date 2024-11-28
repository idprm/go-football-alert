package entity

import (
	"net/url"
	"strconv"
	"strings"
)

type Service struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	Channel    string  `gorm:"size:10;not null" json:"channel"`
	Category   string  `gorm:"size:50;not null" json:"category"`
	Name       string  `gorm:"size:50;not null" json:"name"`
	Code       string  `gorm:"size:15;index:idx_service_code,unique;not null" json:"code"`
	Package    string  `gorm:"size:50" json:"package"`
	Price      float64 `gorm:"size:15" json:"price"`
	Currency   string  `gorm:"size:10" json:"currency"`
	RewardGoal float64 `gorm:"size:15" json:"reward_goal"`
	RenewalDay int     `gorm:"size:2;default:0" json:"renewal_day"`
	FreeDay    int     `gorm:"size:2;default:0" json:"free_day"`
	UrlTelco   string  `gorm:"size:350;not null" json:"url_telco"`
	UserTelco  string  `gorm:"size:100;not null" json:"user_telco"`
	PassTelco  string  `gorm:"size:100;not null" json:"pass_telco"`
	UrlMT      string  `gorm:"size:350;not null" json:"url_mt"`
	UserMT     string  `gorm:"size:100;not null" json:"user_mt"`
	PassMT     string  `gorm:"size:100;not null" json:"pass_mt"`
	ScSubMT    string  `gorm:"size:15;not null" json:"sc_sub_mt"`
	ScUnsubMT  string  `gorm:"size:15;not null" json:"sc_unsub_mt"`
	ShortCode  string  `gorm:"size:15;not null" json:"short_code"`
	UssdCode   string  `gorm:"size:15;not null" json:"ussd_code"`
	IsActive   bool    `gorm:"type:boolean;default:false" json:"is_active"`
}

func (e *Service) GetId() int {
	return e.ID
}

func (e *Service) GetChannel() string {
	return e.Channel
}

func (e *Service) GetName() string {
	return e.Name
}

func (e *Service) GetCategory() string {
	return e.Category
}

func (e *Service) GetCode() string {
	return strings.ToUpper(e.Code)
}

func (s *Service) GetPackage() string {
	return s.Package
}

func (s *Service) GetPrice() float64 {
	return s.Price
}

func (s *Service) GetDiscount(discountPercentage int) float64 {
	return (s.GetPrice() * float64(discountPercentage) / 100)
}

func (s *Service) GetPriceToString() string {
	return strconv.FormatFloat(s.GetPrice(), 'f', 0, 64)
}

func (s *Service) GetPackagePriceToString() string {
	return "(" + strconv.FormatFloat(s.GetPrice(), 'f', 0, 64) + " " + s.GetCurrency() + "/" + s.GetPackage() + ")"
}

func (s *Service) GetCurrency() string {
	return s.Currency
}

func (s *Service) GetRenewalDay() int {
	return s.RenewalDay
}

func (s *Service) GetFreeDay() int {
	return s.FreeDay
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

func (s *Service) SetPriceWithDiscount(discountPercentage int) {
	s.Price = s.GetPrice() - (s.GetPrice() * float64(discountPercentage) / 100)
}
