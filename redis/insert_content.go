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
	//hget testgoredisforssdb_hash field:0
	//hdel testgoredisforssdb_hash field:0 field:1 field:2 field:3
	//zrange testgoredisforssdb_sortedset 0 5 withscores
	//zrem testgoredisforssdb_sortedset membera memberb memberc memberd
	for i, c := range contents {
		field := fmt.Sprintf("field:%d", i)
		pipeline.Process(s.Client.HSet("testgoredisforssdb_hash", field, c))
		cmd := s.Client.ZAdd("testgoredisforssdb_sortedset", redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%s", c),
		})
		pipeline.Process(cmd)
	}
	_, err := pipeline.Exec()
	if err != nil {
		return err
	}

	return nil
}
