package response

import "jchz-admin/model/system"

type ShopPageData struct {
	Total   int64          `json:"total"`
	PageNum int64          `json:"pagenum"`
	Shops   []*system.Shop `json:"shops"`
}

type ShopPageResponse struct {
	Data *ShopPageData `json:"data"`
	Meta *Meta         `json:"meta"`
}

type AddShopResponse struct {
	ShopData *system.Shop `json:"data"`
	Meta     *Meta        `json:"meta"`
}

type UpdateShopResponse struct {
	ShopData *system.Shop `json:"data"`
	Meta     *Meta        `json:"meta"`
}
