package entity

import "time"

type Ussd struct {
	ID        int64     `gorm:"primaryKey" json:"id,omitempty"`
	Msisdn    string    `gorm:"size:15;not null" json:"msisdn"`
	KeyPress  string    `gorm:"size:5;not null" json:"key_press"`
	Action    string    `gorm:"size:45" json:"action,omitempty"`
	Result    string    `gorm:"size:150" json:"result,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *Ussd) GetId() int64 {
	return e.ID
}

func (e *Ussd) GetMsisdn() string {
	return e.Msisdn
}

func (e *Ussd) GetKeyPress() string {
	return e.KeyPress
}

func (e *Ussd) GetAction() string {
	return e.Action
}

func (e *Ussd) GetResult() string {
	return e.Result
}

func (e *Ussd) SetAction(v string) {
	e.Action = v
}

func (e *Ussd) SetResult(v string) {
	e.Result = v
}
