package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-redis/redis/v9"
	"github.com/jackc/pgx/v5/pgxpool"
)

// always look into cache first, go DB if no cache is found.
func (c *Client) Get(dest interface{}) error {
	modelType := reflect.TypeOf(dest)
	ctx := context.Background()

	switch modelType.Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(dest)
		if slice.Index(0).Kind() != reflect.Ptr {
			return errors.New("destination must be a pointer slice")
		}
		table := modelType.Elem().Elem().Name()
		sql := fmt.Sprintf("SELECT * FROM %s", table)
		//look into cache
		if c.Cache != nil {
			if _, err := c.Cache.Pipelined(ctx, func(p redis.Pipeliner) error {
				for i := 0; i < slice.Len(); i++ {
					log.Printf("the %d rounds\nmax:%d rounds", i+1, slice.Len())
					ithValue := reflect.Indirect(slice.Index(i))
					log.Printf("%v\n", ithValue)
					//pkeyVal := reflect.ValueOf(ithValue.Field(0).Interface())
					//field := fmt.Sprintf("%v", pkeyVal)

					err := c.GetCache(&ithValue)
					switch {
					//no cache for this query, 1)fetch from db. 2)cache the record.
					case err == redis.Nil:
						err := pgxscan.Select(ctx, c.Pool, &dest, sql)
						if err != nil {
							panic(err)
						}
						//save to cache
						c.SaveCache(&ithValue)
					case err != nil:
						panic(err)
					}

				}
				return nil
			}); err != nil {
				panic(err)
			}

		}

	default:
		//handle errors

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
	//DEBUG
	//fmt.Printf("%+v\n", pool.Stat())

	redis := redisClient()

	return &Client{
		Pool:  pool,
		Cache: redis,
	}
}

/*
func (c *Client) Insert(model interface{}) error {
	ctx := context.Background()
	/* in normal cases you don't need to Acquire a connection manually,
	you call the query methods on the pool directly.
	Then you don't have to worry about releasing the connection (only about closing Rows).
	conn, err := c.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()


	_, err = c.Pool.Query(context.Background(), model)
	if err != nil {
		return err
	}

	if c.Cache == nil {
		//skip caching
	}

	return nil
}
*/
