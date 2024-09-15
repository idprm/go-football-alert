package entity

import (
	"net/url"
	"strings"
)

type Content struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Name      string   `gorm:"size:50" json:"name"`
	Value     string   `gorm:"size:250" json:"value"`
}

func (e *Content) GetId() int {
	return e.ID
}

func (e *Content) GetServiceId() int {
	return e.ServiceID
}

func (e *Content) GetName() string {
	return e.Name
}

func (e *Content) GetValue() string {
	return e.Value
}

func (e *Content) SetValueSMSAlert(date, price, currency string) {
	replacer := strings.NewReplacer(
		"{date}", date,
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValuePrediction(home, away, credit, price, currency string) {
	replacer := strings.NewReplacer(
		"{home}", url.QueryEscape(home),
		"{away}", url.QueryEscape(away),
		"{credit}", url.QueryEscape(credit),
		"{price}", price,
		"{currency}", url.QueryEscape(currency),
	)
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValueCreditGoal(home, away, score, credit, price, currency string) {
	replacer := strings.NewReplacer(
		"{home}", url.QueryEscape(home),
		"{away}", url.QueryEscape(away),
		"{score}", url.QueryEscape(score),
		"{credit}", url.QueryEscape(credit),
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}
