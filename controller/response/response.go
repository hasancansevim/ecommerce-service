package response

import "go-ecommerce-service/domain"

type ErrorResponse struct {
	ErrorDescription string `json:"error_description"`
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

type UserResponse struct {
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type CartResponse struct {
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type CartItemResponse struct {
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToResponseUserData(user domain.User) UserResponse {
	return UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}
}

func ToResponseCartData(cart domain.Cart) CartResponse {
	return CartResponse{
		UserId:    cart.UserId,
		CreatedAt: cart.CreatedAt.String(),
	}
}

func ToResponseCartItemData(cartItem domain.CartItem) CartItemResponse {
	return CartItemResponse{
		CartId:    cartItem.CartId,
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	}
}

func ToResponseList(products []domain.Product) []ProductResponse {
	var productResponseList = []ProductResponse{}

	for _, product := range products {
		productResponseList = append(productResponseList, ToResponse(product))
	}

	return productResponseList
}

func ToResponseListCarts(carts []domain.Cart) []CartResponse {
	var cartResponseList = []CartResponse{}
	for _, cart := range carts {
		cartResponseList = append(cartResponseList, ToResponseCartData(cart))
	}
	return cartResponseList
}

func ToResponseListCartItems(cartItems []domain.CartItem) []CartItemResponse {
	var cartItemResponseList = []CartItemResponse{}

	for _, cart_item := range cartItems {
		cartItemResponseList = append(cartItemResponseList, ToResponseCartItemData(cart_item))
	}

	return cartItemResponseList
}
