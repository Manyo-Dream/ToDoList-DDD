package task

import (
	"time"

	"github.com/manyodream/todolist-ddd/domain/infrastructure/persistence/user"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
)

func Entity2PO(task *entity.Task) *Task {
	return &Task{
		Uid:       task.UID,
		Title:     task.Title,
		Staust:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func PO2Entity(task *Task, user *user.User) *entity.Task {
	return &entity.Task{
		ID:        task.ID,
		UID:       task.Uid,
		UserName:  user.UserName,
		Title:     task.Title,
		Status:    task.Staust,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
		CreateAt:  time.Time{},
		UpdateAt:  time.Time{},
	}
}
