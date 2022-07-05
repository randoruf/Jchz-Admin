package response

type MenusData struct {
	Id       int          `json:"id"`
	AuthName string       `json:"authName"`
	Path     string       `json:"path"`
	Children []*MenusData `json:"children"`
	Order    interface{}  `json:"order"`
}

type MenusResponse struct {
	Data []*MenusData `json:"data"`
	Meta *Meta        `json:"meta"`
}
