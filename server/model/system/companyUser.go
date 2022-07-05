package system

type CompanyUser struct {
	ID         uint64 `json:"id" gorm:"column:com_id"`
	Username   string `json:"username" gorm:"column:com_username"`
	Password   string `json:"password" gorm:"column:com_password"`
	Email      string `json:"email" gorm:"column:com_email"`
	Phone      string `json:"phone" gorm:"column:com_phone"`
	Avatar     string `json:"avatar" gorm:"column:com_toux"`
	CreateTime string `json:"create_time" gorm:"column:com_create_time"`
}

func (CompanyUser) TableName() string {
	return "com"
}
