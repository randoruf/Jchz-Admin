package service

import (
	"gorm.io/gorm"
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
)

type UserService struct{}

func (u *UserService) QueryUsersList(QueryParam request.UserQueryRequest) ([]*system.User, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var UserList []*system.User
	var total int64

	// 使用事务提交多组 sql 语句，主要为了保证性能：只有一次 RTT
	err := global.JA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.JA_DB.Table((&system.User{}).TableName()).Where("user_username LIKE ?", "%"+query+"%").Count(&total).Error
		if err != nil {
			return err
		}
		err = global.JA_DB.Where("user_username LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&UserList).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}
	return UserList, total, nil
}

func (u *UserService) AddUser(form request.AddUserRequest) (*system.User, error) {
	newUser := &system.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Phone:    form.Phone,
	}
	err := global.JA_DB.Select("user_username", "user_password", "user_email", "user_phone").Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserService) UpdateUser(uid string, form request.UpdateUserRequest) (*system.User, error) {
	newUser := &system.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Phone:    form.Phone,
	}
	err := global.JA_DB.Model(&newUser).Where("user_id = ?", uid).Updates(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserService) DeleteUser(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.User{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
