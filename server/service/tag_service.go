package service

import (
	"jchz-admin/global"
	"jchz-admin/model/request"
	"jchz-admin/model/system"
	"strconv"
)

type TagService struct{}

func (t *TagService) QueryTagsList(QueryParam request.TagQueryRequest) ([]*system.Tag, []*system.TagCountList, int64, error) {
	pageNum := QueryParam.Pagenum - 1
	pageSize := QueryParam.Pagesize
	query := QueryParam.Query
	var TagList []*system.Tag
	var CountList []*system.TagCountList
	var total int64
	err := global.JA_DB.Table((&system.Tag{}).TableName()).Where("tag_name LIKE ?", "%"+query+"%").Count(&total).Error
	if err != nil {
		return nil, nil, 0, err
	}
	err = global.JA_DB.Where("tag_name LIKE ?", "%"+query+"%").Limit(pageSize).Offset(pageNum * pageSize).Find(&TagList).Error
	if err != nil {
		return nil, nil, 0, err
	}
	err = global.JA_DB.Table((&system.Tag{}).TagArticlesViewName()).Find(&CountList).Error
	if err != nil {
		return nil, nil, 0, err
	}
	return TagList, CountList, total, nil
}

func (t *TagService) AddTag(form request.AddTagRequest) (*system.Tag, error) {
	newTag := &system.Tag{
		TagName: form.TagName,
	}
	err := global.JA_DB.Select("tag_name").Create(&newTag).Error
	if err != nil {
		return nil, err
	}
	return newTag, nil
}

func (t *TagService) UpdateTag(uid string, form request.UpdateTagRequest) (*system.Tag, error) {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	newTag := &system.Tag{
		ID:      id,
		TagName: form.TagName,
	}
	err = global.JA_DB.Model(&newTag).Update("tag_name", newTag.TagName).Error
	if err != nil {
		return nil, err
	}
	return newTag, nil
}

func (t *TagService) DeleteTag(uid string) (bool, error) {
	err := global.JA_DB.Delete(&system.Tag{}, uid).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
