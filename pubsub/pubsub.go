package pubsub

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, data []byte) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan Message, error)
}

type Message struct {
	Data []byte
}
