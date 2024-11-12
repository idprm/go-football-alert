package model

import (
	"net/url"
	"strings"
	"unicode"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type UssdRequest struct {
	Slug     string `validate:"required" query:"slug" json:"slug,omitempty"`
	Title    string `query:"title" json:"title,omitempty"`
	Category string `query:"category" json:"category,omitempty"`
	Package  string `query:"package" json:"package,omitempty"`
	Code     string `query:"code" json:"code,omitempty"`
	Action   string `query:"action" json:"action,omitempty"`
	LeagueId int    `query:"league_id" json:"league_id,omitempty"`
	TeamId   int    `query:"team_id" json:"team_id,omitempty"`
	Msisdn   string `json:"msisdn,omitempty"`
	Page     int    `query:"page" json:"page,omitempty"`
}

func (m *UssdRequest) GetSlug() string {
	return m.Slug
}

func (m *UssdRequest) GetTitle() string {
	return m.Title
}

func (e *UssdRequest) GetTitleWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.GetTitle())
	return result
}

func (m *UssdRequest) GetTitleQueryEscape() string {
	return url.QueryEscape(m.GetTitleWithoutAccents())
}

func (m *UssdRequest) GetCategory() string {
	return m.Category
}

func (m *UssdRequest) GetPackage() string {
	return m.Package
}

func (m *UssdRequest) GetCode() string {
	return m.Code
}

func (m *UssdRequest) GetAction() string {
	return m.Action
}

func (m *UssdRequest) GetLeagueId() int {
	return m.LeagueId
}

func (m *UssdRequest) GetTeamId() int {
	return m.TeamId
}

func (m *UssdRequest) GetMsisdn() string {
	return m.Msisdn
}

func (m *UssdRequest) GetPage() int {
	return m.Page
}

func (m *UssdRequest) SetMsisdn(v string) {
	m.Msisdn = v
}

func (m *UssdRequest) IsMsisdn() bool {
	return m.GetMsisdn() != ""
}

func (m *UssdRequest) IsCatLiveMatch() bool {
	return m.GetCategory() == "LIVEMATCH"
}

func (m *UssdRequest) IsCatFlashNews() bool {
	return m.GetCategory() == "FLASHNEWS"
}

func (m *UssdRequest) IsCatSMSAlerte() bool {
	return m.GetCategory() == "SMSALERTE"
}

type SMSRequest struct {
	Smsc     string `query:"smsc,omitempty"`
	Username string `query:"username,omitempty"`
	Password string `query:"password,omitempty"`
	From     string `query:"from,omitempty"`
	To       string `query:"to,omitempty"`
	Text     string `query:"text,omitempty"`
}

func (m *SMSRequest) GetSmsc() string {
	return m.Smsc
}

func (m *SMSRequest) GetFrom() string {
	return m.From
}

func (m *SMSRequest) GetTo() string {
	return m.To
}

func (m *SMSRequest) GetText() string {
	return strings.ToUpper(m.Text)
}

