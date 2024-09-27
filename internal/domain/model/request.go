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
	Slug   string `validate:"required" query:"slug" json:"slug"`
	Msisdn string `json:"msisdn"`
}

func (m *UssdRequest) GetSlug() string {
	return m.Slug
}

func (m *UssdRequest) GetMsisdn() string {
	return m.Msisdn
}

func (m *UssdRequest) SetMsisdn(v string) {
	m.Msisdn = v
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
