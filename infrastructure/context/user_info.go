package context

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("get user info from context error")
	}
	return user, nil
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	user, ok := ctx.Value(userKey).(*UserInfo)
	return user, ok
}

func NewContext(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, user)
}
