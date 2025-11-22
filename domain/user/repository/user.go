package repository

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/user/entity"
)

type User interface {
	UserBase
}

type UserBase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByName(ctx context.Context, name string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
}
