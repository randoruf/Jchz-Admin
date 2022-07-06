package Initialize

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/middleware"
	"jchz-admin/router"
)

// InitRouters 初始化服务端路由
func InitRouters() *gin.Engine {
	r := gin.Default()
	PublicGroup := r.Group("/api/v1/")
	{
		router.RouterGroupApp.AdminRouter.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := r.Group("/api/v1/")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.RouterGroupApp.MenusRouter.InitMenusRouter(PrivateGroup)
		router.RouterGroupApp.UsersRouter.InitUsersRouter(PrivateGroup)
		router.RouterGroupApp.ArticleRouter.InitArticleRouter(PrivateGroup)
		router.RouterGroupApp.TagsRouter.InitTagsRouter(PrivateGroup)
		router.RouterGroupApp.ShopRouter.InitUsersRouter(PrivateGroup)
	}
	return r
}
