package model

import (
	"encoding/xml"
	"time"
)

type CurrencyArray struct {
	CurrencyList []Currency
}

func (c *CurrencyArray) AddCurrency(currency string, amount int) {
	newc := Currency{Amount: amount}
	newc.XMLName.Local = currency
	c.CurrencyList = append(c.CurrencyList, newc)
}

type Currency struct {
	XMLName xml.Name
	Amount  int `xml:",innerxml"`
}

type Plan struct {
	XMLName           xml.Name      `xml:"plan"`
	PlanCode          string        `xml:"plan_code,omitempty"`
	CreatedAt         *time.Time    `xml:"created_at,omitempty"`
	UnitAmountInCents CurrencyArray `xml:"unit_amount_in_cents"`
	SetupFeeInCents   CurrencyArray `xml:"setup_in_cents"`
}

/**
* ####
**/
type PageResponse struct {
	XMLName xml.Name
}

type UssdResponse struct {
	XMLName xml.Name `xml:"Pages"`
	Page    struct {
		Page []struct {
		} `xml:"page"`
	} `xml:"pages"`
}

func (m *UssdResponse) SetAhref(v string) {
	// m.Pages.Page
}
