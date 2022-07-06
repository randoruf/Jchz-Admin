package request

type ShopQueryRequest struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesize"`
}

type AddShopRequest struct {
	Address string      `json:"address"`
	ComId   interface{} `json:"com_id"`
	Content string      `json:"content"`
	Name    string      `json:"name"`
	Phone   string      `json:"phone"`
}

type UpdateShopRequest struct {
	Address string      `json:"address"`
	ComId   interface{} `json:"com_id"`
	Content string      `json:"content"`
	Name    string      `json:"name"`
	Phone   string      `json:"phone"`
}
