package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v9"
)

// pass in pointer or slice of pointers
func (c *Client) SaveCache(model interface{}) error {
	modelType := reflect.TypeOf(model)
	ctx := context.Background()

	switch modelType.Kind() {
	//cache single object by passing in a pointer
	case reflect.Ptr:
		key := modelType.Elem().Name()
		field := getPkeyValue(model)

		val, err := json.Marshal(model)
		if err != nil {
			return nil
		}
		err = c.Cache.HSet(ctx, key, field, val).Err()
		if err != nil {
			return err
		}
		return nil

	//cache multiple objects by passing in a pointer slice e.g., []*Bob{&a,&b,&c}
	case reflect.Slice:
		slice := reflect.ValueOf(model)
		if slice.Index(0).Kind() != reflect.Ptr {
			return errors.New("multiple models must be passed in as a pointer slice")
		}

		key := modelType.Elem().Elem().Name()
		if _, err := c.Cache.Pipelined(ctx, func(p redis.Pipeliner) error {
			for i := 0; i < slice.Len(); i++ {
				modelValue := reflect.Indirect(slice.Index(i))
				pkeyVal := reflect.ValueOf(modelValue.Field(0).Interface())
				field := fmt.Sprintf("%v", pkeyVal)
				//serialize
				val, err := json.Marshal(modelValue.Interface())
				if err != nil {
					return err
				}

				err = p.HSet(ctx, key, field, val).Err()
				if err != nil {
					return err
				}

			}
			return nil
		}); err != nil {
			panic(err)
		}
		return nil

	default:
		return errors.New("passed in model(s) must be a pointer or a slice of pointers")
	}

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

func (c *Client) GetCache(destModel interface{}) error {

	ctx := context.Background()
	destType := reflect.TypeOf(destModel)

	switch destType.Kind() {
	//get single cache
	case reflect.Ptr:

		key := destType.Elem().Name()
		field := getPkeyValue(destModel)

		val, err := c.Cache.HGet(ctx, key, field).Bytes()
		if err != nil {
			return err
		}

		json.Unmarshal(val, destModel)
		return nil

	default:
		return errors.New("passed in model must be a pointer")
	}
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
