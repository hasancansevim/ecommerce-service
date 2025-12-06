package rules

import "go-ecommerce-service/internal/dto"

type OrderRules struct {
	BaseRules[dto.CreateOrderRequest]
}

func NewOrderRules() *OrderRules {
	return &OrderRules{}
}
