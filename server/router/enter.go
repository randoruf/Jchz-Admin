package router

type RouterGroup struct {
	AdminRouter   AdminRouter
	MenusRouter   MenusRouter
	UsersRouter   UsersRouter
	ArticleRouter ArticleRouter
	TagsRouter    TagsRouter
	ShopRouter    ShopRouter
}

var RouterGroupApp = new(RouterGroup)
