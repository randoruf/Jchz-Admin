package request

// user, companyUser, admin 通用请求结构

type UserQueryRequest struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesize"`
	Role     string `json:"role"`
}

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
