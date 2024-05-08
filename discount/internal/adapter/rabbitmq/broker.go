package rabbitmq

import (
	"context"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func (a *Adapter) Publish(event entity.Event, payload string) error {
	//better to get timeout from config
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		err := a.channel.PublishWithContext(
			ctx,
			"",            // exchange
			string(event), // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(payload),
				DeliveryMode: 2,
			},
		)
		if err != nil {
			log.Printf("publish error: %v\n", err)
			return err
		}
		return nil
	}
}
