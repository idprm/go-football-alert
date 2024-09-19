package entity

import "time"

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
	return e.Title
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
