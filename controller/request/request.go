package request

import (
	"go-ecommerce-service/service/model"
	"time"
)

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
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

func (addProductRequest AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}

func (addCartCreate AddCartRequest) ToModel() model.CartCreate {
	return model.CartCreate{
		UserId:    addCartCreate.UserId,
		CreatedAt: addCartCreate.CreatedAt,
	}
}

func (addCartItemRequest AddCartItemRequest) ToModel() model.CartItemCreate {
	return model.CartItemCreate{
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
