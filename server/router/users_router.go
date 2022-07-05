package router

import (
	"github.com/gin-gonic/gin"
	v1 "jchz-admin/api/v1"
)

type UsersRouter struct{}

func (u *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	usersRouter := Router.Group("")
	userApi := v1.UserApiGroup
	{
		usersRouter.GET("/users", userApi.GetUserList)
		usersRouter.POST("/users", userApi.AddUser)
		usersRouter.PUT("/users/:id", userApi.UpdateUser)
		usersRouter.DELETE("/users/:id", userApi.DeleteUser)
		usersRouter.POST("/companyUsers", userApi.AddCompanyUser)
		usersRouter.PUT("/companyUsers/:id", userApi.UpdateCompanyUser)
		usersRouter.DELETE("/companyUsers/:id", userApi.DeleteCompanyUser)
		usersRouter.POST("/admin", userApi.AddAdmin)
		usersRouter.PUT("/admin/:id", userApi.UpdateAdmin)
		usersRouter.DELETE("/admin/:id", userApi.DeleteAdmin)
	}
}
