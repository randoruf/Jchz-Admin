package service

type ServiceGroup struct {
	AdminService       AdminService
	MenuService        MenuService
	UserService        UserService
	CompanyUserService CompanyUserService
	ArticleService     ArticleService
	TagService         TagService
}

var ServiceGroupApp = new(ServiceGroup)
