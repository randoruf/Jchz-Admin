package response

type LoginData struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type LoginResponse struct {
	Data *LoginData `json:"data"`
	Meta *Meta      `json:"meta"`
}
