package task

import (
	"time"

	"github.com/manyodream/todolist-ddd/interfaces/types"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
)

func ListResponse(list []*entity.Task, count int64) types.List[*entity.Task] {
	return types.List[*entity.Task]{
		Items: list,
		Count: count,
	}
}

func UpdateRep2TaskEntity(tid, uid uint, username string, taskReq *types.TaskUpdateReq) *entity.Task {
	return &entity.Task{
		ID:       tid,
		UID:      uid,
		UserName: username,
		Title:    taskReq.Title,
		Status:   taskReq.Status,
		Content:  taskReq.Content,
		UpdateAt: time.Now(),
	}
}
