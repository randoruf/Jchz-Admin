package router

type RouterGroup struct {
	AdminRouter   AdminRouter
	MenusRouter   MenusRouter
	UsersRouter   UsersRouter
	ArticleRouter ArticleRouter
	TagsRouter    TagsRouter
}

var RouterGroupApp = new(RouterGroup)
