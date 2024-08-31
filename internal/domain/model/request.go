package model

import "strings"

const (
	MO_REG       string = "REG"
	MO_UNREG     string = "UNREG"
	VALID_PREFIX string = "_"
)

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

func (r *AfricasTalkingRequest) IsFirst() bool {
	return r.Text == ""
}

func (r *AfricasTalkingRequest) IsOne() bool {
	return r.Text == "1"
}

func (r *AfricasTalkingRequest) IsTwo() bool {
	return r.Text == "2"
}
