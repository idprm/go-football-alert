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
		"é", "e",
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

type NewsSubsciption struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	TeamID         int64         `json:"team_id"`
	Team           *Team         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
}
