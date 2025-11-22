package user

import "github.com/manyodream/todolist-ddd/domain/user/entity"

func Entity2PO(user *entity.User) *User {
	return &User{
		UserName:       user.UserName,
		PassWordDigest: user.PassWord,
	}
}

func PO2Entity(user *User) *entity.User {
	return &entity.User{
		ID:        int64(user.ID),
		UserName:  user.UserName,
		PassWord:  user.PassWordDigest,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
