package entity

type Statistic struct {
	ID        int64   `gorm:"primaryKey" json:"id"`
	PrimaryID int64   `json:"primary_id"`
	LeagueID  int64   `json:"league_id"`
	League    *League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"league,omitempty"`
}
