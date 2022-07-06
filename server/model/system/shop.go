package system

type Shop struct {
	Id         uint64 `json:"id" gorm:"primaryKey;column:shop_id"`
	ComId      uint64 `json:"com_id" gorm:"column:com_id"`
	Name       string `json:"name" gorm:"column:shop_name"`
	Content    string `json:"content" gorm:"column:shop_content"`
	Address    string `json:"address" gorm:"column:shop_address"`
	Phone      string `json:"phone" gorm:"column:shop_phone"`
	CreateTime string `json:"create_time" gorm:"column:shop_create_time"`
}

func (Shop) TableName() string {
	return "shop"
}
