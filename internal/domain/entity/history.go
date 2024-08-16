package entity

type History struct {
	ID        int64    `gorm:"primaryKey" json:"id"`
	CountryID int      `json:"country_id"`
	Country   *Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	ServiceID int      `json:"service_id"`
	Service   *Service `json:"service"`
	Msisdn    string   `gorm:"size:15;not null" json:"msisdn"`
}

func (e *History) GetId() int64 {
	return e.ID
}

func (e *History) GetCountryId() int {
	return e.CountryID
}

func (e *History) GetServiceId() int {
	return e.ServiceID
}

func (e *History) GetMsisdn() string {
	return e.Msisdn
}
