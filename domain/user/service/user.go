package service

import (
	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/domain/user/repository"
)

type UserDomain interface {
	CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx *gin.Context, name string) (*entity.User, error)
}

type UserDomainImpl struct {
	repo    repository.User
	encrypt repository.PwdEncrypt
}

func NewUserDomainImpl(repo repository.User, encrypt repository.PwdEncrypt) UserDomain {
	return &UserDomainImpl{repo: repo, encrypt: encrypt}
}

func (u *UserDomainImpl) CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserDomainImpl) FindUserByName(ctx *gin.Context, name string) (*entity.User, error) {
	return nil, nil
}
