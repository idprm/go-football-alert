package model

import "encoding/xml"

type (
	DeductRequest struct {
		XMLName xml.Name `xml:"soapenv:Envelope"`
		Soapenv string   `xml:"xmlns:soapenv,attr"`
		Header  struct {
			AuthHeader struct {
				Username string `xml:"Username"`
				Password string `xml:"Password"`
			} `xml:"xsd:AuthHeader"`
		} `xml:"soapenv:Header"`
		Body struct {
			DeductFee struct {
				TransactionSN   string `xml:"TransactionSN"`
				TransactionDesc string `xml:"TransactionDesc"`
				ChannelID       string `xml:"Channel_ID"`
				Msisdn          string `xml:"MSISDN"`
				AccountCode     string `xml:"AccountCode"`
				AcctResCode     string `xml:"AcctResCode"`
				DeductBalance   string `xml:"DeductBalance"`
			} `xml:"xsd:DeductFeeRequest"`
		} `xml:"soapenv:Body"`
	}

	DeductResponse struct {
		XMLName xml.Name `xml:"Envelope"`
		SoapEnv string   `xml:"xmlns:soapenv,attr"`
		Body    struct {
			DeductFeeResponse struct {
				TransactionSN string `xml:"TransactionSN"`
				AcctResCode   string `xml:"AcctResCode"`
				AcctResName   string `xml:"AcctResName"`
				BeforeBalance string `xml:"BeforeBalance"`
				AfterBalance  string `xml:"AfterBalance"`
				ExpDate       string `xml:"ExpDate"`
				FaultCode     string `xml:"faultcode"`
				FaultString   string `xml:"faultstring"`
			} `xml:"xsd:DeductFeeResponse"`
		} `xml:"soapenv:Body"`
	}
)

func (e *DeductRequest) SetSoap(v string) {
	e.Soapenv = v
}

func (m *DeductRequest) SetUsername(v string) {
	m.Header.AuthHeader.Username = v
}

func (m *DeductRequest) SetPassword(v string) {
	m.Header.AuthHeader.Password = v
}

func (m *DeductRequest) SetTransactionSN(v string) {
	m.Body.DeductFee.TransactionSN = v
}

func (m *DeductRequest) SetTransactionDesc(v string) {
	m.Body.DeductFee.TransactionDesc = v
}

func (m *DeductRequest) SetChannelID(v string) {
	m.Body.DeductFee.ChannelID = v
}

func (m *DeductRequest) SetMsisdn(v string) {
	m.Body.DeductFee.Msisdn = v
}

func (m *DeductRequest) SetAccountCode(v string) {
	m.Body.DeductFee.AccountCode = v
}

func (m *DeductRequest) SetAcctResCode(v string) {
	m.Body.DeductFee.AcctResCode = v
}

func (m *DeductRequest) SetDeductBalance(v string) {
	m.Body.DeductFee.DeductBalance = v
}

func (m *DeductResponse) GetTransactionSN() string {
	return m.Body.DeductFeeResponse.TransactionSN
}

func (m *DeductResponse) GetAcctResCode() string {
	return m.Body.DeductFeeResponse.AcctResCode
}

func (m *DeductResponse) GetAcctResName() string {
	return m.Body.DeductFeeResponse.AcctResName
}

func (m *DeductResponse) GetBeforeBalance() string {
	return m.Body.DeductFeeResponse.BeforeBalance
}

func (m *DeductResponse) GetAfterBalance() string {
	return m.Body.DeductFeeResponse.AfterBalance
}

func (m *DeductResponse) GetExpDate() string {
	return m.Body.DeductFeeResponse.ExpDate
}

func (m *DeductResponse) GetFaultCode() string {
	return m.Body.DeductFeeResponse.FaultCode
}

func (m *DeductResponse) GetFaultString() string {
	return m.Body.DeductFeeResponse.FaultString
}

func (m *DeductResponse) IsFailed() bool {
	return m.Body.DeductFeeResponse.FaultCode != "" || m.Body.DeductFeeResponse.FaultString != ""
}
