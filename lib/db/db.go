package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/jackc/pgx/v5/pgxpool"
)

// always look into cache first, go DB if no cache is found.
func Get(model interface{}) error {

	err := GetCache(model)
	//no cache for this query, 1)fetch from db. 2)cache the record.
	if err == redis.Nil {

	}
	if err != nil {
		return err
	}

	return nil
}

func CachedInsert() {

}

func CachedUpdate() {

}

func DeleteWithCache() {

}

type Client struct {
	Pool  *pgxpool.Pool
	Cache *redis.Client
}

// keep in mind you are in responsible for closing the pool after use. "defer pool.Close()"
func New() *Client {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:password@localhost:5432/madcaps")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", pool.Stat())

	redis := redisClient()

	return &Client{
		Pool:  pool,
		Cache: redis,
	}
}

func (c *Client) Insert(model string) error {
	ctx := context.Background()
	conn, err := c.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Query(context.Background(), model)
	if err != nil {
		return err
	}

	if c.Cache == nil {
		//skip caching
	}

	return nil
}
