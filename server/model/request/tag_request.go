package request

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
