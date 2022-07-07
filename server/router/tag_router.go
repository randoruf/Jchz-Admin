package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type TagsRouter struct{}

func (t *TagsRouter) InitTagsRouter(Router *gin.RouterGroup) {
	tagsRouter := Router.Group("")
	tagApi := v1.TagApiGroup
	{
		tagsRouter.GET("/tags", tagApi.GetTagsList)
		tagsRouter.POST("/tags", tagApi.AddTag)
		tagsRouter.PUT("/tags/:id", tagApi.UpdateTag)
		tagsRouter.DELETE("/tags/:id", tagApi.DeleteTag)
		tagsRouter.GET("/tags/tagNameExists", tagApi.TagNameExists)
	}
}
