package models

import "time"

type User struct {
	UserID         int64
	Email          string
	HashedPassword string
	CreatedTime    time.Time
	Status         int
}
