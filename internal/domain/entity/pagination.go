package entity

import "time"

type Pagination struct {
	Limit      int         `json:"limit" query:"limit"`
	Page       int         `json:"page" query:"page"`
	Sort       string      `json:"sort" query:"sort"`
	Search     string      `json:"search" query:"search"`
	Date       string      `json:"date" query:"date"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func (p *Pagination) GetSearch() string {
	return p.Search
}

func (p *Pagination) GetDate() time.Time {
	date, _ := time.Parse("2006-01-02", p.Date)
	return date
}

func (p *Pagination) IsDate() bool {
	return p.Date != ""
}

type PaginateUssd struct {
	Limit int         `json:"limit,omitempty" query:"limit"`
	Page  int         `json:"page,omitempty" query:"page"`
	Sort  string      `json:"sort,omitempty" query:"sort"`
	Rows  interface{} `json:"rows"`
}

func (p *PaginateUssd) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PaginateUssd) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func (p *PaginateUssd) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginateUssd) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}
