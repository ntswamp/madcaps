package main

import (
	"fmt"
	"math/big"
	"server/lib/db"
	"server/model"
	"time"
)

func main() {

	db := db.New()
	defer db.Close()

	m := &model.Account{Id: 88, Email: "asd@sscas.com", Name: "Mike", Age: 13, Power: big.NewInt(-25433343), Bot: &model.Bee{Size: 8, Amount: big.NewInt(731287784327)}, CreatedAt: time.Now()}
	t := &model.Account{Id: 99, Email: "huifwe@fas.com", Name: "Tom", Age: 63, Power: big.NewInt(999974637), Bot: &model.Bee{Size: 12, Amount: big.NewInt(4932954387784327)}, CreatedAt: time.Now()}
	err := db.Upsert(&[]*model.Account{m, t})
	if err != nil {
		panic(err)
	}

	dest := []*model.Account{{Id: 88}, {Id: 99}}
	err = db.Get(&dest)
	if err != nil {
		panic(err)
	}

	fmt.Println(*&(dest[0]).Bot.Size)
	fmt.Println(*(dest[1]).Bot.Amount)

}
