package entity

type Menu struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:45" json:"name"`
	KeyPress string `gorm:"size:6" json:"key_press"`
	ParentID int    `gorm:"size:5" json:"parent_id"`
	Child    int    `gorm:"size:5" json:"child"`
	Action   string `gorm:"size:45" json:"action"`
	IsActive bool   `gorm:"type:boolean;default:false;column:is_active" json:"is_active"`
}

func (e *Menu) GetId() int {
	return e.ID
}

func (e *Menu) GetName() string {
	return e.Name
}

func (e *Menu) GetKeyPress() string {
	return e.KeyPress
}

func (e *Menu) GetParentId() int {
	return e.ParentID
}

func (e *Menu) GetChild() int {
	return e.Child
}

func (e *Menu) GetAction() string {
	return e.Action
}
