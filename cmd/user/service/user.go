package service

import (
	"context"
	"tiktok-simple/idl/kitex_gen/user"

)

type UserService struct{}

var UserSrv *UserService

// 懒汉单例获取UserSrv
func GetUserSrv() *UserService {
	if UserSrv == nil {
		UserSrv = new(UserService)
	}
	return UserSrv
}

func (usrsrv *UserService)Register(ctx context.Context, req *user.UserRegisterRequest)(res *user.UserRegisterResponse, err error){

	return
}


func (usrsrv *UserService)Login(ctx context.Context, req *user.UserLoginRequest)(res *user.UserLoginResponse, err error){

	return
}


func (usrsrv *UserService)UserInfo(ctx context.Context, req *user.UserInfoRequest)(res *user.UserInfoResponse, err error){

	return
}