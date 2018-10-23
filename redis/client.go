package redis

import "github.com/go-redis/redis"

// NewClient create a redis client
func NewClient(addr string, db int, poolSize int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		PoolSize: poolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil

}
