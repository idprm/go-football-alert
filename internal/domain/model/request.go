package model

import (
	"net/url"
	"strings"

	"github.com/idprm/go-football-alert/internal/domain/entity"
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

func (m *UssdRequest) GetTitleQueryEscape() string {
	return url.QueryEscape(m.GetTitle())
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

func (m *MORequest) IsInfo() bool {
	return m.GetSMS() == "INFO"
}

func (m *MORequest) IsStop() bool {
	return strings.Contains(m.GetSMS(), "STOP")
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

func (m *MORequest) IsChooseService() bool {
	return m.GetSMS() == "1" || m.GetSMS() == "2" || m.GetSMS() == "3"
}

func (m *MORequest) GetServiceByNumber() string {
	switch m.GetSMS() {
	case "1":
		return "jour"
	case "2":
		return "semaine"
	case "3":
		return "mois"
	default:
		return ""
	}
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
