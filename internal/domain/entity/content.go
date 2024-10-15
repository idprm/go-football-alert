package entity

import (
	"net/url"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.Value)
	return result
}

func (e *Content) SetValueLiveMatch(day, month, price, currency string) {
	replacer := strings.NewReplacer(
		"{day}", day,
		"{month}", month,
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValueFlashNews(day, month, price, currency string) {
	replacer := strings.NewReplacer(
		"{day}", day,
		"{month}", month,
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValueSubFollowCompetition(league, day, month, price, currency string) {
	replacer := strings.NewReplacer(
		"{league}", league,
		"{day}", day,
		"{month}", month,
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValueSubFollowTeam(team, day, month, price, currency string) {
	replacer := strings.NewReplacer(
		"{team}", team,
		"{day}", day,
		"{month}", month,
		"{price}", price,
		"{currency}", url.QueryEscape(currency))
	e.Value = replacer.Replace(e.Value)
}

func (e *Content) SetValueSMSAlerteUnvalid(sc, price, currency string) {
	replacer := strings.NewReplacer(
		"{sc}", sc,
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
