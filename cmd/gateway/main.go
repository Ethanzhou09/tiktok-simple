package main

import (
	"tiktok-simple/cmd/gateway/rpc"
	"tiktok-simple/cmd/gateway/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)


func InitHertz() *server.Hertz {
	h := server.Default()
	return h
}

func main() {
	rpc.InitRPC()

	hz := InitHertz()

	router.RegisterGroup(hz)

	hz.Spin()
}
