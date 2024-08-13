package entity

type Content struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Key       string   `gorm:"size:50" json:"key"`
	Value     string   `gorm:"size:150" json:"value"`
}

func (e *Content) GetId() int {
	return e.ID
}

func (e *Content) GetServiceId() int {
	return e.ServiceID
}

func (e *Content) GetKey() string {
	return e.Key
}

func (e *Content) GetValue() string {
	return e.Value
}