type MORequest struct {
	SMS       string `validate:"required" query:"sms" json:"sms" xml:"sms"`
	To        string `validate:"required" query:"to" json:"to" xml:"to"`
	Msisdn    string `validate:"required" query:"msisdn" json:"msisdn" xml:"msisdn"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

func (s *MORequest) GetSMS() string {
	return strings.ToUpper(s.SMS)
}

func (s *MORequest) GetTo() string {
	return s.To
}

func (s *MORequest) GetMsisdn() string {
	replacer := strings.NewReplacer("+", "")
	return replacer.Replace(s.Msisdn)
}

func (s *MORequest) GetIpAddress() string {
	return s.IpAddress
}

func (s *MORequest) GetStopKeyword() string {
	index := strings.Split(s.GetSMS(), " ")

	if index[0] == "STOP" {
		if strings.Contains(s.GetSMS(), "STOP") {
			if len(index) > 1 {
				return index[1]
			}
			return ""
		}
		return ""
	}
	return ""
}

func (m *MORequest) IsStop() bool {
	index := strings.Split(m.GetSMS(), " ")
	return index[0] == "STOP" && (strings.Contains(m.GetSMS(), "STOP"))
}

func (m *MORequest) IsProno() bool {
	return m.GetSMS() == "PRONO"
}

func (m *MORequest) IsTicket() bool {
	return m.GetSMS() == "TICKET"
}

func (m *MORequest) IsVIP() bool {
	return m.GetSMS() == "VIP"
}

func (m *MORequest) IsInfo() bool {
	return m.GetSMS() == "INFO"
}

func (m *MORequest) IsStopAlive() bool {
	return strings.Contains(m.GetSMS(), "STOP ALIVE")
}

func (m *MORequest) IsStopFlashNews() bool {
	return strings.Contains(m.GetSMS(), "STOP FLASH")
}

func (m *MORequest) IsStopAlerte() bool {
	return strings.Contains(m.GetSMS(), "STOP ALERTE")
}

func (m *MORequest) IsStopProno() bool {
	return strings.Contains(m.GetSMS(), "STOP PRONO")
}

func (m *MORequest) IsStopTicket() bool {
	return strings.Contains(m.GetSMS(), "STOP TICKET")
}

func (m *MORequest) IsStopVIP() bool {
	return strings.Contains(m.GetSMS(), "STOP VIP")
}

func (m *MORequest) IsCreditGoal(s *entity.Service) bool {
	return m.GetTo() == s.ScUnsubMT
}

func (m *MORequest) IsPrediction(s *entity.Service) bool {
	return m.GetTo() == s.ScUnsubMT
}

func (m *MORequest) IsFollowTeam(s *entity.Service) bool {
	return m.GetTo() == s.ScSubMT
}

func (m *MORequest) IsFollowLeague(s *entity.Service) bool {
	return m.GetTo() == s.ScSubMT
}

type MTRequest struct {
	TrxId        string               `json:"trx_id,omitempty"`
	Smsc         string               `json:"smsc,omitempty"`
	Keyword      string               `json:"keyword,omitempty"`
	Service      *entity.Service      `json:"service,omitempty"`
	Content      *entity.Content      `json:"content,omitempty"`
	Subscription *entity.Subscription `json:"subscription,omitempty"`
}

func (m *MTRequest) GetTrxId() string {
	return m.TrxId
}

func (m *MTRequest) SetTrxId(v string) {
	m.TrxId = v
}

func (m *MTRequest) SetKeyword(v string) {
	m.Keyword = v
}

type ErrorResponse struct {
	FailedField string `json:"failed_field" xml:"failed_field"`
	Tag         string `json:"tag" xml:"tag"`
	Value       string `json:"value" xml:"value"`
}

type CampaignSubRequest struct {
	Code      string `validate:"required" query:"code" json:"code" xml:"code"`
	Msisdn    string `validate:"required" query:"msisdn" json:"msisdn" xml:"msisdn"`
	Keyword   string `query:"keyword" json:"keyword" xml:"keyword"`
	Subkey    string `query:"subkey" json:"subkey" xml:"subkey"`
	Adnet     string `query:"adn" json:"adnet" xml:"adnet"`
	PubId     string `query:"pubid" json:"pubid" xml:"pubid"`
	ClickId   string `query:"clickid" json:"clickid" xml:"clickid"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

func (m *CampaignSubRequest) GetCode() string {
	return m.Code
}

func (m *CampaignSubRequest) GetMsisdn() string {
	return m.Msisdn
}

func (m *CampaignSubRequest) GetKeyword() string {
	return m.Keyword
}

func (m *CampaignSubRequest) GetSubkey() string {
	return m.Subkey
}

func (m *CampaignSubRequest) GetAdnet() string {
	return m.Adnet
}

func (m *CampaignSubRequest) GetPubId() string {
	return m.PubId
}

func (m *CampaignSubRequest) GetClickId() string {
	return m.ClickId
}

func (m *CampaignSubRequest) GetIpAddress() string {
	return m.IpAddress
}

func (m *CampaignSubRequest) SetCode(v string) {
	m.Code = v
}

func (m *CampaignSubRequest) SetClickId(v string) {
	m.ClickId = v
}

func (m *CampaignSubRequest) SetIpAddress(v string) {
	m.IpAddress = v
}

type CampaignUnSubRequest struct {
	Code      string `validate:"required" query:"code" json:"code" xml:"code"`
	Msisdn    string `validate:"required" query:"msisdn" json:"msisdn" xml:"msisdn"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

func (m *CampaignUnSubRequest) GetCode() string {
	return m.Code
}

func (m *CampaignUnSubRequest) GetMsisdn() string {
	return m.Msisdn
}

func (m *CampaignUnSubRequest) SetCode(v string) {
	m.Code = v
}

func (m *CampaignUnSubRequest) SetIpAddress(v string) {
	m.IpAddress = v
}

/***
**
***/
func (m *UssdRequest) IsLmLiveMatch() bool {
	return m.GetSlug() == "lm-live-match"
}

func (m *UssdRequest) IsLmSchedule() bool {
	return m.GetSlug() == "lm-schedule"
}

func (m *UssdRequest) IsFlashNews() bool {
	return m.GetSlug() == "flash-news"
}

func (m *UssdRequest) IsCreditGoal() bool {
	return m.GetSlug() == "credit-goal"
}

func (m *UssdRequest) IsChampResults() bool {
	return m.GetSlug() == "champ-results"
}

func (m *UssdRequest) IsChampStandings() bool {
	return m.GetSlug() == "champ-standings"
}

func (m *UssdRequest) IsChampSchedules() bool {
	return m.GetSlug() == "champ-schedule"
}

func (m *UssdRequest) IsChampTeam() bool {
	return m.GetSlug() == "champ-team"
}

func (m *UssdRequest) IsChampCreditScore() bool {
	return m.GetSlug() == "champ-credit-score"
}

func (m *UssdRequest) IsChampCreditGoal() bool {
	return m.GetSlug() == "champ-credit-goal"
}

func (m *UssdRequest) IsChampSMSAlerte() bool {
	return m.GetSlug() == "champ-sms-alerte"
}

func (m *UssdRequest) IsChampSMSAlerteEquipe() bool {
	return m.GetSlug() == "champ-sms-alerte-equipe"
}

func (m *UssdRequest) IsPrediction() bool {
	return m.GetSlug() == "prediction"
}

func (m *UssdRequest) IsSMSAlerte() bool {
	return m.GetSlug() == "sms-alerte"
}

func (m *UssdRequest) IsSMSAlerteEquipe() bool {
	return m.GetSlug() == "sms-alerte-equipe"
}

func (m *UssdRequest) IsSMSFootInternational() bool {
	return m.GetSlug() == "sms-foot-international"
}

func (m *UssdRequest) IsKitFoot() bool {
	return m.GetSlug() == "kit-foot"
}

func (m *UssdRequest) IsKitFootChamp() bool {
	return m.GetSlug() == "kit-foot-champ"
}

type MenuRequest struct {
	Category    string `json:"category"`
	Name        string `json:"name"`
	TemplateXML string `json:"template_xml"`
	IsActive    bool   `json:"is_active"`
}

// Channel    string  `gorm:"size:10;not null" json:"channel"`
// Category   string  `gorm:"size:50;not null" json:"category"`
// Name       string  `gorm:"size:50;not null" json:"name"`
// Code       string  `gorm:"size:15;not null" json:"code"`
// Package    string  `gorm:"size:50" json:"package"`
// Price      float64 `gorm:"size:15" json:"price"`
// Currency   string  `gorm:"size:10" json:"currency"`
// RewardGoal float64 `gorm:"size:15" json:"reward_goal"`
// RenewalDay int     `gorm:"size:2;default:0" json:"renewal_day"`
// FreeDay    int     `gorm:"size:2;default:0" json:"free_day"`
// UrlTelco   string  `gorm:"size:350;not null" json:"url_telco"`
// UserTelco  string  `gorm:"size:100;not null" json:"user_telco"`
// PassTelco  string  `gorm:"size:100;not null" json:"pass_telco"`
// UrlMT      string  `gorm:"size:350;not null" json:"url_mt"`
// UserMT     string  `gorm:"size:100;not null" json:"user_mt"`
// PassMT     string  `gorm:"size:100;not null" json:"pass_mt"`
// ScSubMT    string  `gorm:"size:15;not null" json:"sc_sub_mt"`
// ScUnsubMT  string  `gorm:"size:15;not null" json:"sc_unsub_mt"`
// ShortCode  string  `gorm:"size:15;not null" json:"short_code"`
// UssdCode   string  `gorm:"size:15;not null" json:"ussd_code"`
// IsActive   bool    `gorm:"type:boolean;default:false" json:"is_active"`

type ServiceRequest struct {
	Channel    string  `json:"channel"`
	Category   string  `json:"category"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	Package    string  `json:"package"`
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	RewardGoal float64 `json:"reward_goal"`
	RenewalDay int     `json:"renewal_day"`
	FreeDay    int     `json:"free_day"`
	UrlTelco   string  `json:"url_telco"`
	UserTelco  string  `json:"user_telco"`
	PassTelco  string  `json:"pass_telco"`
	UrlMT      string  `json:"url_mt"`
	UserMT     string  `json:"user_mt"`
	PassMT     string  `json:"pass_mt"`
	ScSubMT    string  `json:"sc_sub_mt"`
	ScUnsubMT  string  `json:"sc_unsub_mt"`
	ShortCode  string  `json:"short_code"`
	UssdCode   string  `json:"ussd_code"`
	IsActive   bool    `json:"is_active"`
}
