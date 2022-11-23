package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func WriteCacheOnly() {

	var ctx = context.Background()
	rdb := client()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

}

func ReadCacheOnly() {

	var ctx = context.Background()
	rdb := client()

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}

func client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
