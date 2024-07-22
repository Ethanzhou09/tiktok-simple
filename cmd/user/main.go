package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"tiktok-simple/idl/kitex_gen/user/userservice"
	"tiktok-simple/cmd/user/service"
	"github.com/cloudwego/kitex/server"
	"log"
)


func main(){
	r := EtcdInit()
	usersrv := service.GetUserSrv()
	server := userservice.NewServer(usersrv,server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "User"}), server.WithRegistry(r))
    err := server.Run()
    if err != nil {
        log.Fatal(err)
    }
}