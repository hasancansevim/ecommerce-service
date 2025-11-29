package request

import (
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/service/model"
	"time"
)

type AddProductRequest struct {
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	BasePrice       float64 `json:"basePrice"`
	Discount        float64 `json:"discount"`
	ImageUrl        string  `json:"imageUrl"`
	MetaDescription string  `json:"metaDescription"`
	StockQuantity   int     `json:"stockQuantity"`
	IsActive        bool    `json:"isActive"`
	IsFeatured      bool    `json:"isFeatured"`
	CategoryId      *uint   `json:"categoryId"`
	StoreId         uint    `json:"storeId"`
}

type UpdateProductRequest struct {
	Name            string  `json:"name"`
	Slug            string  `json:"slug"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	BasePrice       float64 `json:"basePrice"`
	Discount        float64 `json:"discount"`
	ImageUrl        string  `json:"imageUrl"`
	MetaDescription string  `json:"metaDescription"`
	StockQuantity   int     `json:"stockQuantity"`
	IsActive        bool    `json:"isActive"`
	IsFeatured      bool    `json:"isFeatured"`
	CategoryId      *uint   `json:"categoryId"`
	StoreId         uint    `json:"storeId"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddCartRequest struct {
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AddCartItemRequest struct {
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type AddOrderRequest struct {
	UserId     int64     `json:"user_id"`
	TotalPrice float32   `json:"total_price"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"created-at"`
}

type AddOrderItemRequest struct {
	OrderId   int64   `json:"order_id"`
	ProductId int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}

type UpdateOrderItemRequest struct {
	OrderId   int64   `json:"order_id"`
	ProductId int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}

type AddCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type AddStoreRequest struct {
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	LogoUrl        string    `json:"logo_url"`
	ContactEmail   string    `json:"contact_email"`
	ContactPhone   string    `json:"contact_phone"`
	ContactAddress string    `json:"contact_address"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (addProductRequest AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:            addProductRequest.Name,
		Slug:            addProductRequest.Slug,
		Description:     addProductRequest.Description,
		Price:           addProductRequest.Price,
		BasePrice:       addProductRequest.BasePrice,
		Discount:        addProductRequest.Discount,
		ImageUrl:        addProductRequest.ImageUrl,
		MetaDescription: addProductRequest.MetaDescription,
		StockQuantity:   addProductRequest.StockQuantity,
		IsActive:        addProductRequest.IsActive,
		IsFeatured:      addProductRequest.IsFeatured,
		CategoryId:      addProductRequest.CategoryId,
		StoreId:         addProductRequest.StoreId,
	}
}

func (updateProductRequest UpdateProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:            updateProductRequest.Name,
		Slug:            updateProductRequest.Slug,
		Description:     updateProductRequest.Description,
		Price:           updateProductRequest.Price,
		BasePrice:       updateProductRequest.BasePrice,
		Discount:        updateProductRequest.Discount,
		ImageUrl:        updateProductRequest.ImageUrl,
		MetaDescription: updateProductRequest.MetaDescription,
		StockQuantity:   updateProductRequest.StockQuantity,
		IsActive:        updateProductRequest.IsActive,
		IsFeatured:      updateProductRequest.IsFeatured,
		CategoryId:      updateProductRequest.CategoryId,
		StoreId:         updateProductRequest.StoreId,
	}
}

func (addCartCreate AddCartRequest) ToModel() dto.CreateCartRequest {
	return dto.CreateCartRequest{
		UserId: addCartCreate.UserId,
	}
}

func (addCartItemRequest AddCartItemRequest) ToModel() dto.CreateCartItemRequest {
	return dto.CreateCartItemRequest{
		CartId:    addCartItemRequest.CartId,
		ProductId: addCartItemRequest.ProductId,
		Quantity:  addCartItemRequest.Quantity,
	}
}

func (addOrderRequest AddOrderRequest) ToModel() model.OrderCreate {
	return model.OrderCreate{
		UserId:     addOrderRequest.UserId,
		TotalPrice: addOrderRequest.TotalPrice,
		Status:     addOrderRequest.Status,
		CreatedAt:  addOrderRequest.CreatedAt,
	}
}

func (addOrderItemRequest AddOrderItemRequest) ToModel() model.OrderItemCreate {
	return model.OrderItemCreate{
		OrderId:   addOrderItemRequest.OrderId,
		ProductId: addOrderItemRequest.ProductId,
		Quantity:  addOrderItemRequest.Quantity,
		Price:     addOrderItemRequest.Price,
	}
}

func (addCategoryRequest AddCategoryRequest) ToModel() dto.CreateCategoryRequest {
	return dto.CreateCategoryRequest{
		Name:        addCategoryRequest.Name,
		Description: addCategoryRequest.Description,
		IsActive:    addCategoryRequest.IsActive,
	}
}

func (addStoreRequest AddStoreRequest) ToModel() dto.CreateStoreRequest {
	return dto.CreateStoreRequest{
		Name:         addStoreRequest.Name,
		Slug:         addStoreRequest.Slug,
		Description:  addStoreRequest.Description,
		LogoUrl:      addStoreRequest.LogoUrl,
		ContactEmail: addStoreRequest.ContactEmail,
		ContactPhone: addStoreRequest.ContactPhone,
		IsActive:     addStoreRequest.IsActive,
	}
}
