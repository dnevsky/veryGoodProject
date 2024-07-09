package models

import "time"

type Asset struct {
	Name      string
	Uid       uint64
	Data      []byte
	CreatedAt time.Time
}
