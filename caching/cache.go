package caching

import (
	"context"
	"time"
)

type Cache interface {
	Put(ctx context.Context, key string, value string, duration time.Duration) error
	Get(ctx context.Context, key string, value string) error
	Delete(ctx context.Context, key string, value string) error
}
