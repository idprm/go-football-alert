package entity

type Subscription struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `json:"service"`
	Msisdn    string   `gorm:"size:15;not null" json:"msisdn"`
}
