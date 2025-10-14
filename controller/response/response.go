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

type OrderResponse struct {
	UserId     int64   `json:"user_id"`
	TotalPrice float32 `json:"total_price"`
	Status     bool    `json:"status"`
	CreatedAt  string  `json:"created_at"`
}

type OrderItemResponse struct {
	OrderId   int64   `json:"order_id"`
	ProductId int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
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

func ToResponseOrderData(order domain.Order) OrderResponse {
	return OrderResponse{
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt.String(),
	}
}

func ToResponseOrderItemData(orderItem domain.OrderItem) OrderItemResponse {
	return OrderItemResponse{
		OrderId:   orderItem.OrderId,
		ProductId: orderItem.ProductId,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
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

func ToResponseListOrders(orders []domain.Order) []OrderResponse {
	var orderResponseList = []OrderResponse{}

	for _, order := range orders {
		orderResponseList = append(orderResponseList, ToResponseOrderData(order))
	}
	return orderResponseList
}

func ToResponseListOrderItems(orderItems []domain.OrderItem) []OrderItemResponse {
	var orderItemResponseList = []OrderItemResponse{}
	for _, orderItem := range orderItems {
		orderItemResponseList = append(orderItemResponseList, ToResponseOrderItemData(orderItem))
	}
	return orderItemResponseList
}
