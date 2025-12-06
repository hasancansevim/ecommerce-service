package domain

import "time"

type Product struct {
	Id              uint
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
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
