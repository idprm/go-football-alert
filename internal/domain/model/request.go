package model

type MORequest struct {
	Msisdn string `validate:"required" json:"msisdn"`
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
