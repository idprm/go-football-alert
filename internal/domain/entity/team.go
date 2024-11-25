package entity

import (
	"net/url"
	"strconv"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

type Team struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	PrimaryID int64  `json:"primary_id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Code      string `json:"code"`
	Logo      string `json:"logo"`
	Founded   int    `json:"founded"`
	Country   string `json:"country"`
	Keyword   string `json:"keyword"`
	IsActive  bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active,omitempty"`
	gorm.Model
}

func (e *Team) GetId() int64 {
	return e.ID
}

func (e *Team) GetIdToString() string {
	return strconv.Itoa(int(e.ID))
}

func (e *Team) GetName() string {
	return e.Name
}

func (e *Team) GetNameWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.GetName())
	return result
}

func (e *Team) GetNameQueryEscape() string {
	return url.QueryEscape(e.GetName())
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

func (e *Team) GetKeyword() string {
	return e.Keyword
}
