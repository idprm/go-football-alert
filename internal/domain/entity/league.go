package entity

type League struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	PrimaryID int64  `json:"primary_id"`
	Name      string `gorm:"size:110" json:"name"`
	Slug      string `gorm:"size:110" json:"slug"`
	Logo      string `gorm:"size:110" json:"logo"`
	Country   string `gorm:"size:110" json:"country"`
	IsActive  bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
}

func (e *League) GetId() int64 {
	return e.ID
}

func (e *League) GetName() string {
	return e.Name
}

func (e *League) GetSlug() string {
	return e.Slug
}

func (e *League) GetLogo() string {
	return e.Logo
}

func (e *League) GetCountry() string {
	return e.Country
}
