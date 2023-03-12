package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Queue struct {
	Client *redis.Client
}

func (q *Queue) Get(ctx context.Context, key string) (string, error) {
	val, err := q.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (q *Queue) Put(ctx context.Context, key string, value string, duration time.Duration) error {
	if err := q.Client.Set(ctx, key, value, duration).Err(); err != nil {
		return err
	}

	return nil
}
