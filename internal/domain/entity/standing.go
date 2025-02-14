package entity

import (
	"net/url"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Standing struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	LeagueID    int64     `json:"league_id"`
	League      *League   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	TeamID      int64     `json:"team_id"`
	Team        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	Ranking     int       `json:"ranking"`
	TeamName    string    `json:"team_name"`
	Points      int       `json:"points"`
	GoalsDiff   int       `json:"goalsDiff"`
	Group       string    `json:"group"`
	Form        string    `json:"form"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Played      int       `json:"played"`
	Win         int       `json:"win"`
	Draw        int       `json:"draw"`
	Lose        int       `json:"lose"`
	UpdateAt    time.Time `json:"update_at"`
	gorm.Model  `json:"-"`
}

func (e *Standing) GetTitle() string {
	return e.TeamName + " " + e.GetPoints()
}

func (e *Standing) GetTitleQueryEscape() string {
	return url.QueryEscape(e.GetTitle())
}

func (e *Standing) GetPoints() string {
	return "(" + strconv.Itoa(e.Points) + "Pts)"
}
