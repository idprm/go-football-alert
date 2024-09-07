package entity

type Content struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Name      string   `gorm:"size:50" json:"name"`
	Value     string   `gorm:"size:250" json:"value"`
}

func (e *Content) GetId() int {
	return e.ID
}

func (e *Content) GetServiceId() int {
	return e.ServiceID
}

func (e *Content) GetName() string {
	return e.Name
}

func (e *Content) GetValue() string {
	return e.Value
}
