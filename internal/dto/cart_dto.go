package dto

import "time"

type CartResponse struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCartRequest struct {
	UserId int64 `json:"user_id"`
}
