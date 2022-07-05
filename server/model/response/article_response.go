package response

import "jchz-admin/model/system"

type ArticleList struct {
	ArticleId  int    `json:"id" gorm:"primaryKey;column:tiezi_id"`
	Title      string `json:"title" gorm:"column:tiezi_title"`
	Content    string `json:"content" gorm:"column:tiezi_content"`
	Cover      string `json:"cover" gorm:"column:tiezi_page"`
	UserID     string `json:"userid" gorm:"column:user_id"`
	CreateTime string `json:"create_time" gorm:"column:tiezi_create_time"`
	TagName    string `json:"tags"`
}

type ArticlePageData struct {
	Total    int64          `json:"total"`
	PageNum  int64          `json:"pagenum"`
	Articles []*ArticleList `json:"articles"`
}

type ArticlePageResponse struct {
	Data *ArticlePageData `json:"data"`
	Meta *Meta            `json:"meta"`
}

type UpdateArticleResponse struct {
	ArticleData *system.Article `json:"data"`
	Meta        *Meta           `json:"meta"`
}
