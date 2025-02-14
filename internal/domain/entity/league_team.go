package entity

type LeagueTeam struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	LeagueID int64   `json:"league_id"`
	League   *League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
	TeamID   int64   `json:"team_id"`
	Team     *Team   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	IsActive bool    `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
}
