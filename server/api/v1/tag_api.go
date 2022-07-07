package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jchz-admin/core/system"
	"jchz-admin/model/request"
	"jchz-admin/model/response"
	"net/http"
	"strconv"
)

type TagApi struct{}

var TagApiGroup = new(TagApi)

func (t *TagApi) GetTagsList(c *gin.Context) {
	var QueryParam request.TagQueryRequest
	var err error
	QueryParam.Query = c.Query("query")
	pageNumStr, ok := c.GetQuery("pagenum")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagenum, err = strconv.Atoi(pageNumStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	pageSizeStr, ok := c.GetQuery("pagesize")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagesize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	TagsList, CountList, total, err := tagService.QueryTagsList(QueryParam)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "获取标签列表失败", c)
		return
	}

	var ResTagsList []*response.Tags

	for i, j := 0, 0; i < len(TagsList); i++ {
		Tag := &response.Tags{
			ID:       TagsList[i].ID,
			TagName:  TagsList[i].TagName,
			Articles: 0,
		}
		if j < len(CountList) && CountList[j].ID == TagsList[i].ID {
			Tag.Articles = CountList[j].Articles
			j++
		}
		ResTagsList = append(ResTagsList, Tag)
	}

	c.JSON(http.StatusOK, response.TagPageResponse{
		Data: &response.TagPageData{
			Total:   total,
			PageNum: int64(QueryParam.Pagenum),
			Tags:    ResTagsList,
		},
		Meta: &response.Meta{
			Msg:    "获取标签列表成功",
			Status: http.StatusOK,
		},
	})

}

func (t *TagApi) AddTag(c *gin.Context) {
	var addTagQuery request.AddTagRequest
	err := c.ShouldBind(&addTagQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newTag, err := tagService.AddTag(addTagQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "添加标签失败", c)
		return
	}
	c.JSON(http.StatusOK, response.AddTagResponse{
		TagData: newTag,
		Meta: &response.Meta{
			Msg:    "添加标签成功",
			Status: http.StatusOK,
		},
	})
}

func (t *TagApi) UpdateTag(c *gin.Context) {
	var updateTagQuery request.UpdateTagRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateTagQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newTag, err := tagService.UpdateTag(uid, updateTagQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新标签失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateTagResponse{
		TagData: newTag,
		Meta: &response.Meta{
			Msg:    "更新标签成功",
			Status: http.StatusOK,
		},
	})
}

func (t *TagApi) DeleteTag(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := tagService.DeleteTag(uid)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	if flag == false {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除标签失败", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除标签成功", c)
	return
}

func (t *TagApi) TagNameExists(c *gin.Context) {
	name := c.Query("tagName")
	if name == "" {
		c.JSON(http.StatusOK, response.CheckExistsResponse{
			Data: &response.CheckExistsData{Result: "true"},
			Meta: &response.Meta{
				Msg:    "标签为空",
				Status: http.StatusOK,
			},
		})
		return
	}
	flag, err := tagService.TagNameExists(name)
	if err != nil && err != gorm.ErrRecordNotFound {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "检查标签名失败，服务器出错，请稍后再试", c)
		return
	}
	if flag == false {
		c.JSON(http.StatusOK, response.CheckExistsResponse{
			Data: &response.CheckExistsData{Result: "false"},
			Meta: &response.Meta{
				Msg:    "该标签不存在",
				Status: http.StatusOK,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.CheckExistsResponse{
		Data: &response.CheckExistsData{Result: "true"},
		Meta: &response.Meta{
			Msg:    "该标签存在",
			Status: http.StatusOK,
		},
	})
	return
}
