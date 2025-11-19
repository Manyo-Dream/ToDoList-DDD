package user

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/domain/user/service"
)

type Service interface {
	Register(ctx *gin.Context, user *entity.User) (any, error)
}

type ServiceImpl struct {
	ud service.UserDomain
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func NewServiceImpl(srv service.UserDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{
			ud: srv,
		}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) Register(ctx *gin.Context, userEntity *entity.User) (any, error) {
	// 用户名是否重复
	userExist, err := s.ud.FindUserByName(ctx, userEntity.UserName)
	if err != nil {
		return nil, err
	}
	if userExist.IsActive() {
		return nil, errors.New("user is already exsit")
	}

	// 创建用户
	user, err := s.ud.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}
