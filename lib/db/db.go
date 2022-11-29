package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// accept a pointer to a slice of pointers e.g., &[]*Bob{{Id:1},{Id:2},{Id:3}}
func (c *Client) Get(dest interface{}) error {
	ctx := context.Background()
	destVal := reflect.ValueOf(dest)
	slice := destVal.Elem()

	if slice.Index(0).Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer slice")
	}

	/***DEBUG
	destType := reflect.TypeOf(dest)
	tableName := destType.Elem().Elem().Elem().Name()
	pKeyName := slice.Index(0).Type().Elem().Field(0).Name
	fmt.Println("Table Name:", tableName, ", Pkey Name:", pKeyName, ", Slice:", slice)
	***/

	//look into cache
	for i := 0; i < slice.Len(); i++ {
		//get primary key value
		pVal := getPkValue(slice.Index(i).Interface())
		//try getting cache
		err := c.GetCache(slice.Index(i).Interface())
		switch {
		//no cache for this query, 1)fetch from db. 2)cache the record.
		case err == redis.Nil:
			err := c.Db.NewSelect().Model(slice.Index(i).Interface()).Where("?PKs = ?", pVal).Scan(ctx)
			if err != nil {
				return err
			}

			/*
				err := c.Db.NewRaw(
					"SELECT * FROM ? WHERE ? = ?",
					bun.Ident(strings.ToLower(tableName)), pKeyName, pVal,
				).Scan(ctx, dest)
				if err != nil {
					println("asdasdasdsadasasasdasd")

				}
				println("ninin")
			*/
			/*
				sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", tableName, pKeyName, pVal)
				fmt.Println(sql)
				err := pgxscan.Select(ctx, c.Pool, dest, sql)
				if err != nil {
					return err
				}
			*/
			//make cache
			if c.Cache != nil {
				err = c.SaveCache(slice.Index(i).Interface())
				if err != nil {
					panic(err)
				}
			}
		case err != nil:
			panic(err)
		}
	}
	return nil
}

// accept a pointer to a slice of pointers e.g., &[]*Bob{{Id:1},{Id:2},{Id:3}}
// the order is upsert to DB before upserting cache
func (c *Client) Upsert(model interface{}) error {
	ctx := context.Background()
	//modelType := reflect.TypeOf(model)
	modelVal := reflect.ValueOf(model)
	//slice would be [*struct,*struct,*struct...]
	slice := modelVal.Elem()

	if slice.Index(0).Kind() != reflect.Ptr {
		return errors.New("accept only pointer slices")
	}

	for i := 0; i < slice.Len(); i++ {

		_, err := c.Db.NewInsert().
			Model(slice.Index(i).Interface()).
			On("CONFLICT (?PKs) DO UPDATE").
			Exec(ctx)
		if err != nil {
			return err
		}
		/*
			err := pgxscan.Select(ctx, c.Pool, dest, sql)
			if err != nil {
				panic(err)
			}
		*/
		//save to cache
		if c.Cache != nil {
			err = c.SaveCache(slice.Index(i).Interface())
			if err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func DeleteWithCache() {

}

type Client struct {
	Cache *redis.Client
	Db    *bun.DB
}

// keep in mind you are in responsible for closing the client afteruse: defer client.Close()
func New() *Client {
	dsn := "postgres://postgres:password@localhost:5432/madcaps?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	redis := redisClient()

	return &Client{
		Db:    db,
		Cache: redis,
	}
}

func (c *Client) Close() {
	c.Db.Close()
	c.Cache.Close()
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

// accept a pointer to a struct.
func getPkName(model interface{}) string {
	t := reflect.TypeOf(model).Elem()
	if t.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := strings.Split(field.Tag.Get("bun"), ",") // use split to ignore tag "options" like omitempty, etc.

		for _, tag := range tags {
			if tag == "pk" {
				return field.Name
			}
		}
	}
	return ""
}

// accept a pointer to a struct.
func getPkValue(model interface{}) string {
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()
	if t.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := strings.Split(field.Tag.Get("bun"), ",") // use split to ignore tag "options" like omitempty, etc.

		for _, tag := range tags {
			if tag == "pk" {
				return fmt.Sprintf("%v", reflect.ValueOf(v.Field(i).Interface()))
			}
		}
	}
	return ""
}
