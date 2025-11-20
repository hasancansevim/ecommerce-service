package domain

import "time"

type Cart struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
