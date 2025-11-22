package persistence

import (
	"github.com/manyodream/todolist-ddd/domain/user/repository"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepository(db),
	}
}
