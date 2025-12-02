package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/domain/user/repository"
)

type UserDomain interface {
	CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx *gin.Context, name string) (*entity.User, error)
	CheckPwd(ctx *gin.Context, user *entity.User, pwd string) error
}

type UserDomainImpl struct {
	repo       repository.User
	PwdEncrypt repository.PwdEncrypt
}

func NewUserDomainImpl(repo repository.User, pwd repository.PwdEncrypt) UserDomain {
	return &UserDomainImpl{repo: repo, PwdEncrypt: pwd}
}

func (u *UserDomainImpl) CreateUser(ctx *gin.Context, user *entity.User) (*entity.User, error) {
	encryptPwd, err := u.PwdEncrypt.Encrypt([]byte(user.PassWord))
	if err != nil {
		return nil, err
	}
	user.SetPwd(encryptPwd)
	return u.repo.CreateUser(ctx, user)
}

func (u *UserDomainImpl) FindUserByName(ctx *gin.Context, username string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, username)
}

func (u *UserDomainImpl) CheckPwd(ctx *gin.Context, user *entity.User, pwd string) error {
	if u.PwdEncrypt.Check([]byte(user.PassWord), []byte(pwd)) {
		return nil
	}
	return errors.New("pwd is error")
}
