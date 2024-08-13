package entity

import "time"

type News struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	FixtureID   int64     `json:"fixture_id"`
	Fixture     *Fixture  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	Title       string    `gorm:"size:300;not null" json:"title"`
	Slug        string    `gorm:"size:300;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (e *News) GetId() int64 {
	return e.ID
}

func (e *News) GetFixtureId() int64 {
	return e.FixtureID
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

func (e *News) GetCreatedAt() time.Time {
	return e.CreatedAt
}
