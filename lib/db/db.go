package db

import (
	"errors"
	"fmt"
	"reflect"
)

// always look into cache first, go DB if no cache is found.
func Get(model interface{}) error {

	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr {
		return errors.New("model passed in must be a pointer")
	}

	v := reflect.ValueOf(model).Elem()

	pkey := reflect.ValueOf(v.Field(0).Interface())
	table := t.Elem().Name()

	cacheKey := fmt.Sprintf("%s|%v", table, pkey)
	println(cacheKey)
	//looking up cache

	return nil
}

func CachedInsert() {

}

func CachedUpdate() {

}

func DeleteWithCache() {

}
