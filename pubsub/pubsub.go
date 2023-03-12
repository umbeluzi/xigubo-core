package pubsub

import "context"

type PubSub interface {
	Publish(ctx context.Context, topic string, data []byte) error
	Subscribe(ctx context.Context, topic string) (<-chan Message, error)
}

type Message struct {
	Data []byte
}
