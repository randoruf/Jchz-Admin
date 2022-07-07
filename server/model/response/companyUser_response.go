package response

import (
	"jchz-admin/model/system"
)

type CompanyUserData struct {
	Total   int64                 `json:"total"`
	PageNum int64                 `json:"pagenum"`
	Users   []*system.CompanyUser `json:"users"`
}

type CompanyUserResponse struct {
	Data *CompanyUserData `json:"data"`
	Meta *Meta            `json:"meta"`
}

type AddCompanyUserResponse struct {
	UserData *system.CompanyUser `json:"data"`
	Meta     *Meta               `json:"meta"`
}

type UpdateCompanyUserResponse struct {
	UserData *system.CompanyUser `json:"data"`
	Meta     *Meta               `json:"meta"`
}
