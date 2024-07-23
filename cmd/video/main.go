package main

import (
	"tiktok-simple/cmd/video/cfginit"
	"log"
	"tiktok-simple/cmd/video/service"
	"tiktok-simple/idl/kitex_gen/video/videoservice"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)


func main(){
	cfginit.InitViper()
	r := cfginit.EtcdInit()
	addr := cfginit.GetSrvAddr()
	videosrv := service.GetVideoSrv()
	server := videoservice.NewServer(videosrv,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Video"}), server.WithRegistry(r), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
    err := server.Run()
    if err != nil {
        log.Fatal(err)
    }
}