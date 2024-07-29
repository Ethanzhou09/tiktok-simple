package main

import (
	"tiktok-simple/cmd/user/cfginit"
	"log"
	"tiktok-simple/cmd/user/service"
	"tiktok-simple/idl/kitex_gen/user/userservice"
	"tiktok-simple/dal/db"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)


func main(){
	db.Init()
	cfginit.InitViper()
	r := cfginit.EtcdInit()
	addr := cfginit.GetSrvAddr()
	srv := service.GetUserSrv()
	server := userservice.NewServer(srv,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "User"}), server.WithRegistry(r), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
    err := server.Run()
    if err != nil {
        log.Fatal(err)
    }
}