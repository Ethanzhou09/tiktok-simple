package main

import (
	"tiktok-simple/cmd/favorite/cfginit"
	"log"
	"tiktok-simple/cmd/favorite/service"
	"tiktok-simple/idl/kitex_gen/favorite/favoriteservice"
	"time"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)


func main(){
	cfginit.InitViper()
	r := cfginit.EtcdInit()
	addr := cfginit.GetSrvAddr()
	srv := service.GetFavoriteSrv()
	server := favoriteservice.NewServer(srv,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Favorite"}), server.WithRegistry(r), server.WithServiceAddr(addr), server.WithReadWriteTimeout(5*time.Second), server.WithExitWaitTime(5*time.Second))
    err := server.Run()
    if err != nil {
        log.Fatal(err)
    }
}