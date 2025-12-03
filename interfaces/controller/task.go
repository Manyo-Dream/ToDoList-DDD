package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/application/task"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

func TaskCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task create binding error"))
			return
		}
		result, err := task.ServiceImplIns.TaskCreate(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task create error"))
			return
		}
		resp := types.Entity2TaskResp(result)
		ctx.JSON(http.StatusOK, resp)
	}
}

func TaskListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func TaskDetailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func TaskUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func TaskSearchHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func TaskDeteleHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
