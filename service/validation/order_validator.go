package validation

import "go-ecommerce-service/service/model"

func ValidateOrderCreate(orderCreate model.OrderCreate) error {
	return NewValidator().
		RequiredInt(int(orderCreate.UserId), "user_id").
		GreaterThanFloat32(orderCreate.TotalPrice, 0, "total_price").
		Error()
}
