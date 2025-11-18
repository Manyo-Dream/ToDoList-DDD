package service

import (
	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
)

type UserDomain interface {
	UserCreate(ctx *gin.Context, user *entity.User) (*entity.User, error)
}

type UserDomainImpl struct{}

func NewUserDomainImpl() UserDomain {
	return &UserDomainImpl{}
}

func (u *UserDomainImpl) UserCreate(ctx *gin.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}
