package main

import (
	"tiktok-simple/cmd/gateway/rpc"
	"tiktok-simple/cmd/gateway/router"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
)

type Server struct{
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Cfg struct {
	Server Server `mapstructure:"server"`
}

func InitHertz() *server.Hertz {
	v := viper.New()
	v.SetConfigFile("../../cfg/gateway.yml")
	v.ReadInConfig()
	cfg := Cfg{}
	v.Unmarshal(&cfg)
	servercfg := cfg.Server
	h := server.Default(server.WithHostPorts(servercfg.Host + ":" + servercfg.Port))
	return h
}

func main() {
	rpc.InitRPC()

	hz := InitHertz()

	router.RegisterGroup(hz)

	hz.Spin()
}
