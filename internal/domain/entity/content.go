package entity

import (
	"net/url"
	"strings"
)

type Content struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Category string `gorm:"size:20" json:"category"`
	Channel  string `gorm:"size:15" json:"channel"`
	Name     string `gorm:"size:50" json:"name"`
	Value    string `gorm:"size:250" json:"value"`
}

func (e *Content) GetId() int {
	return e.ID
}

func (e *Content) GetCategory() string {
	return e.Category
}

func (e *Content) GetChannel() string {
	return e.Channel
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
