package caching

import (
	"context"
	"time"
)

type Cache interface {
	Putter
	Getter
	Deleter
}

type Putter interface {
	Put(ctx context.Context, key string, value string, duration time.Duration) error
}

type Getter interface {
	Get(ctx context.Context, key string, value string) error
}

type Deleter interface {
	Delete(ctx context.Context, key string) error
}
