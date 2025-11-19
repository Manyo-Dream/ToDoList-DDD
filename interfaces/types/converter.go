package types

import "github.com/manyodream/todolist-ddd/domain/user/entity"

func UserReq2Entity(user *UserReq) *entity.User {
	return &entity.User{
		UserName: user.UserName,
		PassWord: user.PassWord,
	}
}
