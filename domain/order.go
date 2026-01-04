package domain

import "time"

type Order struct {
	Id         int64
	UserId     int64
	TotalPrice float32
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
