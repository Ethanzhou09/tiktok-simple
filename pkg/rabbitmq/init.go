package rabbitmq

import(
	"tiktok-simple/pkg/viper"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	MqUrl string
	conn *amqp.Connection
)

func Init() {
	MqUrl = viper.InitRabbitmq()
}