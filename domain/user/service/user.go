package service

import (
	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
)

type UserDomain interface {
	CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx *gin.Context, name string) (*entity.User, error)
}

type UserDomainImpl struct{}

func NewUserDomainImpl() UserDomain {
	return &UserDomainImpl{}
}

func (u *UserDomainImpl) CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserDomainImpl) FindUserByName(ctx *gin.Context, name string) (*entity.User, error) {
	return nil, nil
}
