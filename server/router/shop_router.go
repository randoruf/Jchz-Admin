package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type ShopRouter struct{}

func (s *ShopRouter) InitUsersRouter(Router *gin.RouterGroup) {
	shopRouter := Router.Group("")
	shopApi := v1.ShopApiGroup
	{
		shopRouter.GET("/shop", shopApi.GetShopList)
		shopRouter.POST("/shop", shopApi.AddShop)
		shopRouter.PUT("/shop/:id", shopApi.UpdateShop)
		shopRouter.DELETE("/shop/:id", shopApi.DeleteShop)
	}
}
