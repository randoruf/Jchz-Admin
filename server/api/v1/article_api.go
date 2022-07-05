package v1

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/core/system"
	"jchz-admin/model/request"
	"jchz-admin/model/response"
	"net/http"
	"strconv"
	"strings"
)

type ArticleApi struct{}

var ArticleApiGroup = new(ArticleApi)

func (u *ArticleApi) GetArticlesList(c *gin.Context) {
	var QueryParam request.ArticleQueryRequest
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

	ArticlesList, ArticleTags, total, err := articleService.QueryArticlesList(QueryParam)

	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "获取文章列表失败", c)
		return
	}
	var ArticlesDataList []*response.ArticleList
	for i, j := 0, 0; i < len(ArticlesList); i++ {
		temp := &response.ArticleList{
			ArticleId:  ArticlesList[i].ArticleId,
			Title:      ArticlesList[i].Title,
			Content:    ArticlesList[i].Content,
			Cover:      ArticlesList[i].Cover,
			UserID:     ArticlesList[i].UserID,
			CreateTime: ArticlesList[i].CreateTime,
		}
		var TagNames []string
		for j < len(ArticleTags) {
			if ArticleTags[j].ArticleID == ArticlesList[i].ArticleId {
				TagNames = append(TagNames, ArticleTags[j].TagName)
				j++
			} else {
				break
			}
		}
		TagNamesStr := strings.Join(TagNames, "，")
		temp.TagName = TagNamesStr
		ArticlesDataList = append(ArticlesDataList, temp)
	}
	c.JSON(http.StatusOK, response.ArticlePageResponse{
		Data: &response.ArticlePageData{
			Total:    total,
			PageNum:  int64(QueryParam.Pagenum),
			Articles: ArticlesDataList,
		},
		Meta: &response.Meta{
			Msg:    "获取文章列表成功",
			Status: http.StatusOK,
		},
	})
}

func (u *ArticleApi) UpdateArticle(c *gin.Context) {
	var updateArticleQuery request.UpdateArticleRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateArticleQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newArticle, err := articleService.UpdateArticle(uid, updateArticleQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新文章失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateArticleResponse{
		ArticleData: newArticle,
		Meta: &response.Meta{
			Msg:    "更新文章成功",
			Status: http.StatusOK,
		},
	})
}

func (u *ArticleApi) DeleteArticle(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := articleService.DeleteArticle(uid)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	if flag == false {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除文章失败", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除文章成功", c)
	return
}
