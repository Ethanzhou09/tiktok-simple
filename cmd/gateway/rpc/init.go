package rpc

import (
	"fmt"
	"log"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"github.com/cloudwego/kitex/client"
	"time"

	"tiktok-simple/idl/kitex_gen/comment/commentservice"
	"tiktok-simple/idl/kitex_gen/favorite/favoriteservice"
	"tiktok-simple/idl/kitex_gen/message/messageservice"
	"tiktok-simple/idl/kitex_gen/relation/relationservice"
	"tiktok-simple/idl/kitex_gen/user/userservice"
	"tiktok-simple/idl/kitex_gen/video/videoservice"
)

var (
	UserClient     userservice.Client
	VideoClient    videoservice.Client
	CommentClient  commentservice.Client
	RelationClient relationservice.Client
	FavoriteClient favoriteservice.Client
	MessageClient  messageservice.Client
)

type Cfg struct {
	EtcdCfg EtcdCfg `mapstructure:"etcd"`
}

type EtcdCfg struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func InitRPC() {
	v := viper.New()
	v.SetConfigFile("../../cfg/gateway.yml")
	v.ReadInConfig()
	cfg := Cfg{}

	v.Unmarshal(&cfg)
	etcdcfg := cfg.EtcdCfg
	r, err := etcd.NewEtcdResolver([]string{fmt.Sprintf("%s:%d", etcdcfg.Host, etcdcfg.Port)})
	if err!= nil {
		log.Fatal(err)
	}

	UserClient, err = userservice.NewClient("user", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}

	VideoClient, err = videoservice.NewClient("video", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}

	CommentClient, err = commentservice.NewClient("comment", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}

	RelationClient, err = relationservice.NewClient("relation", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}

	FavoriteClient, err = favoriteservice.NewClient("favorite", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}

	MessageClient, err = messageservice.NewClient("message", client.WithResolver(r),client.WithConnectTimeout(5*time.Second),
	client.WithRPCTimeout(0))
	if err!= nil {
		log.Fatal(err)
	}
}