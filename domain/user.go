package domain

import "time"

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}
