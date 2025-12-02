package user

import (
	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

func RegisterResponse(user *entity.User) *types.UserResp {
	return &types.UserResp{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

func LoginResponse(user *entity.User, token string) *types.TokenData {
	return &types.TokenData{
		User: &types.UserResp{
			ID:        user.ID,
			UserName:  user.UserName,
			CreatedAt: user.CreatedAt.Unix(),
		},
		Token: token,
	}
}
