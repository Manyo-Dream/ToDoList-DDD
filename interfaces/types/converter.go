package types

import "github.com/manyodream/todolist-ddd/domain/userentity"

func UserReq2Entity(user *UserReq) *userentity.User {
	return &userentity.User{
		UserName: user.UserName,
		PassWord: user.PassWord,
	}
}
