package domain

import "time"

type Order struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	TotalPrice float32   `json:"total_price"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
