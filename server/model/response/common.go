package response

import (
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

type Result struct {
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta"`
}

func FailWithDetailed(status int, message string, c *gin.Context) {
	c.JSON(status, Result{
		Data: nil,
		Meta: &Meta{Msg: message, Status: status},
	})
}

func SuccessWithNullData(status int, message string, c *gin.Context) {
	c.JSON(status, Result{
		Data: nil,
		Meta: &Meta{Msg: message, Status: status},
	})
}
