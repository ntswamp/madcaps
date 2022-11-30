package main

import (
	"fmt"
	"math/big"
	"server/lib/db"
	"server/model"
	"time"
)

func main() {

	db := db.New(true)
	defer db.Close()

	a := &model.Account{Id: 5, Item: []int64{1, 2, 3, 4, 5}, Email: "allen@a.com", Name: "Allen", Age: 23, Power: big.NewInt(3243), Bot: &model.Bee{Size: 81, Amount: big.NewInt(5)}, CreatedAt: time.Now()}
	b := &model.Account{Id: 6, Item: []int64{1, 2, 3, 4, 5}, Email: "bob@bbb.com", Name: "Bob", Age: 66, Power: big.NewInt(123123), Bot: &model.Bee{Size: 41, Amount: big.NewInt(51)}, CreatedAt: time.Now()}
	e := &model.Account{Id: 7, Item: []int64{1, 2, 3, 4, 5}, Email: "eric@err.com", Name: "Eric", Age: 1, Power: big.NewInt(121314), Bot: &model.Bee{Size: 1, Amount: big.NewInt(1)}, CreatedAt: time.Now()}
	m := &model.Account{Id: 8, Item: []int64{1, 2, 3, 4, 5}, Email: "asd@sscas.com", Name: "Mike", Age: 13, Power: big.NewInt(-25433343), Bot: &model.Bee{Size: 8, Amount: big.NewInt(731287784327)}, CreatedAt: time.Now()}
	t := &model.Account{Id: 9, Item: []int64{1, 2, 3, 4, 5}, Email: "huifwe@fas.com", Name: "Tommy", Age: 63, Power: big.NewInt(999974637), Bot: &model.Bee{Size: 12, Amount: big.NewInt(4932954387784327)}, CreatedAt: time.Now()}

	db.BeginTx()
	defer db.CommitTx()

	err := db.Upsert(&[]*model.Account{a, b, e, m, t})
	if err != nil {
		panic(err)
	}

	err = db.Delete(&[]*model.Account{a, b})
	if err != nil {
		panic(err)
	}

	dest := []*model.Account{{Id: 8}, {Id: 9}}
	err = db.Get(&dest)
	if err != nil {
		panic(err)
	}

	fmt.Println(*&(dest[0]).Name)
	fmt.Println(*&(dest[1]).Name)

}
