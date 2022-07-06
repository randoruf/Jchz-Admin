package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
	"jchz-admin/utils"
)

type AdminService struct{}

// Login 后台页面登录逻辑处理
func (AdminService *AdminService) Login(u *system.Admin) (adminInter *system.Admin, err error) {
	if global.JA_DB == nil {
		return nil, fmt.Errorf("DB does not init")
	}
	var admin system.Admin
	err = global.JA_DB.Where("admin_username = ?", u.Username).First(&admin).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, admin.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &admin, err
}

func (AdminService *AdminService) QueryUsersList(QueryParam request.UserQueryRequest) ([]*system.Admin, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var AdminList []*system.Admin
	var total int64

	err := global.JA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.JA_DB.Table((&system.Admin{}).TableName()).Where("admin_username LIKE ?", "%"+query+"%").Count(&total).Error
		if err != nil {
			return err
		}
		err = global.JA_DB.Where("admin_username LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&AdminList).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return AdminList, total, nil
}

func (AdminService *AdminService) AddAdmin(form request.AddUserRequest) (*system.Admin, error) {
	newUser := &system.Admin{
		Username: form.Username,
		Password: utils.BcryptHash(form.Password),
		Email:    form.Email,
		Phone:    form.Phone,
	}
	err := global.JA_DB.Select("admin_username", "admin_password", "admin_email", "admin_phone").Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (AdminService *AdminService) UpdateAdmin(uid string, form request.UpdateUserRequest) (*system.Admin, error) {
	newUser := &system.Admin{
		Username: form.Username,
		Email:    form.Email,
		Phone:    form.Phone,
	}
	if form.Password != "" {
		newUser.Password = utils.BcryptHash(form.Password)
	}
	err := global.JA_DB.Model(&newUser).Where("admin_id = ?", uid).Updates(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (AdminService *AdminService) DeleteAdmin(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.Admin{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
