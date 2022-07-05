package middleware

import (
	"github.com/gin-gonic/gin"
	"jchz-admin/core/system"
	"jchz-admin/model/response"
	"jchz-admin/utils"
	"net/http"
	"time"
)

// JWTAuth JWT鉴权
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithDetailed(http.StatusUnauthorized, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(http.StatusUnauthorized, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(http.StatusUnauthorized, err.Error(), c)
			c.Abort()
			return
		}

		// 更新 token 以续期
		if time.Now().Unix()-claims.IssuedAt.Unix() > claims.BufferTime {
			newToken, err := j.CreateTokenByOldToken(token, claims)
			if err != nil {
				system.PrintError(err)
				response.FailWithDetailed(http.StatusInternalServerError, err.Error(), c)
				c.Abort()
				return
			}
			if err != nil {
				system.PrintError(err)
				response.FailWithDetailed(http.StatusInternalServerError, err.Error(), c)
				c.Abort()
				return
			}
			c.Header("newToken", newToken)
		}
		c.Set("claims", claims)
		c.Next()
	}
}
