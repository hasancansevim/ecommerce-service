package dto

import "time"

type OrderResponse struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	TotalPrice float32   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOrderRequest struct {
	UserId     int64   `json:"user_id"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
}
