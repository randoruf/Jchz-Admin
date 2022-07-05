package system

type Menu struct {
	Id       int    `json:"id" gorm:"column:id"`
	AuthName string `json:"authname" gorm:"column:authname"`
	Path     string `json:"path" gorm:"column:path"`
	ParentID int    `json:"parent_id" gorm:"column:parent_id"`
	Order    int    `json:"order_id" gorm:"column:order_id"`
}

func (Menu) TableName() string {
	return "back_sys_menu"
}
