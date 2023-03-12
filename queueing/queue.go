package queueing

import (
	"context"
	"time"
)

type Pusher interface {
	Push(ctx context.Context, queue string) (interface{}, error)
}

type Popper interface {
	Pop(ctx context.Context, queue string, data interface{}, dur time.Duration) error
}
