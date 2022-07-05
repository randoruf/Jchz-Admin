package request

type ArticleQueryRequest struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesize"`
}

type UpdateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`
}
