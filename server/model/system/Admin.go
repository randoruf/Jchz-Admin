package system

type Admin struct {
	ID         uint64 `json:"id" gorm:"column:admin_id"`
	Username   string `json:"username" gorm:"column:admin_username"`
	Password   string `json:"password" gorm:"column:admin_password"`
	Email      string `json:"email" gorm:"column:admin_email"`
	Phone      string `json:"phone" gorm:"column:admin_phone"`
	Avatar     string `json:"avatar" gorm:"column:admin_toux"`
	CreateTime string `json:"create_time" gorm:"column:admin_create_time"`
}

func (Admin) TableName() string {
	return "admin"
}
