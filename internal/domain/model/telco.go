package model

import "encoding/xml"

type (
	QueryProfileAndBalRequest struct {
		XMLName xml.Name `xml:"soapenv:Envelope"`
		Soapenv string   `xml:"xmlns:soapenv,attr"`
		Xsd     string   `xml:"xmlns:xsd,attr"`
		Header  struct {
			AuthHeader struct {
				Username string `xml:"Username,omitempty"`
				Password string `xml:"Password,omitempty"`
			} `xml:"xsd:AuthHeader"`
		} `xml:"soapenv:Header"`
		Body struct {
			QueryProfileAndBal struct {
				Msisdn        string `xml:"MSISDN,omitempty"`
				TransactionSN string `xml:"TransactionSN,omitempty"`
				UserPwd       string `xml:"UserPwd,omitempty"`
			} `xml:"xsd:QueryProfileAndBalRequest"`
		} `xml:"soapenv:Body"`
	}

	QueryProfileAndBalResponse struct {
		XMLName xml.Name `xml:"Envelope"`
		SoapEnv string   `xml:"xmlns:soapenv,attr"`
		Body    struct {
			QueryProfileAndBal struct {
				Msisdn          string `xml:"MSISDN,omitempty"`
				DefLang         string `xml:"DefLang,omitempty"`
				State           string `xml:"State,omitempty"`
				StateSet        string `xml:"StateSet,omitempty"`
				ActiveStopDate  string `xml:"ActiveStopDate,omitempty"`
				SuspendStopDate string `xml:"SuspendStopDate,omitempty"`
				DisableStopDate string `xml:"DisableStopDate,omitempty"`
				ServiceStopDate string `xml:"ServiceStopDate,omitempty"`
				BrandIndex      string `xml:"BrandIndex,omitempty"`
				ServiceClass    string `xml:"ServiceClass,omitempty"`
				TransactionSN   string `xml:"TransactionSN,omitempty"`
				BalDtoList      struct {
				} `xml:"BalDtoList"`
			} `xml:"QueryProfileAndBalResponse"`
		} `xml:"soapenv:Body"`
	}
)

type (
	DeductRequest struct {
		XMLName xml.Name `xml:"soapenv:Envelope"`
		Soapenv string   `xml:"xmlns:soapenv,attr"`
		Xsd     string   `xml:"xmlns:xsd,attr"`
		Header  struct {
			AuthHeader struct {
				Username string `xml:"Username"`
				Password string `xml:"Password"`
			} `xml:"xsd:AuthHeader"`
		} `xml:"soapenv:Header"`
		Body struct {
			DeductFee struct {
				TransactionSN   string `xml:"TransactionSN,omitempty"`
				TransactionDesc string `xml:"TransactionDesc,omitempty"`
				ChannelID       string `xml:"Channel_ID,omitempty"`
				Msisdn          string `xml:"MSISDN,omitempty"`
				AccountCode     string `xml:"AccountCode,omitempty"`
				AcctResCode     string `xml:"AcctResCode,omitempty"`
				DeductBalance   string `xml:"DeductBalance,omitempty"`
			} `xml:"xsd:DeductFeeRequest"`
		} `xml:"soapenv:Body"`
	}

	DeductResponse struct {
		XMLName xml.Name `xml:"Envelope"`
		SoapEnv string   `xml:"xmlns:soapenv,attr"`
		Body    struct {
			DeductFee struct {
				TransactionSN string `xml:"TransactionSN,omitempty"`
				AcctResCode   string `xml:"AcctResCode,omitempty"`
				AcctResName   string `xml:"AcctResName,omitempty"`
				BeforeBalance string `xml:"BeforeBalance,omitempty"`
				AfterBalance  string `xml:"AfterBalance,omitempty"`
				ExpDate       string `xml:"ExpDate,omitempty"`
			} `xml:"DeductFeeResponse"`
			Fault struct {
				FaultCode   string `xml:"faultcode,omitempty"`
				FaultString string `xml:"faultstring,omitempty"`
			} `xml:"soapenv:Fault"`
		} `xml:"soapenv:Body"`
	}
)

func (e *QueryProfileAndBalRequest) SetSoap(v string) {
	e.Soapenv = v
}

func (e *QueryProfileAndBalRequest) SetXsd(v string) {
	e.Xsd = v
}

func (m *QueryProfileAndBalRequest) SetUsername(v string) {
	m.Header.AuthHeader.Username = v
}

func (m *QueryProfileAndBalRequest) SetPassword(v string) {
	m.Header.AuthHeader.Password = v
}

func (m *QueryProfileAndBalRequest) SetMsisdn(v string) {
	m.Body.QueryProfileAndBal.Msisdn = v
}

func (m *QueryProfileAndBalRequest) SetTransactionSN(v string) {
	m.Body.QueryProfileAndBal.TransactionSN = v
}

func (m *QueryProfileAndBalRequest) SetUserPwd(v string) {
	m.Body.QueryProfileAndBal.UserPwd = v
}

func (e *DeductRequest) SetSoap(v string) {
	e.Soapenv = v
}

func (e *DeductRequest) SetXsd(v string) {
	e.Xsd = v
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
	return m.Body.DeductFee.TransactionSN
}

func (m *DeductResponse) GetAcctResCode() string {
	return m.Body.DeductFee.AcctResCode
}

func (m *DeductResponse) GetAcctResName() string {
	return m.Body.DeductFee.AcctResName
}

func (m *DeductResponse) GetBeforeBalance() string {
	return m.Body.DeductFee.BeforeBalance
}

func (m *DeductResponse) GetAfterBalance() string {
	return m.Body.DeductFee.AfterBalance
}

func (m *DeductResponse) GetExpDate() string {
	return m.Body.DeductFee.ExpDate
}

func (m *DeductResponse) GetFaultCode() string {
	return m.Body.Fault.FaultCode
}

func (m *DeductResponse) GetFaultString() string {
	return m.Body.Fault.FaultString
}

func (m *DeductResponse) IsFailed() bool {
	return m.Body.Fault.FaultCode == "" || m.Body.Fault.FaultString == "" || m == nil
}
