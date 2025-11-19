package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
)

type User interface {
	UserBase
}

type UserBase interface {
	CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error)
	GetUserByName(ctx *gin.Context, name string) (*entity.User, error)
	GetUserByID(ctx *gin.Context, id int64) (*entity.User, error)
}
