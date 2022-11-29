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

// accept a pointer to a slice of pointers e.g., &[]*Bob{{Id:1},{Id:2},{Id:3}}
func (c *Client) Get(dest interface{}) error {
	ctx := context.Background()
	destType := reflect.TypeOf(dest)
	destVal := reflect.ValueOf(dest)
	slice := destVal.Elem()
	//resultSlice := reflect.MakeSlice(slice.Elem().Type(), 0, 0)
	if slice.Index(0).Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer slice")
	}
	tableName := destType.Elem().Elem().Elem().Name()
	pKeyName := slice.Index(0).Type().Elem().Field(0).Name

	log.Println("Table Name:", tableName, ", Pkey Name:", pKeyName)

	//look into cache

	for i := 0; i < slice.Len(); i++ {
		log.Printf("the %d rounds\nmax:%d rounds", i+1, slice.Len())
		ithValue := reflect.Indirect(slice.Index(i))

		//get pv
		pkeyVal := reflect.ValueOf(ithValue.Field(0).Interface())
		pVal := fmt.Sprintf("%v", pkeyVal)

		//getting cache
		err := c.GetCache(slice.Index(i).Interface())
		switch {
		//no cache for this query, 1)fetch from db. 2)cache the record.
		case err == redis.Nil:
			sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tableName, pKeyName, pVal)
			log.Println(sql)
			err := pgxscan.Select(ctx, c.Pool, dest, sql)
			if err != nil {
				panic(err)
			}
			//save to cache
			if c.Cache != nil {
				c.SaveCache(slice.Index(i).Interface())
			}
		case err != nil:
			panic(err)
		}

	}

	return nil
}

// accept a pointer to a slice of pointers e.g., &[]*Bob{{Id:1},{Id:2},{Id:3}}
func (c *Client) CachedInsert(model interface{}) error {

	return nil
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
