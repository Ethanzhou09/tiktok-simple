package viper 

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

type DalConfig struct{
	Db Db `mapstructure:"db"`
	Redis Redis `mapstructure:"redis"`
}
type Db struct{
	Drivername string `mapstructure:"drivername"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset string `mapstructure:"charset"`
}
type Redis struct{
	Host string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

func Init(){
	v = viper.New()
	v.SetConfigFile("../../cfg/dal.yml")
	v.ReadInConfig()
}

func GetViper() *viper.Viper{
	return v
}

func GetdbConfig() *Db{
	var config DalConfig
	v.Unmarshal(&config)
	return &config.Db
}