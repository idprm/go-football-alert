package entity

import "strings"

type Menu struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:45" json:"name"`
	Slug        string `gorm:"size:45" json:"slug"`
	KeyPress    string `gorm:"size:6" json:"key_press"`
	TemplateXML string `gorm:"type:text" json:"template_xml"`
	IsActive    bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
}

func (e *Menu) GetId() int {
	return e.ID
}

func (e *Menu) GetName() string {
	return e.Name
}

func (e *Menu) GetSlug() string {
	return e.Slug
}

func (e *Menu) GetKeyPress() string {
	return e.KeyPress
}

func (e *Menu) GetTemplateXML() string {
	return strings.TrimSpace(e.TemplateXML)
}
