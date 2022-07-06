package v1

import (
	"github.com/gin-gonic/gin"
	LogReq "jchz-admin/model/request"
	LogRes "jchz-admin/model/response"
	"jchz-admin/model/system"
	"jchz-admin/utils"
	"net/http"
	"strconv"
)

type BaseApi struct{}

var BaseApiGroup = new(BaseApi)

func (a *BaseApi) Login(c *gin.Context) {
	var form LogReq.LoginRequest

	if err := c.ShouldBind(&form); err != nil {
		LogRes.FailWithDetailed(http.StatusBadRequest, "参数错误", c)
		return
	}
	u := &system.Admin{Username: form.Username, Password: form.Password}
	admin, err := adminService.Login(u)
	if err != nil {
		LogRes.FailWithDetailed(http.StatusUnauthorized, "用户名或密码错误", c)
		return
	}
	uid := strconv.FormatUint(uint64(admin.ID), 10)
	JWT := utils.NewJWT()
	tok, err := JWT.CreateToken(uid)
	if err != nil {
		panic(err)
	}
	res := LogRes.LoginResponse{
		Data: &LogRes.LoginData{ID: admin.ID, Username: admin.Username, Avatar: admin.Avatar, Token: tok},
		Meta: &LogRes.Meta{Msg: "login ok", Status: http.StatusOK},
	}
	c.JSON(http.StatusOK, res)
}
