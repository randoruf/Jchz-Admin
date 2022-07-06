package service

import (
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
	"strconv"
)

type ShopService struct{}

func (s *ShopService) QueryShopsList(QueryParam request.ShopQueryRequest) ([]*system.Shop, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var ShopList []*system.Shop
	var total int64

	err := global.JA_DB.Table((&system.Shop{}).TableName()).Where("shop_name LIKE ?", "%"+query+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = global.JA_DB.Where("shop_name LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&ShopList).Error
	if err != nil {
		return nil, 0, err
	}
	return ShopList, total, nil
}

func (s *ShopService) AddShop(form request.AddShopRequest) (*system.Shop, error) {
	var comid uint64
	var err error
	switch form.ComId.(type) {
	case float64:
		comid = uint64(form.ComId.(float64))
		break
	default:
		comid, err = strconv.ParseUint(form.ComId.(string), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	newShop := &system.Shop{
		ComId:   comid,
		Name:    form.Name,
		Content: form.Content,
		Address: form.Address,
		Phone:   form.Phone,
	}
	err = global.JA_DB.Select("com_id", "shop_name", "shop_content", "shop_address", "shop_phone").Create(&newShop).Error
	if err != nil {
		return nil, err
	}
	return newShop, nil
}

func (s *ShopService) UpdateShop(uid string, form request.UpdateShopRequest) (*system.Shop, error) {
	var comid uint64
	var err error
	switch form.ComId.(type) {
	case float64:
		comid = uint64(form.ComId.(float64))
		break
	default:
		comid, err = strconv.ParseUint(form.ComId.(string), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	newShop := &system.Shop{
		ComId:   comid,
		Name:    form.Name,
		Content: form.Content,
		Address: form.Address,
		Phone:   form.Phone,
	}
	err = global.JA_DB.Model(&newShop).Where("shop_id = ?", uid).Updates(&newShop).Error
	if err != nil {
		return nil, err
	}
	return newShop, nil
}

func (s *ShopService) DeleteShop(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.Shop{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
