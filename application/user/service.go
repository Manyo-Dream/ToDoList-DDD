package user

import (
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

func (s *ServiceImpl) Register(ctx *gin.Context, user *entity.User) (any, error) {
	user, err := s.ud.UserCreate(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
