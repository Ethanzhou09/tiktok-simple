package service

import (
	"context"
	"fmt"
	"math/rand"
	"tiktok-simple/dal/db"
	"tiktok-simple/idl/kitex_gen/user"
	"tiktok-simple/pkg/crypt"
	"tiktok-simple/pkg/jwt"
	"tiktok-simple/pkg/minio"
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
	usr, err := db.GetUserByName(req.Username)
	if err!=nil{
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：服务器内部错误",
		}
		return res, err
	}
	if usr == nil {
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：用户名不存在",
		}
		return res, nil
	}
	if usr.Password!= crypt.Md5Encrypt(req.Password) {
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：密码错误",
		}
		return res, nil
	}
	token, err := jwt.GenerateToken(usr.ID)
	if err != nil {
		res := &user.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：token生成失败",
		}
		return res, nil
	}
	res = &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      token,
	}
	return res, nil
}

func (usrsrv *UserService) UserInfo(ctx context.Context, req *user.UserInfoRequest) (res *user.UserInfoResponse, err error) {
	userid := req.UserId
	usr,err := db.GetUserById(ctx,userid)
	if err!=nil{
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：服务器内部错误",
		}
		return res, err
	}
	if usr == nil {
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "获取用户信息失败：用户不存在",
		}
		return res, nil
	}
	avatar, err := minio.GetFileTemporaryURL(minio.AvatarBucketName, usr.Avatar)

	if err != nil {
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：获取头像失败",
		}
		return res, nil
	}
	backgroundImage, err := minio.GetFileTemporaryURL(minio.BackgroundImageBucketName, usr.BackgroundImage)
	if err != nil {
		res := &user.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "服务器内部错误：获取背景图失败",
		}
		return res, nil
	}
	//返回结果
	res = &user.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User: &user.User{
			Id:              int64(usr.ID),
			Name:            usr.UserName,
			FollowCount:     int64(usr.FollowingCount),
			FollowerCount:   int64(usr.FollowerCount),
			IsFollow:        userid == int64(usr.ID),
			Avatar:          avatar,
			BackgroundImage: backgroundImage,
			Signature:       usr.Signature,
			TotalFavorited:  int64(usr.TotalFavorited),
			WorkCount:       int64(usr.WorkCount),
			FavoriteCount:   int64(usr.FavoriteCount),
		},
	}
	return res, nil
}
