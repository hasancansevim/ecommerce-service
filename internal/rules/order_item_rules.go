package rules

import "go-ecommerce-service/internal/dto"

type OrderItemRules struct {
	BaseRules[dto.CreateOrderItemRequest]
}

func NewOrderItemRules() *OrderItemRules {
	return &OrderItemRules{}
}
