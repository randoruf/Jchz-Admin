package system

type User struct {
	ID         uint64 `json:"id" gorm:"primaryKey;column:user_id"`
	Username   string `json:"username" gorm:"column:user_username"`
	Password   string `json:"password" gorm:"column:user_password"`
	Sex        string `json:"sex" gorm:"column:user_sex"`
	Birth      string `json:"birth" gorm:"column:user_birth"`
	Hometown   string `json:"hometown" gorm:"column:user_hometown"`
	Email      string `json:"email" gorm:"column:user_email"`
	Phone      string `json:"phone" gorm:"column:user_phone"`
	Avatar     string `json:"avatar" gorm:"column:user_toux"`
	CreateTime string `json:"create_time" gorm:"column:user_create_time"`
}

func (u *User) TableName() string {
	return "user"
}
