package rabbitmq

import (
	"fmt"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	UserName string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
}

type Adapter struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func New(config Config) *Adapter {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.UserName, config.Password, config.Host, config.Port))
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		panic(err)
	}
	_, err = channel.QueueDeclare(
		string(entity.AddAmountEvent),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		channel.Close()
		conn.Close()
		panic(err)
	}
	return &Adapter{
		connection: conn,
		channel:    channel,
	}
}

func (a *Adapter) Close() {
	if a.channel != nil {
		a.channel.Close()
	}
	if a.connection != nil {
		a.connection.Close()
	}
}
