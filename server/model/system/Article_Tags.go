package system

type Article_Tags struct {
	ArticleID uint64 `json:"article_id" gorm:"column:tiezi_id"`
	TagId     uint64 `json:"tag_id" gorm:"column:tag_id"`
}

func (Article_Tags) TableName() string {
	return "tiezi_to_tag"
}
