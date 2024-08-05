package entity

type Fixture struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Referee   string `json:"referee"`
	Timezone  string `json:"timezone"`
	Date      string `json:"date"`
	TimeStamp int    `json:"timestamp"`
	HomeID    int64  `json:"home_id"`
	Home      *Home  `json:"home"`
	AwayID    int64  `json:"away_id"`
	Away      *Away  `json:"away"`
}
