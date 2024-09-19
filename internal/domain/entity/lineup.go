package entity

type Lineup struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	PrimaryID int64 `json:"primary_id"`
}
