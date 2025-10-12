package validation

import "go-ecommerce-service/service/model"

func ValidateCartItemCreate(cartItemCreate model.CartItemCreate) error {
	return NewValidator().
		RequiredInt(int(cartItemCreate.CartId), "cart_id").
		RequiredInt(int(cartItemCreate.ProductId), "product_id").
		Error()
}
