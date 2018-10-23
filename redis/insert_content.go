package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

type Inserter struct {
	Client *redis.Client
}

func (s *Inserter) Insert(ctx context.Context, contents []string) error {
	pipeline := s.Client.Pipeline()
	//hget testgoredisforssdb field:0
	//hdel testgoredisforssdb field:0 field:1 field:2 field:3
	for i, c := range contents {
		field := fmt.Sprintf("field:%d", i)
		pipeline.Process(s.Client.HSet("testgoredisforssdb", field, c))
	}
	_, err := pipeline.Exec()
	if err != nil {
		return err
	}

	return nil
}
