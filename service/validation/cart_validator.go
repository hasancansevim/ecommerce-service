package validation

import "go-ecommerce-service/service/model"

func ValidateCartCreate(cartCreate model.CartCreate) error {
	return NewValidator().
		RequiredInt(int(cartCreate.UserId), "user_id").
		Error()
}
