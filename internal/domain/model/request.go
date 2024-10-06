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

func (m *UssdRequest) IsYes() bool {
	return m.GetAction() == "yes"
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
	return s.Msisdn
}

func (s *MORequest) GetIpAddress() string {
	return s.IpAddress
}

func (m *MORequest) IsInfo() bool {
	return m.GetSMS() == "INFO"
}

func (m *MORequest) IsStop() bool {
	return m.GetSMS() == "STOP"
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

func (m *MORequest) IsFollowCompetition(s *entity.Service) bool {
	return m.GetTo() == s.ScSubMT
}

func (m *MORequest) IsChooseService() bool {
	return m.GetSMS() == "1" || m.GetSMS() == "2" || m.GetSMS() == "3"
}

func (m *MORequest) GetServiceByNumber() string {
	switch m.GetSMS() {
	case "1":
		return "daily"
	case "2":
		return "weekly"
	case "3":
		return "monthly"
	default:
		return ""
	}
}

type MTRequest struct {
	Smsc         string               `json:"smsc,omitempty"`
	Content      *entity.Content      `json:"content,omitempty"`
	Subscription *entity.Subscription `json:"subscription,omitempty"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field" xml:"failed_field"`
	Tag         string `json:"tag" xml:"tag"`
	Value       string `json:"value" xml:"value"`
}

// Read the variables sent via POST from our API
type AfricasTalkingRequest struct {
	SessionId   string `form:"sessionId" json:"session_id"`
	ServiceCode string `form:"serviceCode" json:"service_code"`
	PhoneNumber string `form:"phoneNumber" json:"phone_number"`
	Text        string `form:"text" json:"text"`
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
	return m.GetSlug() == "champ-creditgoal"
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

func (m *UssdRequest) IsSMSKitFoot() bool {
	return m.GetSlug() == "sms-kit-foot"
}

func (m *UssdRequest) IsSMSFootEurope() bool {
	return m.GetSlug() == "sms-foot-europe"
}

func (m *UssdRequest) IsSMSFootAfrique() bool {
	return m.GetSlug() == "sms-foot-afrique"
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

func (m *UssdRequest) IsKitFootPremierLeague() bool {
	return m.GetSlug() == "kit-foot-premier-league"
}

func (m *UssdRequest) IsPremierLeagues() bool {
	return m.GetSlug() == "foot-europe-premier-league"
}

func (m *UssdRequest) IsLaLigas() bool {
	return m.GetSlug() == "foot-europe-la-liga"
}

func (m *UssdRequest) IsLigue1s() bool {
	return m.GetSlug() == "foot-europe-ligue-1"
}

func (m *UssdRequest) IsSerieA() bool {
	return m.GetSlug() == "foot-europe-serie-a"
}

// if req.GetSlug() == "foot-europe-serie-a" {
// data = h.SerieA(req.GetPage() + 1)
// }

// if req.GetSlug() == "foot-europe-bundesligua" {
// data = h.Bundesliguas(req.GetPage() + 1)
// }

// if req.GetSlug() == "foot-europe-champ-portugal" {
// data = h.PrimeiraLigas(req.GetPage() + 1)
// }
