package validation

import "go-ecommerce-service/service/model"

func ValidateOrderItemCreate(orderItemCreate model.OrderItemCreate) error {
	return NewValidator().
		RequiredInt(int(orderItemCreate.OrderId), "order_id").
		RequiredInt(int(orderItemCreate.ProductId), "product_id").
		RequiredInt(orderItemCreate.Quantity, "quantity").
		Error()
}
