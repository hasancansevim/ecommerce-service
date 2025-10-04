package domain

import "time"

type Order struct {
	Id         int64
	UserId     int64
	TotalPrice float32
	Status     bool
	CreatedAt  time.Time
}
