package model

import (
	"net/url"
	"strings"
	"time"
	"unicode"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type UssdRequest struct {
	Slug       string `validate:"required" query:"slug" json:"slug,omitempty"`
	Title      string `query:"title" json:"title,omitempty"`
	Category   string `query:"category" json:"category,omitempty"`
	Package    string `query:"package" json:"package,omitempty"`
	Code       string `query:"code" json:"code,omitempty"`
	UniqueCode string `query:"unique_code" json:"unique_code,omitempty"`
	Action     string `query:"action" json:"action,omitempty"`
	LeagueId   int    `query:"league_id" json:"league_id,omitempty"`
	TeamId     int    `query:"team_id" json:"team_id,omitempty"`
	Msisdn     string `json:"msisdn,omitempty"`
	Page       int    `query:"page" json:"page,omitempty"`
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

func (m *UssdRequest) GetUniqueCode() string {
	return m.UniqueCode
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

func (m *MORequest) IsLive() bool {
	return m.GetSMS() == "LIVE"
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
func (m *UssdRequest) IsLmLiveMatchToday() bool {
	return m.GetSlug() == "lm-live-match"
}

func (m *UssdRequest) IsLmLiveMatchLater() bool {
	return m.GetSlug() == "lm-live-match-later"
}

func (m *UssdRequest) IsLmStanding() bool {
	return m.GetSlug() == "lm-standing"
}

func (m *UssdRequest) IsLmClassement() bool {
	return m.GetSlug() == "lm-classement"
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

func (m *UssdRequest) IsLiveMatch() bool {
	return m.GetSlug() == "lm-live-match" || m.GetSlug() == "lm-schedule" || m.GetSlug() == "lm-lineup" || m.GetSlug() == "lm-display"
}

func (m *UssdRequest) IsSMSAlerte() bool {
	return m.GetSlug() == "alerte-sms-competition" || m.GetSlug() == "kit-foot-by-league" || m.GetSlug() == "kit-foot-by-team"
}

func (m *UssdRequest) IsSMSAlerteEquipe() bool {
	return m.GetSlug() == "sms-alerte-equipe"
}

func (m *UssdRequest) IsSMSFootInternational() bool {
	return m.GetSlug() == "sms-foot-international"
}

func (m *UssdRequest) IsPronostic() bool {
	return m.GetSlug() == "ticket-safe" || m.GetSlug() == "combine"
}

func (m *UssdRequest) IsKitFootChamp() bool {
	return m.GetSlug() == "kit-foot-champ"
}

func (m *UssdRequest) IsCatSMSAlerteCompetition() bool {
	return m.GetCategory() == "SMSALERTE_COMPETITION"
}

func (m *UssdRequest) IsCatSMSAlerteEquipe() bool {
	return m.GetCategory() == "SMSALERTE_EQUIPE"
}

func (m *UssdRequest) IsCatPronostic() bool {
	return m.GetCategory() == "PRONOSTIC_COMBINED"
}

type MenuRequest struct {
	Category    string `validate:"required" json:"category"`
	Name        string `validate:"required" json:"name"`
	TemplateXML string `validate:"required" json:"template_xml"`
	IsActive    bool   `json:"is_active"`
}

func (r *MenuRequest) GetName() string {
	return strings.ToUpper(r.Name)
}

type ServiceRequest struct {
	Channel    string  `validate:"required" json:"channel"`
	Category   string  `validate:"required" json:"category"`
	Name       string  `validate:"required" json:"name"`
	Code       string  `validate:"required" json:"code"`
	Package    string  `validate:"required" json:"package"`
	Price      float64 `validate:"required" json:"price"`
	Currency   string  `validate:"required" json:"currency"`
	RewardGoal float64 `validate:"required" json:"reward_goal"`
	RenewalDay int     `validate:"required" json:"renewal_day"`
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

func (r *ServiceRequest) GetCode() string {
	return strings.ToUpper(r.Code)
}

type ContentRequest struct {
	Category string `validate:"required" json:"category"`
	Channel  string `validate:"required" json:"channel"`
	Name     string `validate:"required" json:"name"`
	Value    string `validate:"required" json:"value"`
}

func (r *ContentRequest) GetName() string {
	return strings.ToUpper(r.Name)
}

type LeagueRequest struct {
	PrimaryID int64  `validate:"required" json:"primary_id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	Keyword   string `json:"keyword"`
}

type TeamRequest struct {
	PrimaryID int64  `validate:"required" json:"primary_id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	Keyword   string `json:"keyword"`
}

type PronosticRequest struct {
	FixtureID int64     `validate:"required" json:"fixture_id"`
	Category  string    `validate:"required" json:"category"`
	Value     string    `validate:"required" json:"value"`
	PublishAt time.Time `validate:"required" json:"publish_at"`
}
