package model

import (
	"strings"

	"github.com/idprm/go-football-alert/internal/domain/entity"
)

const (
	MO_REG   string = "REG"
	MO_UNREG string = "UNREG"
)

type UssdRequest struct {
	ServiceCode string `validate:"required" form:"serviceCode" json:"service_code"`
	PhoneNumber string `validate:"required" form:"phoneNumber" json:"phone_number"`
	Text        string `validate:"required" form:"text" json:"text"`
	Action      string `json:"action,omitempty"`
}

func (m *UssdRequest) GetServiceCode() string {
	return m.ServiceCode
}

func (m *UssdRequest) GetMsisdn() string {
	return m.PhoneNumber
}

func (m *UssdRequest) GetText() string {
	return m.Text
}

func (m *UssdRequest) GetAction() string {
	return m.Action
}

func (m *UssdRequest) SetAction(v string) {
	m.Action = v
}

func (m *UssdRequest) IsREG() bool {
	return m.Action == "REG"
}

func (m *UssdRequest) IsFilterLevel1() bool {
	filter := []string{
		"1", "2", "3", "4*", "5", "6", "7",
	}
	for _, s := range filter {
		if strings.HasPrefix(m.GetText(), s) {
			return true
		}
	}
	return false
}

func (m *UssdRequest) IsFilterLevel3() bool {
	filter := []string{
		"0", "1*1*", "1*2*", "1*3*", "1*4*",
		"1*5*", "2*", "3*", "4*1*", "4*2*",
		"4*3*", "4*4*", "4*5*", "4*6*", "4*7*",
		"4*8*", "5*", "6*1*", "6*2*", "6*3*",
		"6*4*", "6*5*", "7*1*", "7*2*", "7*3*",
		"7*4*", "7*5*", "7*6*", "8*1*", "8*2*",
		"8*3*", "8*4*", "8*5*", "8*6*", "8*7*",
		"8*8*", "8*9*",
	}
	for _, s := range filter {
		if strings.HasPrefix(m.GetText(), s) {
			return true
		}
	}
	return false
}

func (m *UssdRequest) IsMain() bool {
	return m.Text == "" || m.Text == "98" || m.Text == "99"
}

func (m *UssdRequest) IsPrev() bool {
	return m.Text == "0"
}

func (m *UssdRequest) IsLevel() bool {
	return strings.Contains(m.GetText(), "*")
}

func (m *UssdRequest) IsLiveMatch() bool {
	return strings.HasPrefix(m.GetText(), "1*1*")
}

func (m *UssdRequest) IsSchedule() bool {
	return strings.HasPrefix(m.GetText(), "1*2*")
}

func (m *UssdRequest) IsLineup() bool {
	return strings.HasPrefix(m.GetText(), "1*3*")
}

func (m *UssdRequest) IsMatchStats() bool {
	return strings.HasPrefix(m.GetText(), "1*4*")
}

func (m *UssdRequest) IsDisplayLiveMatch() bool {
	return strings.HasPrefix(m.GetText(), "1*5*")
}

func (m *UssdRequest) IsFlashNews() bool {
	return strings.HasPrefix(m.GetText(), "2*")
}

func (m *UssdRequest) IsCreditGoal() bool {
	return strings.HasPrefix(m.GetText(), "3*")
}

func (m *UssdRequest) IsChampResults() bool {
	return strings.HasPrefix(m.GetText(), "4*1*")
}

func (m *UssdRequest) IsChampStandings() bool {
	return strings.HasPrefix(m.GetText(), "4*2*")
}

func (m *UssdRequest) IsChampSchedule() bool {
	return strings.HasPrefix(m.GetText(), "4*3*")
}

func (m *UssdRequest) IsChampTeam() bool {
	return strings.HasPrefix(m.GetText(), "4*4*")
}

func (m *UssdRequest) IsChampCreditScore() bool {
	return strings.HasPrefix(m.GetText(), "4*5*")
}

func (m *UssdRequest) IsChampCreditGoal() bool {
	return strings.HasPrefix(m.GetText(), "4*6*")
}

func (m *UssdRequest) IsChampSMSAlerte() bool {
	return strings.HasPrefix(m.GetText(), "4*7*")
}

func (m *UssdRequest) IsChampSMSAlerteEquipe() bool {
	return strings.HasPrefix(m.GetText(), "4*8*")
}

func (m *UssdRequest) IsPrediction() bool {
	return strings.HasPrefix(m.GetText(), "5*")
}

func (m *UssdRequest) IsKitFoot() bool {
	return strings.HasPrefix(m.GetText(), "6*1*")
}

func (m *UssdRequest) IsEurope() bool {
	return strings.HasPrefix(m.GetText(), "6*2*")
}

