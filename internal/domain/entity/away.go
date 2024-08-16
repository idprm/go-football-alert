package entity

type Away struct {
	ID        int64 `json:"id,omitempty"`
	PrimaryID int64 `json:"primary_id"`
	TeamID    int64 `json:"team_id"`
	Team      *Team `json:"team"`
	Goal      int   `json:"goal"`
	IsWinner  bool  `gorm:"type:bool;default:false" json:"is_winner"`
}

func (e *Away) GetId() int64 {
	return e.ID
}

func (e *Away) GetTeamId() int64 {
	return e.TeamID
}

func (e *Away) GetGoal() int {
	return e.Goal
}
