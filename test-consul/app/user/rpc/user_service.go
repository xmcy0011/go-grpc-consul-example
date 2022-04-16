package rpc

import (
	"context"
	"test-consul/api/user"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (u *UserService) CreateUser(context.Context, *user.CreateUserReq) (*user.CreateUserRes, error) {
	return nil, nil
}