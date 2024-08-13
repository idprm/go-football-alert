package entity

type League struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	PrimaryID int64  `json:"primary_id"`
	Name      string `gorm:"size:100" json:"name"`
	Slug      string `gorm:"size:110" json:"slug"`
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
