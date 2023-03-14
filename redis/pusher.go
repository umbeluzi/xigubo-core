package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewPusher(rdb *redis.Client) *Subscriber {
	return &Subscriber{
		client: rdb,
	}
}

type Pusher struct {
	client *redis.Client
}

func (s *Pusher) Push(ctx context.Context, queue string, data []byte) error {
	return s.client.RPush(ctx, queue, data).Err()
}
