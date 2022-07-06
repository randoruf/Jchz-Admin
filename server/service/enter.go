package service

type ServiceGroup struct {
	AdminService       AdminService
	MenuService        MenuService
	UserService        UserService
	CompanyUserService CompanyUserService
	ArticleService     ArticleService
	TagService         TagService
	ShopService        ShopService
}

var ServiceGroupApp = new(ServiceGroup)
