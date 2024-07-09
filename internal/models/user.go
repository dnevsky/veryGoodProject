package models

import "time"

type User struct {
	Id           uint64
	Login        string
	PasswordHash []byte
	CreatedAt    time.Time
}
