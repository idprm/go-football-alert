package entity

type Season struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:35;not null" json:"name"`
	Slug string `gorm:"size:35;not null" json:"slug"`
}
