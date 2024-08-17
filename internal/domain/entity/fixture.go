package entity

type Fixture struct {
	ID        int64   `gorm:"primaryKey" json:"id"`
	PrimaryID int64   `json:"primary_id"`
	Timezone  string  `json:"timezone"`
	Date      string  `json:"date"`
	TimeStamp int     `json:"timestamp"`
	LeagueID  int64   `json:"league_id"`
	League    *League `json:"league"`
	HomeID    int64   `json:"home_id"`
	Home      *Home   `json:"home"`
	AwayID    int64   `json:"away_id"`
	Away      *Away   `json:"away"`
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

func (e *Fixture) GetDate() string {
	return e.Date
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
