package entity

import (
	"time"

	"gorm.io/gorm"
)

// type PredictionResponse struct {
// 	Winner struct {
// 		PrimaryID int    `json:"id"`
// 		Name      string `json:"name"`
// 		Comment   string `json:"comment"`
// 	} `json:"winner"`
// 	Advice  string `json:"advice"`
// 	Percent struct {
// 		Home string `json:"home"`
// 		Draw string `json:"draw"`
// 		Away string `json:"away"`
// 	} `json:"percent"`
// }

type Prediction struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	FixtureID     int64     `json:"fixture_id"`
	Fixture       *Fixture  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"fixture,omitempty"`
	FixtureDate   time.Time `json:"fixture_date"`
	WinnerID      int64     `json:"winner_id"`
	Winner        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"winner,omitempty"`
	WinnerName    string    `json:"winner_name"`
	WinnerComment string    `json:"winner_comment"`
	Advice        string    `json:"advice"`
	PercentHome   string    `json:"percent_home"`
	PercentDraw   string    `json:"percent_draw"`
	PercentAway   string    `json:"percent_away"`
	gorm.Model    `json:"-"`
}

func (e *Prediction) GetId() int64 {
	return e.ID
}

func (e *Prediction) GetFixtureId() int64 {
	return e.FixtureID
}

type PredictionSubsciption struct {
	ID             int64         `gorm:"primaryKey" json:"id"`
	SubscriptionID int64         `json:"subscription_id"`
	Subscription   *Subscription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subscription,omitempty"`
	TeamID         int64         `json:"team_id"`
	Team           *Team         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
}
