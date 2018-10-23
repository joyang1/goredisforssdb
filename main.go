package main

import (
	"context"
	"mygo/goredisforssdb/redis"
)

func main() {
	client, err := redis.NewClient("10.10.73.140:8765", 0, 10)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	inserter := &redis.Inserter{Client: client}

	contents := make([]string, 0)
	contents = append(contents, "a", "b", "c", "d")
	err = inserter.Insert(ctx, contents)
	if err != nil {
		panic(err)
	}
}
