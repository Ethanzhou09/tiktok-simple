package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok-simple/cmd/gateway/handler"
)

func registerGroup(hz *server.Hertz) {
	douyin := hz.Group("/douyin")
	{
		user := douyin.Group("/user")
		{
			user.GET("/", handler.UserInfo)
			user.POST("/register/", handler.Login)
			user.POST("/login/", handler.Register)
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
		douyin.GET("/feed", handler.Feed)
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
	return h
}

func main() {
	hz := InitHertz()

	registerGroup(hz)

	hz.Spin()
}
