package main

import (
	"github.com/spf13/viper"
	"fmt"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/pkg/registry"
)
type Cfg struct{
	EtcdCfg EtcdCfg `mapstructure:"etcd"`
}

type EtcdCfg struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

func EtcdInit() registry.Registry{
	v := viper.New()
	v.SetConfigFile("../cfg/user.yml")
	v.ReadInConfig()
	cfg := Cfg{}
	v.Unmarshal(&cfg)
	etcdcfg := cfg.EtcdCfg
	r, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%d", etcdcfg.Host, etcdcfg.Port)})
	if err!= nil {
		panic(err)
	}
	return r
}