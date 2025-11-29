package domain

import "time"

type Cart struct {
	Id        int64
	UserId    int64
	CreatedAt time.Time
}
