package entity

import (
	"net/url"
	"time"

	"gorm.io/gorm"
)

type Fixture struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	PrimaryID   int64     `json:"primary_id"`
	Timezone    string    `json:"timezone"`
	FixtureDate time.Time `json:"fixture_date"`
	TimeStamp   int       `json:"timestamp"`
	LeagueID    int64     `json:"league_id"`
	League      *League   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	HomeID      int64     `json:"home_id"`
	Home        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"home,omitempty"`
	AwayID      int64     `json:"away_id"`
	Away        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"away,omitempty"`
	Goal        string    `json:"goal"`
	IsDone      bool      `gorm:"type:boolean;default:false" json:"is_done"`
	gorm.Model  `json:"-"`
}

func (e *Fixture) GetId() int64 {
	return e.ID
}

func (e *Fixture) GetPrimaryId() int64 {
	return e.PrimaryID
}

func (e *Fixture) GetTimezone() string {
	return e.Timezone
}

func (e *Fixture) GetDate() time.Time {
	return e.FixtureDate
}

func (e *Fixture) GetFixtureName() string {
	return e.Home.GetName() + " - " + e.Away.GetName() + " (" + e.GetFixtureDateToString() + ")"
}

func (e *Fixture) GetFixtureAndTimeName() string {
	return e.Home.GetName() + " - " + e.Away.GetName() + " (" + e.GetFixtureDateAndTimeToString() + ")"
}

func (e *Fixture) GetFixtureNameQueryEscape() string {
	return url.QueryEscape(e.GetFixtureAndTimeName())
}

func (e *Fixture) GetFixtureDateToString() string {
	return e.FixtureDate.Format("2 Jan 06")
}

func (e *Fixture) GetFixtureDateAndTimeToString() string {
	return e.FixtureDate.Format("2 Jan 06 15:04")
}

func (e *Fixture) GetTimeStamp() int {
	return e.TimeStamp
}

func (e *Fixture) GetHomeId() int64 {
	return e.HomeID
}

func (e *Fixture) GetAwayId() int64 {
	return e.AwayID
}
