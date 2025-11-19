package validation

import "go-ecommerce-service/service/model"

func ValidateOrderCreate(orderCreate model.OrderCreate) error {
	return NewValidator().
		RequiredInt(int(orderCreate.UserId), "user_id").
		Error()
}
