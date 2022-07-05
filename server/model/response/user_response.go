package response

import "jchz-admin/model/system"

type UserPageData struct {
	Total   int64          `json:"total"`
	PageNum int64          `json:"pagenum"`
	Users   []*system.User `json:"users"`
}

type UserPageResponse struct {
	Data *UserPageData `json:"data"`
	Meta *Meta         `json:"meta"`
}

type AddUserResponse struct {
	UserData *system.User `json:"data"`
	Meta     *Meta        `json:"meta"`
}

type UpdateUserResponse struct {
	UserData *system.User `json:"data"`
	Meta     *Meta        `json:"meta"`
}
