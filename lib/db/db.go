package db

import "github.com/go-redis/redis/v9"

// always look into cache first, go DB if no cache is found.
func Get(model interface{}) error {

	err := getCache(model)
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

func dbClient() {

}
