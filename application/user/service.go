package user

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/domain/user/service"
	"github.com/manyodream/todolist-ddd/infrastructure/auth"
)

type Service interface {
	Register(ctx *gin.Context, user *entity.User) (any, error)
	Login(ctx *gin.Context, user *entity.User) (any, error)
}

type ServiceImpl struct {
	ud           service.UserDomain
	tokenService auth.TokenService
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func NewServiceImpl(srv service.UserDomain, ts auth.TokenService) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{
			ud:           srv,
			tokenService: ts,
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

	if userExist != nil && userExist.IsActive() {
		return nil, errors.New("[Register] user is already exist")
	}

	// 创建用户
	user, err := s.ud.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return RegisterResponse(user), nil
}

func (s *ServiceImpl) Login(ctx *gin.Context, userEntity *entity.User) (any, error) {
	// 用户名是否重复
	userExist, err := s.ud.FindUserByName(ctx, userEntity.UserName)
	if err != nil {
		return nil, err
	}

	// 检查密码
	err = s.ud.CheckPwd(ctx, userExist, userEntity.PassWord)
	if err != nil {
		return nil, err
	}

	// 返回token
	token, err := s.tokenService.GenerateToken(ctx, userExist.ID, userExist.UserName)
	if err != nil {
		return nil, err
	}
	return LoginResponse(userExist, token), nil
}
