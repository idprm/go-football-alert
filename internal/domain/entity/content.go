package entity

type Content struct {
	ID        int      `gorm:"primaryKey" json:"id"`
	ServiceID int      `json:"service_id"`
	Service   *Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"service,omitempty"`
	Key       string   `gorm:"size:50" json:"key"`
	Value     string   `gorm:"size:150" json:"value"`
}
