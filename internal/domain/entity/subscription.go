package entity

type Subscription struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `json:"service"`
	Msisdn    string   `gorm:"size:15;not null" json:"msisdn"`
}

func (e *Subscription) GetId() int64 {
	return e.ID
}

func (e *Subscription) GetServiceId() int {
	return e.ServiceID
}

func (e *Subscription) GetMsisdn() string {
	return e.Msisdn
}
