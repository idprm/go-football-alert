package entity

import (
	"net/url"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

type News struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:300;not null" json:"title"`
	Slug        string    `gorm:"size:300;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	Source      string    `gorm:"size:45" json:"source"`
	PublishAt   time.Time `json:"publish_at"`
	gorm.Model
}

func (e *News) GetId() int64 {
	return e.ID
}

func (e *News) GetTitle() string {
	replacer := strings.NewReplacer(
		`"`, "",
		`’`, "'",
	)
	return replacer.Replace(e.Title)
}

func (e *News) GetTitleLimited(maxLength int) string {
	if len(e.Title) >= maxLength {
		return e.GetTitleWithoutAccents()[:maxLength]
	}
	return e.GetTitleWithoutAccents()
}

func (e *News) GetTitleWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.GetTitle())
	return result
}

func (e *News) GetWithoutAccent(v string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, v)
	return result
}

func (e *News) GetTitleQueryEscape() string {
	return url.QueryEscape(e.GetTitleWithoutAccents())
}

func (e *News) GetSlug() string {
	return e.Slug
}

func (e *News) GetDescription() string {
	return e.Description
}

func (e *News) GetSource() string {
	return e.Source
}

func (e *News) GetPublishAt() time.Time {
	return e.PublishAt
}

func (e *News) GetParseTitleLeft() string {
	t := strings.TrimSpace(e.Title[:strings.IndexByte(e.Title, ':')])
	replacer := strings.NewReplacer(
		`Mercato`, "",
		`OM`, "",
	)
	if len(strings.TrimSpace(replacer.Replace(t))) > 0 {
		return strings.TrimSpace(replacer.Replace(t))
	}
	return "-"
}

func (e *News) GetParseTitleRight() string {
	return strings.TrimSpace(e.Title[strings.IndexByte(e.Title, ':')+1:])
}

func (e *News) GetHomeTeam() string {
	return strings.TrimSpace(e.GetParseTitleLeft()[:strings.IndexByte(e.GetParseTitleLeft(), '-')])
}

func (e *News) GetAwayTeam() string {
	return strings.TrimSpace(e.GetParseTitleLeft()[strings.IndexByte(e.GetParseTitleLeft(), '-')+1:])
}

func (e *News) SetId(v int64) {
	e.ID = v
}

func (e *News) IsHeadTitle() bool {
	return strings.Contains(e.GetTitle(), ":")
}

func (e *News) IsMatch() bool {
	return strings.Contains(e.GetParseTitleLeft(), "-")
}

func (e *News) IsMaxiFoot() bool {
	return e.Source == "MAXIFOOT"
}

func (e *News) IsMadeInFoot() bool {
	return e.Source == "MADEINFOOT"
}

func (e *News) IsAfricaTopSports() bool {
	return e.Source == "AFRICATOPSPORTS"
}

func (e *News) IsFootMercato() bool {
	return e.Source == "FOOTMERCATO"
}

func (e *News) IsMobimiumNews() bool {
	return e.Source == "MOBIMIUMNEWS"
}

func (m *News) IsFilteredKeyword(v string) bool {
	return v == "OL"
}

type NewsLeagues struct {
	ID         int64   `gorm:"primaryKey" json:"id"`
	NewsID     int64   `json:"news_id"`
	News       *News   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	LeagueID   int64   `json:"league_id"`
	League     *League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	gorm.Model `json:"-"`
}

type NewsTeams struct {
	ID         int64 `gorm:"primaryKey" json:"id"`
	NewsID     int64 `json:"news_id"`
	News       *News `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"news,omitempty"`
	TeamID     int64 `json:"team_id"`
	Team       *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	gorm.Model `json:"-"`
}
