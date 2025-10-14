package model

import "time"

type ProductCreate struct {
	Name     string
	Price    float32
	Discount float32
	Store    string
}

type UserCreate struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type RegisterCreate struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginCreate struct {
	Email    string
	Password string
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

type OrderCreate struct {
	UserId     int64
	TotalPrice float32
	Status     bool
	CreatedAt  time.Time
}
type OrderItemCreate struct {
	OrderId   int64
	ProductId int64
	Quantity  int
	Price     float32
}
