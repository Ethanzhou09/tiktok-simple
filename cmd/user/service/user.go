package service

import (
	"context"
	"fmt"
	"math/rand"
	"tiktok-simple/dal/db"
	"tiktok-simple/idl/kitex_gen/user"
	"tiktok-simple/pkg/crypt"
	"tiktok-simple/pkg/jwt"
	"time"
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

func (usrsrv *UserService) Register(ctx context.Context, req *user.UserRegisterRequest) (res *user.UserRegisterResponse, err error) {
	usr, err := db.GetUserByName(req.Username)
	if err != nil {
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, err
	}
	if usr != nil {
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：用户名已存在",
		}
		return res, nil
	}
	// 开始创建user账号
	rand.Seed(time.Now().UnixMilli())
	usr = &db.User{
		UserName: req.Username,
		// 密码md5加密
		Password: crypt.Md5Encrypt(req.Password),
		Avatar:   fmt.Sprintf("default%d.png", rand.Intn(10)),
	}
	if err := db.CreateUser(ctx, usr); err != nil {
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	}
	token, err := jwt.GenerateToken(usr.ID)
	if err != nil {
		res := &user.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	}
	res = &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

func (usrsrv *UserService) Login(ctx context.Context, req *user.UserLoginRequest) (res *user.UserLoginResponse, err error) {

	return
}

func (usrsrv *UserService) UserInfo(ctx context.Context, req *user.UserInfoRequest) (res *user.UserInfoResponse, err error) {

	return
}
