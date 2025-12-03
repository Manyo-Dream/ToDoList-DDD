package types

import (
	task "github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
)

func UserReq2Entity(user *UserReq) *entity.User {
	return &entity.User{
		UserName: user.UserName,
		PassWord: user.PassWord,
	}
}

func Entity2TaskResp(task *task.Task) *TaskResp {
	return &TaskResp{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
		Status:  task.Status,
	}
}
