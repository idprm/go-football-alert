package model

import "strings"

const (
	MO_REG   string = "REG"
	MO_UNREG string = "UNREG"
)

type UssdRequest struct {
	Text string `validate:"required" form:"text"`
}

func (m *UssdRequest) GetText() string {
	return m.Text
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
	return m.GetText() == "1*1"
}

func (m *UssdRequest) IsSchedule() bool {
	return m.GetText() == "1*2"
}

func (m *UssdRequest) IsLineup() bool {
	return m.GetText() == "1*3"
}

func (m *UssdRequest) IsMatchStats() bool {
	return m.GetText() == "1*4"
}

func (m *UssdRequest) IsDisplayLiveMatch() bool {
	return m.GetText() == "1*5"
}

func (m *UssdRequest) IsFlashNews() bool {
	return m.GetText() == "2"
}

func (m *UssdRequest) IsCreditGoal() bool {
	return m.GetText() == "3"
}

func (m *UssdRequest) IsChampResults() bool {
	return m.GetText() == "4*1"
}

func (m *UssdRequest) IsChampStandings() bool {
	return m.GetText() == "4*2"
}

func (m *UssdRequest) IsChampSchedule() bool {
	return m.GetText() == "4*3"
}

func (m *UssdRequest) IsChampTeam() bool {
	return m.GetText() == "4*4"
}

func (m *UssdRequest) IsChampCreditScore() bool {
	return m.GetText() == "4*5"
}

func (m *UssdRequest) IsChampCreditGoal() bool {
	return m.GetText() == "4*6"
}

func (m *UssdRequest) IsChampSMSAlerte() bool {
	return m.GetText() == "4*7"
}

func (m *UssdRequest) IsChampSMSAlerteEquipe() bool {
	return m.GetText() == "4*8"
}

func (m *UssdRequest) IsPrediction() bool {
	return m.GetText() == "5"
}

func (m *UssdRequest) IsKitFoot() bool {
	return m.GetText() == "6*1"
}

func (m *UssdRequest) IsEurope() bool {
	return m.GetText() == "6*2"
}

func (m *UssdRequest) IsAfrique() bool {
	return m.GetText() == "6*3"
}

func (m *UssdRequest) IsSMSAlerteEquipe() bool {
	return m.GetText() == "6*4"
}

func (m *UssdRequest) IsFootInternational() bool {
	return m.GetText() == "6*5"
}

func (m *UssdRequest) IsAlerteChampMaliEquipe() bool {
	return m.GetText() == "7*1"
}

func (m *UssdRequest) IsAlertePremierLeagueEquipe() bool {
	return m.GetText() == "7*2"
}

func (m *UssdRequest) IsAlerteLaLigaEquipe() bool {
	return m.GetText() == "7*3"
}

func (m *UssdRequest) IsAlerteLigue1Equipe() bool {
	return m.GetText() == "7*4"
}

func (m *UssdRequest) IsAlerteSerieAEquipe() bool {
	return m.GetText() == "7*5"
}

func (m *UssdRequest) IsAlerteBundesligueEquipe() bool {
	return m.GetText() == "7*6"
}

func (m *UssdRequest) IsChampionLeague() bool {
	return m.GetText() == "8*1"
}

func (m *UssdRequest) IsPremierLeague() bool {
	return m.GetText() == "8*2"
}

func (m *UssdRequest) IsLaLiga() bool {
	return m.GetText() == "8*3"
}

func (m *UssdRequest) IsLigue1() bool {
	return m.GetText() == "8*4"
}

func (m *UssdRequest) IsLEuropa() bool {
	return m.GetText() == "8*5"
}

func (m *UssdRequest) IsSerieA() bool {
	return m.GetText() == "8*6"
}

func (m *UssdRequest) IsBundesligua() bool {
	return m.GetText() == "8*7"
}

func (m *UssdRequest) IsChampPortugal() bool {
	return m.GetText() == "8*8"
}

func (m *UssdRequest) IsSaudiLeague() bool {
	return m.GetText() == "8*9"
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
