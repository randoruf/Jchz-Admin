package response

type LoginData struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginResponse struct {
	Data *LoginData `json:"data""`
	Meta *Meta      `json:"meta"`
}
