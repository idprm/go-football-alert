package entity

type Ussd struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	Msisdn string `json:"msisdn"`
}
