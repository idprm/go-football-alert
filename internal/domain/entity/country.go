package entity

import "gorm.io/gorm"

type Country struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:75;index:uidx_country_name,unique" json:"name"`
	Slug string `gorm:"size:90;index:uidx_country_slug,unique" json:"slug"`
	gorm.Model
}
