package redis

import (
	"context"

	"github.com/getumbeluzi/xigubo-core/pubsub"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) Subscribe(ctx context.Context, topic string) (<-chan pubsub.Message, error) {
	events := r.Client.Subscribe(ctx, topic)

	c := make(chan pubsub.Message)

	go func() {
		for {
			msg, err := events.ReceiveMessage(ctx)
			if err != nil {
				return
			}

			c <- pubsub.Message{
				Data: []byte(msg.Payload),
			}
		}
	}()

	return c, nil

}
