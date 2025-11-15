package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "user register binding error"))
			return
		}

		// 定义entity
		userEntity := types.UserReq2Entity(&req)
	}
}
