package entity

type Home struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	PrimaryID int64 `json:"primary_id"`
	TeamID    int64 `json:"team_id"`
	Team      *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`
	Goal      int   `gorm:"size:10" json:"goal"`
	IsWinner  bool  `gorm:"type:boolean;default:false" json:"is_winner"`
}

func (e *Home) GetId() int64 {
	return e.ID
}

func (e *Home) GetTeamId() int64 {
	return e.TeamID
}

func (e *Home) GetGoal() int {
	return e.Goal
}
