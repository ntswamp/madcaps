package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v9"
)

func saveCache(model interface{}) error {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		return errors.New("passed in model must be a pointer")
	}

	rc := redisClient()

	key := t.Elem().Name()
	ctx := context.Background()
	//modelVal := reflect.ValueOf(model).Elem()
	field := getPkeyValue(model)

	val, err := json.Marshal(model)
	if err != nil {
		return nil
	}
	err = rc.HSet(ctx, key, field, val).Err()
	if err != nil {
		return err
	}

	return nil
	/*
		key := makeCacheKeyName(model)
		if _, err := rc.Pipelined(ctx, func(p redis.Pipeliner) error {
			for i := 0; i < modelVal.NumField(); i++ {
				field := reflect.ValueOf(modelVal.Field(i).Interface())
				//fmt.Println(modelVal.Type().Field(i).Type == time.Time)
				p.HSet(ctx, key, modelVal.Type().Field(i).Name, fmt.Sprintf("%v", field))
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	*/
}

func getCache(destModel interface{}) error {
	t := reflect.TypeOf(destModel)
	if t.Kind() != reflect.Ptr {
		return errors.New("passed in model must be a pointer")
	}

	rc := redisClient()

	key := t.Elem().Name()
	ctx := context.Background()
	field := getPkeyValue(destModel)

	val, err := rc.HGet(ctx, key, field).Bytes()
	if err != nil {
		return err
	}

	json.Unmarshal(val, destModel)
	return nil
}

// the first field of passed in model is supposed to be the primary key.
func getPkeyValue(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		panic("passed in model must be a pointer")
	}

	v := reflect.ValueOf(model).Elem()
	pkeyVal := reflect.ValueOf(v.Field(0).Interface())

	return fmt.Sprintf("%v", pkeyVal)
}

func redisClient() *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rc
}

/*
// return "" on failure
func getCacheKeyName(model interface{}) string {

	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		panic("passed in model must be a pointer")
	}

	v := reflect.ValueOf(model).Elem()

	pkeyVal := reflect.ValueOf(v.Field(0).Interface())
	table := t.Elem().Name()

	cacheKey := fmt.Sprintf("%s|%v", table, pkeyVal)

	return cacheKey
}
*/
