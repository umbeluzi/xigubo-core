package pubsub

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, msg *Message) error
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan Message, error)
	Close() error
}

type Message struct {
	ID          string
	ContentType string
	Data        []byte
	AckFunc     func() error
	NackFunc    func() error
}

func (m Message) Ack() error {
	return m.AckFunc()
}

func (m Message) Nack() error {
	return m.NackFunc()
}
