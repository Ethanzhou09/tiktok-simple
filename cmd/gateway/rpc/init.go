package rpc

import (
	"github.com/spf13/viper"
	"fmt"
)
type Cfg struct{
	EtcdCfg EtcdCfg `mapstructure:"etcd"`
}

type EtcdCfg struct{
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

func Init() {
	v := viper.New()
	v.SetConfigFile("../../cfg/gateway.yml")
	v.ReadInConfig()
	etcdcfg := Cfg{}

	v.Unmarshal(&etcdcfg)
	fmt.Println(etcdcfg.EtcdCfg)
}