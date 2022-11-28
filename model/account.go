package model

import "time"

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
	Bot       *Bee
}
