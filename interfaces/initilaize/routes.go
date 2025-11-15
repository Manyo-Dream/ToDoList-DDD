package initilaize

import (
	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/interfaces/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use()
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "success"})
		})

		// 用户部分
		v1.POST("user/register", controller.UserRegisterHandler())
		v1.POST("user/login", controller.UserRegisterHandler())
		// 备忘录部分
	}
	return r
}
