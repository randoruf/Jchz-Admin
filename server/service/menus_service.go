package service

import (
	"fmt"
	"jchz-admin/global"
	"jchz-admin/model/system"
)

type MenuService struct{}

func (m *MenuService) GetMenuList() ([]system.Menu, error) {
	if global.JA_DB == nil {
		return nil, fmt.Errorf("DB does not init")
	}
	var menus []system.Menu
	err := global.JA_DB.Order("order_id ASC").Find(&menus).Error
	if err != nil {
		return nil, err
	}
	return menus, nil
}
