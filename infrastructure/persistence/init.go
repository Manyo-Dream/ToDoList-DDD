package persistence

import (
	taskRepo "github.com/manyodream/todolist-ddd/domain/task/repository"
	"github.com/manyodream/todolist-ddd/domain/user/repository"
	taskPer "github.com/manyodream/todolist-ddd/infrastructure/persistence/task"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.User
	Task taskRepo.Task
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepository(db),
		Task: taskPer.NewRepository(db),
	}
}
