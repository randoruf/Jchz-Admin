package v1

import "jchz-admin/service"

var (
	adminService   = service.ServiceGroupApp.AdminService
	menuService    = service.ServiceGroupApp.MenuService
	userService    = service.ServiceGroupApp.UserService
	companyService = service.ServiceGroupApp.CompanyUserService
	articleService = service.ServiceGroupApp.ArticleService
	tagService     = service.ServiceGroupApp.TagService
)
