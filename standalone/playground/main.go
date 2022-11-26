package main

import (
	"fmt"
	"server/lib/db"
	"time"

	"github.com/go-redis/redis/v9"
)

func main() {

	a := Account{Id: 7888, Email: "asd@sscas.com", Name: "Mike", Age: 13, Power: -992239, B: &Bee{Size: 8}, CreatedAt: time.Now()}
	err := db.SaveCache(&a)
	if err != nil {
		fmt.Println(err)
	}

	b := Account{Id: 7999, Email: "huifwe@fas.com", Name: "Tommy", Age: 63, Power: 2342423442, B: &Bee{Size: 12}, CreatedAt: time.Now()}
	err = db.SaveCache(&b)
	if err != nil {
		fmt.Println(err)
	}

	a.Power = 8888888888888
	db.SaveCache(&a)

	mike := &Account{Id: 7888}
	err = db.GetCache(mike)
	if err != nil {
		fmt.Println(err)
	}

	if err == redis.Nil {
		println("no records!")
	}

	fmt.Println(mike)

}

type Bee struct {
	Size uint8
}

type Account struct {
	Id        uint64
	Email     string
	Power     int64
	Age       int64
	CreatedAt time.Time
	Name      string
	B         *Bee
}
