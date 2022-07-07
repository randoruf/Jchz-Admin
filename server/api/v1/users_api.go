package v1

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/core/system"
	"jchz-admin/model/request"
	"jchz-admin/model/response"
	"net/http"
	"strconv"
)

type UserApi struct{}

var UserApiGroup = new(UserApi)

func (u *UserApi) GetUserList(c *gin.Context) {
	var QueryParam request.UserQueryRequest
	var err error
	QueryParam.Query = c.Query("query")
	pageNumStr, ok := c.GetQuery("pagenum")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagenum, err = strconv.Atoi(pageNumStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	pageSizeStr, ok := c.GetQuery("pagesize")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Pagesize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	QueryParam.Role, ok = c.GetQuery("role")
	if !ok {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	switch QueryParam.Role {
	case "user":
		UserList, total, err := userService.QueryUsersList(QueryParam)
		if err != nil {
			response.FailWithDetailed(http.StatusInternalServerError, "获取用户列表失败", c)
			return
		}

		c.JSON(http.StatusOK, response.UserPageResponse{
			Data: &response.UserPageData{
				Total:   total,
				PageNum: int64(QueryParam.Pagenum),
				Users:   UserList,
			},
			Meta: &response.Meta{
				Msg:    "获取用户列表成功",
				Status: http.StatusOK,
			},
		})
		break
	case "companyUser":
		companyUserList, total, err := companyService.QueryUsersList(QueryParam)
		if err != nil {
			response.FailWithDetailed(http.StatusInternalServerError, "获取商家用户列表失败", c)
			return
		}

		c.JSON(http.StatusOK, response.CompanyUserResponse{
			Data: &response.CompanyUserData{
				Total:   total,
				PageNum: int64(QueryParam.Pagenum),
				Users:   companyUserList,
			},
			Meta: &response.Meta{
				Msg:    "获取商家用户列表成功",
				Status: http.StatusOK,
			},
		})
		break
	case "admin":
		AdminList, total, err := adminService.QueryUsersList(QueryParam)
		if err != nil {
			response.FailWithDetailed(http.StatusInternalServerError, "获取管理员列表失败", c)
			return
		}

		c.JSON(http.StatusOK, response.AdminResponse{
			Data: &response.AdminData{
				Total:   total,
				PageNum: int64(QueryParam.Pagenum),
				Users:   AdminList,
			},
			Meta: &response.Meta{
				Msg:    "获取管理员列表成功",
				Status: http.StatusOK,
			},
		})
		break
	}

}

func (u *UserApi) AddUser(c *gin.Context) {
	var addUserQuery request.AddUserRequest
	err := c.ShouldBind(&addUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := userService.AddUser(addUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusInternalServerError, "添加用户失败", c)
		return
	}
	c.JSON(http.StatusOK, response.AddUserResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "添加用户成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) UpdateUser(c *gin.Context) {
	var updateUserQuery request.UpdateUserRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := userService.UpdateUser(uid, updateUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新用户失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateUserResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "更新用户成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := userService.DeleteUser(uid)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	if flag == false {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除用户失败", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除用户成功", c)
	return
}

func (u *UserApi) AddCompanyUser(c *gin.Context) {
	var addUserQuery request.AddUserRequest
	err := c.ShouldBind(&addUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := companyService.AddCompanyUser(addUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "添加商家用户失败", c)
		return
	}
	c.JSON(http.StatusOK, response.AddCompanyUserResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "添加商家用户成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) UpdateCompanyUser(c *gin.Context) {
	var updateUserQuery request.UpdateUserRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := companyService.UpdateCompanyUser(uid, updateUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新商家用户失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateCompanyUserResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "更新商家用户成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) DeleteCompanyUser(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := companyService.DeleteCompanyUser(uid)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除商家用户失败", c)
		return
	}
	if flag == false {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除商家用户成功", c)
	return
}

func (u *UserApi) CheckComID(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := companyService.CheckComID(uid)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "检查商家用户id失败，服务器出错，请稍后再试", c)
		return
	}
	if flag == false {
		c.JSON(http.StatusOK, response.CheckComIDResponse{
			Data: &response.CheckComIdData{Result: "false"},
			Meta: &response.Meta{
				Msg:    "该商家用户id不存在",
				Status: http.StatusOK,
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.CheckComIDResponse{
		Data: &response.CheckComIdData{Result: "true"},
		Meta: &response.Meta{
			Msg:    "该商家用户id存在",
			Status: http.StatusOK,
		},
	})
	return
}

func (u *UserApi) AddAdmin(c *gin.Context) {
	var addUserQuery request.AddUserRequest
	err := c.ShouldBind(&addUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := adminService.AddAdmin(addUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "添加管理员失败", c)
		return
	}
	c.JSON(http.StatusOK, response.AddAdminResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "添加管理员成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) UpdateAdmin(c *gin.Context) {
	var updateUserQuery request.UpdateUserRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := adminService.UpdateAdmin(uid, updateUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新管理员信息失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateAdminResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "更新管理员信息成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) UpdateSelf(c *gin.Context) {
	var updateUserQuery request.UpdateSelfRequest
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	err := c.ShouldBind(&updateUserQuery)
	if err != nil {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	newUser, err := adminService.UpdateSelf(uid, updateUserQuery)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "更新管理员个人信息失败", c)
		return
	}
	c.JSON(http.StatusOK, response.UpdateAdminResponse{
		UserData: newUser,
		Meta: &response.Meta{
			Msg:    "更新管理员个人信息成功",
			Status: http.StatusOK,
		},
	})
}

func (u *UserApi) DeleteAdmin(c *gin.Context) {
	uid := c.Param("id")
	if uid == "" {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	flag, err := adminService.DeleteAdmin(uid)
	if err != nil {
		system.PrintError(err)
		response.FailWithDetailed(http.StatusInternalServerError, "删除管理员失败", c)
		return
	}
	if flag == false {
		response.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	response.SuccessWithNullData(http.StatusOK, "删除管理员成功", c)
	return
}
