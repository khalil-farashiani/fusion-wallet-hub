package rabbitmq

import (
	"fmt"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (a *Adapter) Consume(queueName entity.Event) (<-chan amqp.Delivery, error) {
	msgs, err := a.channel.Consume(
		string(queueName), // queue
		"",                // consumer tag - empty string will let RabbitMQ generate a unique one
		false,             // auto-acknowledge messages - set to false for manual ack
		false,             // exclusive
		false,
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("error setting up consumer: %v", err)
	}
	return msgs, nil
}
