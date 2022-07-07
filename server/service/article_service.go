package service

import (
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
	"strings"
)

type ArticleService struct{}

func (u *ArticleService) QueryArticlesList(QueryParam request.ArticleQueryRequest) ([]*system.Article, []*system.ArticleTags, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var ArticleList []*system.Article
	var TagNames []*system.ArticleTags
	var total int64
	err := global.JA_DB.Table("tiezi").Where("tiezi_title LIKE ?", "%"+query+"%").Count(&total).Error
	if err != nil {
		return nil, nil, 0, err
	}
	err = global.JA_DB.Where("tiezi_title LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&ArticleList).Error
	if err != nil {
		return nil, nil, 0, err
	}
	if len(ArticleList) == 0 {
		return nil, nil, 0, nil
	}
	err = global.JA_DB.Table((&system.Article{}).ArticleTagsViewName()).Where("tiezi_id >= ?", ArticleList[0].ArticleId).Find(&TagNames).Error
	if err != nil {
		return nil, nil, 0, err
	}
	return ArticleList, TagNames, total, nil
}

func (u *ArticleService) UpdateArticle(uid string, form request.UpdateArticleRequest) (*system.Article, error) {
	newArticle := &system.Article{
		Title:   form.Title,
		Content: form.Content,
		Cover:   form.Cover,
	}
	err := global.JA_DB.Model(&newArticle).Where("tiezi_id = ?", uid).Updates(&newArticle).Error
	if err != nil {
		return nil, err
	}
	Tags := strings.Split(form.Tags, "，")
	var Old system.Tag
	var New system.Tag

	// 将更新文章对应标签封装成函数以减少重复代码
	updateTag := func(oldTag, newTag string) error {
		if newTag != "" {
			err = global.JA_DB.Table((&system.Tag{}).TableName()).Where("tag_name = ?", newTag).First(&New).Error
			if err != nil {
				return err
			}
		}
		if oldTag == "" {
			err = global.JA_DB.Table("tiezi_to_tag").Create([]map[string]interface{}{{"tag_id": New.ID, "tiezi_id": form.Id}}).Error
			if err != nil {
				return err
			}
		} else {
			err = global.JA_DB.Table((&system.Tag{}).TableName()).Where("tag_name = ?", oldTag).First(&Old).Error
			if err != nil {
				return err
			}
			if newTag != "" {
				err = global.JA_DB.Model(&system.Article_Tags{}).Where("tiezi_id = ? AND tag_id = ?", form.Id, Old.ID).Update("tag_id", New.ID).Error
				if err != nil {
					return err
				}
			} else {
				err = global.JA_DB.Where("tiezi_id = ? AND tag_id = ?", form.Id, Old.ID).Delete(&system.Article_Tags{}).Error
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	if (len(Tags) > 0 && form.Tag1 != Tags[0]) || (len(Tags) < 1 && form.Tag1 != "") {
		if len(Tags) < 1 {
			err = updateTag("", form.Tag1)
		} else {
			err = updateTag(Tags[0], form.Tag1)
		}
		if err != nil {
			return nil, err
		}
	}
	if (len(Tags) > 1 && form.Tag2 != Tags[1]) || (len(Tags) < 2 && form.Tag2 != "") {
		if len(Tags) < 2 {
			err = updateTag("", form.Tag2)
		} else {
			err = updateTag(Tags[1], form.Tag2)
		}
		if err != nil {
			return nil, err
		}
	}
	if (len(Tags) > 2 && form.Tag3 != Tags[2]) || (len(Tags) < 3 && form.Tag3 != "") {
		if len(Tags) < 3 {
			err = updateTag("", form.Tag3)
		} else {
			err = updateTag(Tags[2], form.Tag3)
		}
		if err != nil {
			return nil, err
		}
	}

	return newArticle, nil
}

func (u *ArticleService) DeleteArticle(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.Article{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
