package entity

type Team struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	Logo string `json:"logo"`
}

func (e *Team) GetId() int64 {
	return e.ID
}

func (e *Team) GetName() string {
	return e.Name
}

func (e *Team) GetSlug() string {
	return e.Slug
}

func (e *Team) GetCode() string {
	return e.Code
}

func (e *Team) GetLogo() string {
	return e.Logo
}
