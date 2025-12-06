package dto

import "time"

type ProductResponse struct {
	Id              uint      `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	BasePrice       float64   `json:"base_price"`
	Discount        float64   `json:"discount"`
	ImageUrl        string    `json:"image_url"`
	MetaDescription string    `json:"meta_description"`
	StockQuantity   int       `json:"stock_quantity"`
	IsActive        bool      `json:"is_active"`
	IsFeatured      bool      `json:"is_featured"`
	CategoryId      *uint     `json:"category_id"`
	StoreId         uint      `json:"store_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	BasePrice       float64 `json:"base_price"`
	Discount        float64 `json:"discount"`
	ImageUrl        string  `json:"image_url"`
	MetaDescription string  `json:"meta_description"`
	StockQuantity   int     `json:"stock_quantity"`
	IsActive        bool    `json:"is_active"`
	IsFeatured      bool    `json:"is_featured"`
	CategoryId      *uint   `json:"category_id"`
	StoreId         uint    `json:"store_id"`
}
