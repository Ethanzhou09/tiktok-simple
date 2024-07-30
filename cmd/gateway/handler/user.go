package handler

import (
	"context"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktok-simple/internal/response"
	"tiktok-simple/cmd/gateway/rpc"
	"tiktok-simple/idl/kitex_gen/user"
)


func Register(ctx context.Context, c *app.RequestContext){
	username := c.Query("username")
	password := c.Query("password")
	//校验参数
	if len(username) == 0 || len(password) == 0 {
		c.JSON(http.StatusBadRequest, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "用户名或密码不能为空",
			},
		})
		return
	}
	if len(username) > 32 || len(password) > 32 {
		c.JSON(http.StatusOK, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "用户名或密码长度不能大于32个字符",
			},
		})
		return
	}
	req := &user.UserRegisterRequest{
		Username: username,
		Password: password,
	}
	res,err := rpc.Register(ctx,req)
	if err != nil {
		c.JSON(http.StatusOK, response.Register{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "注册失败",
			},
		})
		return
	}
	c.JSON(http.StatusOK, response.Register{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserID: res.UserId,
		Token: res.Token,
	})
}

func Login(ctx context.Context, c *app.RequestContext){
	

}

func UserInfo(ctx context.Context, c *app.RequestContext){
	

}