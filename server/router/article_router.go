package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articlesRouter := Router.Group("")
	articleApi := v1.ArticleApiGroup
	{
		articlesRouter.GET("/article", articleApi.GetArticlesList)
		articlesRouter.PUT("/article/:id", articleApi.UpdateArticle)
		articlesRouter.DELETE("/article/:id", articleApi.DeleteArticle)
	}
}
