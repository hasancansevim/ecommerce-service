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

type AddUserRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type AddCartCreate struct {
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AddCartItemRequest struct {
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func (addProductRequest AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:     addProductRequest.Name,
		Price:    addProductRequest.Price,
		Discount: addProductRequest.Discount,
		Store:    addProductRequest.Store,
	}
}

func (addUserRequest AddUserRequest) ToModel() model.UserCreate {
	return model.UserCreate{
		FirstName: addUserRequest.FirstName,
		LastName:  addUserRequest.LastName,
		Email:     addUserRequest.Email,
		Password:  addUserRequest.Password,
		CreatedAt: addUserRequest.CreatedAt,
	}
}

func (addCartCreate AddCartCreate) ToModel() model.CartCreate {
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