func (m *UssdRequest) IsAfrique() bool {
	return strings.HasPrefix(m.GetText(), "6*3*")
}

func (m *UssdRequest) IsSMSAlerteEquipe() bool {
	return strings.HasPrefix(m.GetText(), "6*4*")
}

func (m *UssdRequest) IsFootInternational() bool {
	return strings.HasPrefix(m.GetText(), "6*5*")
}

func (m *UssdRequest) IsAlerteChampMaliEquipe() bool {
	return strings.HasPrefix(m.GetText(), "7*1*")
}

func (m *UssdRequest) IsAlertePremierLeagueEquipe() bool {
	return strings.HasPrefix(m.GetText(), "7*2*")
}

func (m *UssdRequest) IsAlerteLaLigaEquipe() bool {
	return strings.HasPrefix(m.GetText(), "7*3*")
}

func (m *UssdRequest) IsAlerteLigue1Equipe() bool {
	return strings.HasPrefix(m.GetText(), "7*4*")
}

func (m *UssdRequest) IsAlerteSerieAEquipe() bool {
	return strings.HasPrefix(m.GetText(), "7*5*")
}

func (m *UssdRequest) IsAlerteBundesligueEquipe() bool {
	return strings.HasPrefix(m.GetText(), "7*6*")
}

func (m *UssdRequest) IsChampionLeague() bool {
	return strings.HasPrefix(m.GetText(), "8*1*")
}

func (m *UssdRequest) IsPremierLeague() bool {
	return strings.HasPrefix(m.GetText(), "8*2*")
}

func (m *UssdRequest) IsLaLiga() bool {
	return strings.HasPrefix(m.GetText(), "8*3*")
}

func (m *UssdRequest) IsLigue1() bool {
	return strings.HasPrefix(m.GetText(), "8*4*")
}

func (m *UssdRequest) IsLEuropa() bool {
	return strings.HasPrefix(m.GetText(), "8*5*")
}

func (m *UssdRequest) IsSerieA() bool {
	return strings.HasPrefix(m.GetText(), "8*6*")
}

func (m *UssdRequest) IsBundesligua() bool {
	return strings.HasPrefix(m.GetText(), "8*7*")
}

func (m *UssdRequest) IsChampPortugal() bool {
	return strings.HasPrefix(m.GetText(), "8*8*")
}

func (m *UssdRequest) IsSaudiLeague() bool {
	return strings.HasPrefix(m.GetText(), "8*9*")
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

func (m *SMSRequest) IsInfo() bool {
	return m.GetText() == "INFO"
}

func (m *SMSRequest) IsStop() bool {
	return m.GetText() == "STOP"
}

func (m *SMSRequest) IsCreditGoal() bool {
	return m.GetSmsc() == "8021"
}

func (m *SMSRequest) IsPrediction() bool {
	return m.GetSmsc() == "8033"
}

func (m *SMSRequest) IsFollowTeam() bool {
	return m.GetSmsc() == "8023"
}

func (m *SMSRequest) IsFollowCompetition() bool {
	return m.GetSmsc() == "8024"
}

func (m *SMSRequest) IsChooseService() bool {
	return m.GetText() == "1" || m.GetText() == "2" || m.GetText() == "3"
}

func (m *SMSRequest) GetServiceByNumber() string {
	switch m.GetText() {
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

type MORequest struct {
	SMS       string `validate:"required" json:"sms" xml:"sms"`
	Msisdn    string `validate:"required" json:"msisdn" xml:"msisdn"`
	IpAddress string `query:"ip_address" json:"ip_address"`
}

func (s *MORequest) GetKeyword() string {
	return strings.ToUpper(s.SMS)
}

func (s *MORequest) GetMsisdn() string {
	return s.Msisdn
}

func (s *MORequest) GetSubKeyword() string {
	message := strings.ToUpper(s.SMS)
	index := strings.Split(message, " ")

	if index[0] == MO_REG || index[0] == MO_UNREG {
		if strings.Contains(message, MO_REG) || strings.Contains(message, MO_UNREG) {
			if len(index) > 1 {
				return index[1]
			}
			return ""
		}
		return ""
	}
	return ""
}

func (s *MORequest) IsREG() bool {
	message := strings.ToUpper(s.SMS)
	index := strings.Split(message, " ")
	if index[0] == MO_REG && (strings.Contains(message, MO_REG)) {
		return true
	}
	return false
}

func (s *MORequest) IsUNREG() bool {
	message := strings.ToUpper(s.SMS)
	index := strings.Split(message, " ")
	if index[0] == MO_UNREG && (strings.Contains(message, MO_UNREG)) {
		return true
	}
	return false
}

func (s *MORequest) GetIpAddress() string {
	return s.IpAddress
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
