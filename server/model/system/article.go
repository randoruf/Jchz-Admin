package system

type ArticleTags struct {
	ArticleID int    `json:"id" gorm:"column:tiezi_id"`
	TagName   string `json:"tags" gorm:"column:tag_name"`
}

type Article struct {
	ArticleId  int    `json:"id" gorm:"primaryKey;column:tiezi_id"`
	Title      string `json:"title" gorm:"column:tiezi_title"`
	Content    string `json:"content" gorm:"column:tiezi_content"`
	Cover      string `json:"cover" gorm:"column:tiezi_page"`
	UserID     string `json:"userid" gorm:"column:user_id"`
	CreateTime string `json:"create_time" gorm:"column:tiezi_create_time"`
}

func (Article) TableName() string {
	return "tiezi"
}

func (Article) ArticleTagsViewName() string {
	return "articletags"
}
