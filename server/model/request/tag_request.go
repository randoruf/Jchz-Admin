package request

// user, companyUser, admin 通用请求结构

type TagQueryRequest struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesize"`
}

type AddTagRequest struct {
	TagName string `json:"tagname"`
}

type UpdateTagRequest struct {
	TagName string `json:"tagname"`
}
