package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/consts"
	"github.com/manyodream/todolist-ddd/infrastructure/auth"
	"github.com/manyodream/todolist-ddd/infrastructure/context"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = consts.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = consts.NotFound
			c.JSON(http.StatusNotFound, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}
		jwtService := auth.NewJWTTokenService()
		claims, err := jwtService.ParseToken(c.Request.Context(), token)
		if err != nil {
			code = consts.Error
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = consts.InvalidParams
		}
		if code != consts.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": code,
				"msg":    consts.GetMsg(code),
				"data":   "身份可能失效，请重新登录",
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(
			context.NewContext(
				c.Request.Context(),
				&context.UserInfo{
					Id:   claims.ID,
					Name: claims.Username,
				},
			),
		)
		c.Next()
	}
}
