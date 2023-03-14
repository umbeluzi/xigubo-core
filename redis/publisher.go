package redis

import (
	"context"

	"github.com/getumbeluzi/xigubo-core/pubsub"
	"github.com/redis/go-redis/v9"
)

func NewSubscriber(rdb *redis.Client) *Subscriber {
	return &Subscriber{
		client: rdb,
	}
}

type Subscriber struct {
	client *redis.Client
}

func (s *Subscriber) Subscribe(ctx context.Context, topic string) (<-chan pubsub.Message, error) {
	events := s.client.Subscribe(ctx, topic)

	event := make(chan pubsub.Message)

	go func() {
		for {
			msg, err := events.ReceiveMessage(ctx)
			if err != nil {
				return
			}

			event <- pubsub.Message{
				Data: []byte(msg.Payload),
			}
		}
	}()

	return event, nil
}
