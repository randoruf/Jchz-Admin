package service

import (
	"gorm.io/gorm"
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
)

type ArticleService struct{}

func (u *ArticleService) QueryArticlesList(QueryParam request.ArticleQueryRequest) ([]*system.Article, []*system.ArticleTags, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var ArticleList []*system.Article
	var total int64
	err := global.JA_DB.Transaction(func(tx *gorm.DB) error {
		err := global.JA_DB.Table("tiezi").Where("tiezi_title LIKE ?", "%"+query+"%").Count(&total).Error
		if err != nil {
			return err
		}
		err = global.JA_DB.Where("tiezi_title LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&ArticleList).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, nil, 0, err
	}
	var TagNames []*system.ArticleTags
	err = global.JA_DB.Table((&system.Article{}).ArticleTagsViewName()).Find(&TagNames).Error
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
	return newArticle, nil
}

func (u *ArticleService) DeleteArticle(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.Article{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
