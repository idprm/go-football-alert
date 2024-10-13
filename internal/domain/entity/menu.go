package entity

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Menu struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Category    string `gorm:"size:50;not null" json:"category"`
	Name        string `gorm:"size:45" json:"name"`
	Slug        string `gorm:"size:45" json:"slug"`
	TemplateXML string `gorm:"type:text" json:"template_xml"`
	IsConfirm   bool   `gorm:"type:boolean;default:false;column:is_confirm" json:"is_confirm"`
	IsActive    bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
}

func (e *Menu) GetId() int {
	return e.ID
}

func (e *Menu) GetCategory() string {
	return strings.ToLower(e.Category)
}

func (e *Menu) GetName() string {
	return e.Name
}

func (e *Menu) GetSlug() string {
	return e.Slug
}

func (e *Menu) GetTemplateXML() string {
	return strings.TrimSpace(e.GetTemplateWithoutAccents())
}

func (e *Menu) GetTemplateWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.TemplateXML)
	return result
}
