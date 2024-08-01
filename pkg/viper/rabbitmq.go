package viper

import (
	"github.com/spf13/viper"
)


type Rabbitmq struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string    `mapstructure:"port"`
}

func InitRabbitmq()string {
	v = viper.New() 
	viper.SetConfigFile("../../cfg/rabbitmq.yml")
	v.ReadInConfig()
	var cfg Rabbitmq
	v.Unmarshal(&cfg)
	MqUrl := "amqp://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" +cfg.Port
	return MqUrl
}