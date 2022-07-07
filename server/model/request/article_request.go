package request

type ArticleQueryRequest struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesize"`
}

type UpdateArticleRequest struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
	Tag1    string `json:"tag1"`
	Tag2    string `json:"tag2"`
	Tag3    string `json:"tag3"`
	Tags    string `json:"tags"` // 修改前 tag
}
