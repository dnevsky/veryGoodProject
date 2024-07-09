package models

import "time"

type Session struct {
	Id        string
	Uid       uint64
	Ip        string
	CreatedAt time.Time
}
