package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Bee struct {
	Size uint8
}

type Account struct {
	Id        uint64 `bun:"id,pk"`
	Email     string
	Power     int64
	Age       int64
	CreatedAt time.Time
	Name      string
	Bot       *Bee

	bun.BaseModel `bun:"table:accounts"`
}
