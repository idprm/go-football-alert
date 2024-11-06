package entity

import (
	"net/url"
	"strconv"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type League struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	PrimaryID int64  `json:"primary_id"`
	Name      string `gorm:"size:110" json:"name"`
	Slug      string `gorm:"size:110" json:"slug"`
	Code      string `gorm:"size:50" json:"code"`
	Logo      string `gorm:"size:110" json:"logo"`
	Country   string `gorm:"size:110" json:"country"`
	IsActive  bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
	Keyword   string `json:"keyword"`
}

func (e *League) GetId() int64 {
	return e.ID
}

func (e *League) GetIdToString() string {
	return strconv.Itoa(int(e.ID))
}

func (e *League) GetName() string {
	return e.Name
}

func (e *League) GetNameQueryEscape() string {
	return url.QueryEscape(e.GetNameWithoutAccents())
}

func (e *League) GetNameWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, e.GetName())
	return result
}

func (e *League) GetSlug() string {
	return e.Slug
}

func (e *League) GetLogo() string {
	return e.Logo
}

func (e *League) GetCode() string {
	return e.Code
}

func (e *League) GetCountry() string {
	return e.Country
}

func (e *League) GetKeyword() string {
	return e.Keyword
}
