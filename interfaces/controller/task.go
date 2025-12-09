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

		ctx.JSON(http.StatusOK, types.RespSuccessWithData(types.Entity2TaskResp(result)))
	}
}

func TaskListHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task list binding error"))
			return
		}
		results, err := task.ServiceImplIns.TaskList(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task list error"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(results))
	}
}

func TaskDetailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskDetailReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task detail binding error"))
			return
		}
		result, err := task.ServiceImplIns.TaskDetail(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task detail error"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(result))
	}
}

func TaskUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task update binding error"))
			return
		}
		err := task.ServiceImplIns.TaskUpdate(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task update error"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}

func TaskSearchHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskSearchReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task search binding error"))
			return
		}
		result, err := task.ServiceImplIns.TaskSearch(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task search error"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccessWithData(result))
	}
}

func TaskDeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskDeleteReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task delete binding error"))
			return
		}
		err := task.ServiceImplIns.TaskDelete(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, types.RespError(err, "task delete error"))
			return
		}
		ctx.JSON(http.StatusOK, types.RespSuccess())
	}
}
