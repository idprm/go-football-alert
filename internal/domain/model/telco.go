package model

import (
	"encoding/xml"
	"strconv"

	"github.com/idprm/go-football-alert/internal/domain/entity"
)

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
				TransactionSN int64  `xml:"TransactionSN,omitempty"`
				UserPwd       string `xml:"UserPwd,omitempty"`
			} `xml:"xsd:QueryProfileAndBalRequest"`
		} `xml:"soapenv:Body"`
	}

	QueryProfileAndBalResponse struct {
		XMLName xml.Name
		Body    BodyQueryProfileAndBalResponse `xml:"Body"`
	}

	BodyQueryProfileAndBalResponse struct {
		XMLName xml.Name
		Item    ItemQueryProfileAndBalResponse `xml:"QueryProfileAndBalResponse"`
	}

	ItemQueryProfileAndBalResponse struct {
		XMLName         xml.Name
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
			BalDto []ItemBalDto `xml:"BalDto"`
		} `xml:"BalDtoList"`
	}

	ItemBalDto struct {
		BalID       string `xml:"BalID"`
		AcctResCode string `xml:"AcctResCode"`
		AcctResName string `xml:"AcctResName"`
		Balance     string `xml:"Balance"`
		EffDate     string `xml:"EffDate"`
		ExpDate     string `xml:"ExpDate"`
		UpdateDate  string `xml:"UpdateDate"`
	}
)

func (m *QueryProfileAndBalResponse) IsEnoughBalance(s *entity.Service) bool {
	var b int = 0
	if len(m.Body.Item.BalDtoList.BalDto) > 0 {
		for _, data := range m.Body.Item.BalDtoList.BalDto {
			if data.AcctResName == "Principal" {
				b, _ = strconv.Atoi(data.Balance)
			}
		}
	}
	return (b * -1) >= int(s.GetPrice())
}

func (m *QueryProfileAndBalResponse) GetBalance() string {
	return m.Body.Item.BalDtoList.BalDto[0].Balance
}

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
				TransactionSN   int64  `xml:"TransactionSN,omitempty"`
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
		XMLName xml.Name
		Body    BodyDeductFeeResponse `xml:"Body"`
	}

	BodyDeductFeeResponse struct {
		XMLName xml.Name
		Item    ItemDeductFeeResponse  `xml:"DeductFeeResponse"`
		Fault   FaultDeductFeeResponse `xml:"Fault"`
	}

	ItemDeductFeeResponse struct {
		XMLName       xml.Name
		TransactionSN string `xml:"TransactionSN"`
		AcctResCode   string `xml:"AcctResCode"`
		AcctResName   string `xml:"AcctResName"`
		BeforeBalance string `xml:"BeforeBalance"`
		AfterBalance  string `xml:"AfterBalance"`
		ExpDate       string `xml:"ExpDate"`
	}

	FaultDeductFeeResponse struct {
		XMLName     xml.Name
		FaultCode   string `xml:"faultcode"`
		FaultString string `xml:"faultstring"`
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

func (m *QueryProfileAndBalRequest) SetTransactionSN(v int64) {
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

func (m *DeductRequest) SetTransactionSN(v int64) {
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
	return m.Body.Item.TransactionSN
}

func (m *DeductResponse) GetAcctResCode() string {
	return m.Body.Item.AcctResCode
}

func (m *DeductResponse) GetAcctResName() string {
	return m.Body.Item.AcctResName
}

func (m *DeductResponse) GetBeforeBalance() string {
	return m.Body.Item.BeforeBalance
}

func (m *DeductResponse) GetBeforeBalanceToFloat() float64 {
	f, _ := strconv.ParseFloat(m.GetBeforeBalance(), 64)
	return f
}

func (m *DeductResponse) GetAfterBalance() string {
	return m.Body.Item.AfterBalance
}

func (m *DeductResponse) GetAfterBalanceToFloat() float64 {
	f, _ := strconv.ParseFloat(m.GetAfterBalance(), 64)
	return f
}

func (m *DeductResponse) GetExpDate() string {
	return m.Body.Item.ExpDate
}

func (m *DeductResponse) GetFaultCode() string {
	return m.Body.Fault.FaultCode
}

func (m *DeductResponse) GetFaultString() string {
	return m.Body.Fault.FaultString
}

func (m *DeductResponse) IsSuccess() bool {
	return m.Body.Item.AcctResCode == "1"
}

func (m *DeductResponse) IsFailed() bool {
	return m.Body.Fault.FaultCode != ""
}
