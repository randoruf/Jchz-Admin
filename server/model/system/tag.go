package system

type Tag struct {
	ID      int64  `json:"id" gorm:"primaryKey;column:tag_id"`
	TagName string `json:"tagname" gorm:"column:tag_name"`
}

type TagCountList struct {
	ID       int64 `json:"id" gorm:"column:tag_id"`
	Articles int64 `json:"articles" grom:"column:Articles"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func (t *Tag) TagArticlesViewName() string {
	return "tagarticles"
}
