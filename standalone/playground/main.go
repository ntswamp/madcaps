package main

import (
	"context"
	"fmt"
	"server/lib/db"
	"server/model"
)

func main() {

	db := db.New()
	db.Pool.Ping(context.Background())
	defer db.Pool.Close()
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

	{
		/*DB

			insert := `insert into accounts(
		   		wallet_address,
		   		username,
		   		language,
		   		api_token,
		   		sst_wei,
		   		sbg_wei,
		   		created_at
		   	) values (
		   		'UA1151', 'Banana', 'zh', 'tokentest', '343242', '82', '1961-06-16'
		   	)`
		*/
	}

	accounts := []*model.Account{{Id: 1}}
	err := db.Get(&accounts)
	if err != nil {
		panic(err)
	}
	fmt.Println(*accounts[0])

}
func mike() {
	db := db.New()
	db.Pool.Ping(context.Background())
	defer db.Pool.Close()

	mike := &model.Account{Id: 88}
	_ = db.GetCache(mike)
	fmt.Println(mike.Bot)

}
