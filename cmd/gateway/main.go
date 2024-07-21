package main

import (
	"context"
	"github.com/Ethanzhou09/tiktok-simple/cmd/gateway/handler"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func registerGroup(hz *server.Hertz) {
	douyin := hz.Group("/douyin")
	{
		user := douyin.Group("/user")
		{
			user.GET("/", )
			user.POST("/register/", handler.Login)
			user.POST("/login/",handler.Register )

		}
		// message := douyin.Group("/message")
		// {
		// 	message.GET("/chat/", )
		// 	message.POST("/action/",)
		// }
		// relation := douyin.Group("/relation")
		// {
		// 	// 粉丝列表
		// 	relation.GET("/follower/list/", )
		// 	// 关注列表
		// 	relation.GET("/follow/list/",)
		// 	// 朋友列表
		// 	relation.GET("/friend/list/",)
		// 	relation.POST("/action/",)
		// }
		// publish := douyin.Group("/publish")
		// {
		// 	publish.GET("/list/",)
		// 	publish.POST("/action/",)
		// }
		douyin.GET("/feed",)
		// favorite := douyin.Group("/favorite")
		// {
		// 	favorite.POST("/action/",)
		// 	favorite.GET("/list/",)
		// }
		// comment := douyin.Group("/comment")
		// {
		// 	comment.POST("/action/",)
		// 	comment.GET("/list/",)
		// }
	}
}

func InitHertz() *server.Hertz {
    h := server.Default()
    registerGroup(h)
    return h
}

func main() {
	hz := InitHertz()

	registerGroup(hz)

    hz.Spin()
}