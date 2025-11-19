package model

import "time"

type ProductCreate struct {
	Name            string
	Slug            string
	Description     string
	Price           float64
	BasePrice       float64
	Discount        float64
	ImageUrl        string
	MetaDescription string
	StockQuantity   int
	IsActive        bool
	IsFeatured      bool
	CategoryId      *uint
	StoreId         uint
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

type CategoryCreate struct {
	Name        string
	Description string
	IsActive    bool
}

type StoreCreate struct {
	Name           string
	Slug           string
	Description    string
	LogoUrl        string
	ContactEmail   string
	ContactPhone   string
	ContactAddress string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
