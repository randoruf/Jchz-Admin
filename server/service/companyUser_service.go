package service

import (
	"gorm.io/gorm"
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
)

type CompanyUserService struct{}

func (u *CompanyUserService) QueryUsersList(QueryParam request.UserQueryRequest) ([]*system.CompanyUser, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var ComUserList []*system.CompanyUser
	var total int64

	err := global.JA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.JA_DB.Table((&system.CompanyUser{}).TableName()).Where("com_username LIKE ?", "%"+query+"%").Count(&total).Error
		if err != nil {
			return err
		}
		err = global.JA_DB.Where("com_username LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&ComUserList).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return ComUserList, total, nil
}

func (u *CompanyUserService) AddCompanyUser(form request.AddUserRequest) (*system.CompanyUser, error) {
	newUser := &system.CompanyUser{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Phone:    form.Phone,
	}
	err := global.JA_DB.Select("com_username", "com_password", "com_email", "com_phone").Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *CompanyUserService) UpdateCompanyUser(uid string, form request.UpdateUserRequest) (*system.CompanyUser, error) {
	newUser := &system.CompanyUser{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Phone:    form.Phone,
	}
	err := global.JA_DB.Model(&newUser).Where("com_id = ?", uid).Updates(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *CompanyUserService) DeleteCompanyUser(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.CompanyUser{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
