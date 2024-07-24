package main

import (
	"tiktok-simple/cmd/comment/cfginit"
	"log"
	"tiktok-simple/cmd/comment/service"
	"tiktok-simple/idl/kitex_gen/comment/commentservice"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)


func main(){
	cfginit.InitViper()
	r := cfginit.EtcdInit()
	addr := cfginit.GetSrvAddr()
	srv := service.GetCommentSrv()
	server := commentservice.NewServer(srv,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Video"}), server.WithRegistry(r), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
    err := server.Run()
    if err != nil {
        log.Fatal(err)
    }
}