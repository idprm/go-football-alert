package entity

type Service struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	CountryID int      `json:"country_id"`
	Country   *Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"country,omitempty"`
	Name      string   `gorm:"size:50;not null" json:"name"`
	Code      string   `gorm:"size:15;not null" json:"code"`
	UrlTelco  string   `gorm:"size:350;not null" json:"url_telco"`
}

func (e *Service) GetId() int {
	return e.ID
}

func (e *Service) GetName() string {
	return e.Name
}

func (e *Service) GetCode() string {
	return e.Code
}

func (e *Service) GetUrlTelco() string {
	return e.UrlTelco
}
