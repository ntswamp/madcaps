package main

import (
	"database/sql"
	"fmt"
	"server/lib/db"
	"server/model"
)

func main() {

	db := db.New()
	defer db.Close()

	accounts := []*model.Account{{Id: 1}, {Id: 2}, {Id: 4}}
	err := db.Get(&accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	for _, a := range accounts {
		fmt.Println(*a)
	}

	/*
		{
			//REDIS
			m := Account{Id: 88, Email: "asd@sscas.com", Name: "Mike", Age: 13, Power: -992239, Bot: &Bee{Size: 8}, CreatedAt: time.Now()}
			t := Account{Id: 99, Email: "huifwe@fas.com", Name: "Tom", Age: 63, Power: 2342423442, Bot: &Bee{Size: 12}, CreatedAt: time.Now()}
			err := db.SaveCache([]Account{m, t})
			if err != nil {
				panic(err)
			}

			mike := &Account{Id: 88}
			tom := &Account{Id: 99}
			err = db.GetCache(mike)
			if err != nil {
				fmt.Println(err)
			}
			if err == redis.Nil {
				println("no 88!")
			}

			err = db.GetCache(tom)
			if err != nil {
				fmt.Println(err)
			}
			if err == redis.Nil {
				println("no 99!")
			}

			fmt.Println(mike)
			fmt.Println(tom)

			m.Power = 8888888888888
			db.SaveCache(m)
		}
	*/

}
