package cfginit

import (
	"fmt"
	"net"
	"log"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
)

var v *viper.Viper


type Cfg struct{
	Srv SrvCfg `mapstructure:"srv"`
	EtcdCfg EtcdCfg `mapstructure:"etcd"`
}

type SrvCfg struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type EtcdCfg struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

func InitViper(){
	v = viper.New()
	v.SetConfigFile("../../cfg/video.yml")
	v.ReadInConfig()
}

func EtcdInit() registry.Registry{
	cfg := Cfg{}
	v.Unmarshal(&cfg)
	etcdcfg := cfg.EtcdCfg
	r, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%d", etcdcfg.Host, etcdcfg.Port)})
	if err!= nil {
		panic(err)
	}
	return r
}

func GetSrvAddr()net.Addr{
	cfg := Cfg{}
	v.Unmarshal(&cfg)
	srvcfg := cfg.Srv
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", srvcfg.Host, srvcfg.Port))
	if err!= nil {
		log.Fatal(err)
	}
	return addr
}
