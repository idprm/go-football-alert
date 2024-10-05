package entity

import (
	"net/url"
	"strings"
	"time"
)

type News struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:300;not null" json:"title"`
	Slug        string    `gorm:"size:300;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	Source      string    `gorm:"size:45" json:"source"`
	PublishAt   time.Time `json:"publish_at"`
}

func (e *News) GetId() int64 {
	return e.ID
}

func (e *News) GetTitle() string {
	replacer := strings.NewReplacer(
		`"`, "",
	)
	return replacer.Replace(e.Title)

}

func (e *News) GetTitleLimited(maxLength int) string {
	if len(e.Title) >= maxLength {
		return e.GetTitle()[:maxLength]
	}
	return e.GetTitle()
}

func (e *News) GetTitleQueryEscape() string {
	return url.QueryEscape(e.GetTitle())
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

func (e *News) GetParseTitle() string {
	if strings.Contains(e.GetTitle(), ":") {
		return strings.TrimSpace(e.Title[:strings.IndexByte(e.Title, ':')])
	}
	return ""
}

func (e *News) GetHomeTeam() string {
	return strings.TrimSpace(e.GetParseTitle()[:strings.IndexByte(e.GetParseTitle(), '-')])
}

func (e *News) GetAwayTeam() string {
	return strings.TrimSpace(e.GetParseTitle()[strings.IndexByte(e.GetParseTitle(), '-')+1:])
}

func (e *News) IsParseTitle() bool {
	return e.GetParseTitle() != ""
}

func (e *News) IsMatch() bool {
	return strings.Contains(e.GetParseTitle(), "-")
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

type NewsSubsciption struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	TeamID         int64         `json:"team_id"`
	Team           *Team         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
}
