package entity

type Subscription struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	CountryID int      `json:"country_id"`
	Country   *Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
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
