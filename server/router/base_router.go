package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type AdminRouter struct{}

func (a *AdminRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("")
	baseApi := v1.BaseApiGroup
	{
		baseRouter.POST("/login", baseApi.Login)
	}
}
