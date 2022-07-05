package response

import "jchz-admin/model/system"

type AdminData struct {
	Total   int64           `json:"total"`
	PageNum int64           `json:"pagenum"`
	Users   []*system.Admin `json:"users"`
}

type AdminResponse struct {
	Data *AdminData `json:"data"`
	Meta *Meta      `json:"meta"`
}

type AddAdminResponse struct {
	UserData *system.Admin `json:"data"`
	Meta     *Meta         `json:"meta"`
}

type UpdateAdminResponse struct {
	UserData *system.Admin `json:"data"`
	Meta     *Meta         `json:"meta"`
}
