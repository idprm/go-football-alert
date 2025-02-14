package entity

type Verify struct {
	Msisdn   string `json:"msisdn"`
	Pin      string `json:"pin"`
	Status   string `json:"status"`
	Category string `json:"category"`
}

func (e *Verify) GetMsisdn() string {
	return e.Msisdn
}

func (e *Verify) GetPin() string {
	return e.Pin
}

func (e *Verify) GetStatus() string {
	return e.Status
}

func (e *Verify) GetCategory() string {
	return e.Category
}

func (e *Verify) SetStatus(data string) {
	e.Status = data
}

func (e *Verify) IsValid() bool {
	return e.Status == "PONG"
}
