package model

import (
	"math/big"
	"time"

	"github.com/uptrace/bun"
)

type Bee struct {
	Size   uint8
	Amount *big.Int
}

type Account struct {
	Id        uint64 `bun:"id,pk,autoincrement"`
	Email     string
	Power     *big.Int
	Age       int64
	CreatedAt time.Time
	Name      string
	Item      []int64 `bun:",array"`
	Bot       *Bee

	bun.BaseModel `bun:"table:accounts"`
}
