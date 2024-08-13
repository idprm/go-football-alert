package entity

type Season struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:35;not null" json:"name"`
	Slug string `gorm:"size:35;not null" json:"slug"`
}

func (e *Season) GetId() int {
	return e.ID
}

func (e *Season) GetName() string {
	return e.Name
}

func (e *Season) GetSlug() string {
	return e.Slug
}
