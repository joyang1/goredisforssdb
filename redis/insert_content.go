package redis

import (
	"context"

	"github.com/go-redis/redis"
)

type Inserter struct {
	client *redis.Client
}

func (s *Inserter) Insert(ctx context.Context, content []string) error {
	return nil
}
