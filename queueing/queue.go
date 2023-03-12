package queueing

import "context"

type Queue interface {
	Produce(ctx context.Context, queue string, data interface{}) error
	Consume(ctx context.Context, queue string, data interface{}) (<-chan Message, error)
}

type Message struct {
	Header map[string]string
	Data   []byte
}
