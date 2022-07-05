package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type MenusRouter struct{}

func (m *MenusRouter) InitMenusRouter(Router *gin.RouterGroup) {
	menusRouter := Router.Group("")
	menuApi := v1.MenuApiGroup
	{
		menusRouter.GET("/menus", menuApi.MenuList)
	}
}
