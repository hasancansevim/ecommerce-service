package model

import "time"

type ProductCreate struct {
	Name     string
	Price    float32
	Discount float32
	Store    string
}

type UserCreate struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type CartCreate struct {
	UserId    int64
	CreatedAt time.Time
}

type CartItemCreate struct {
	CartId    int64
	ProductId int64
	Quantity  int
}
