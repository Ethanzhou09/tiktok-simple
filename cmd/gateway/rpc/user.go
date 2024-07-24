package rpc

import (
	"context"
	"tiktok-simple/idl/kitex_gen/user"
	"errors"
)

func Register(ctx context.Context, req *user.UserRegisterRequest) (res *user.UserRegisterResponse, err error) {
	if req == nil {
		return nil,errors.New("register req is nil")
	}
	return UserClient.Register(ctx,req)
}

func Login(ctx context.Context, req *user.UserLoginRequest) (res *user.UserLoginResponse, err error) {
	if req == nil {
		return nil,errors.New("login req is nil")
	}
	return UserClient.Login(ctx,req)
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (res *user.UserInfoResponse, err error) {
	if req == nil {
		return nil,errors.New("userInfo req is nil")
	}
	return UserClient.UserInfo(ctx,req)
}
