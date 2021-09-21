package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisConnection struct {
	RDB 	*redis.Client
}

var connection RedisConnection

func NewClient() *RedisConnection {
	rdbC := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	connection := &RedisConnection{
		RDB: rdbC,
	}
	return connection
}

func (client *RedisConnection) Set(key string, val string)  {
	err := client.RDB.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (client *RedisConnection) Get(key string) string {
	val, err := client.RDB.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return val
}

func (client *RedisConnection) Incr(key string) int64 {
	val, err := client.RDB.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return val
}