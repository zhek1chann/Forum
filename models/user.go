package models

import "time"

type User struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Status         int
}
