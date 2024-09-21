package entity

import (
	"time"

	"gorm.io/gorm"
)

type Standing struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Rank        int       `json:"primary_id"`
	LeagueID    int64     `json:"league_id"`
	League      *League   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	TeamID      int64     `json:"team_id"`
	Team        *Team     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
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

// Rank int `json:"rank"`
// Team struct {
// 	PrimaryID   int    `json:"id"`
// 	Name        string `json:"name"`
// 	Points      int    `json:"points"`
// 	GoalsDiff   int    `json:"goalsDiff"`
// 	Group       string `json:"group"`
// 	Form        string `json:"form"`
// 	Status      string `json:"status"`
// 	Description string `json:"description"`
// } `json:"team"`
// All struct {
// 	Played int `json:"played"`
// 	Win    int `json:"win"`
// 	Draw   int `json:"draw"`
// 	Lose   int `json:"lose"`
// } `json:"all"`
// UpdateAt string `json:"update"`
