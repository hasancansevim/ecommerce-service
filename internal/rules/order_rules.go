package rules

import (
	"errors"
	"go-ecommerce-service/internal/dto"
)

type OrderRules struct {
	BaseRules[dto.CreateOrderRequest]
}

func NewOrderRules() *OrderRules {
	return &OrderRules{}
}

func (r *OrderRules) ValidateCreateOrder(req dto.CreateOrderRequest) error {
	if err := r.ValidateStructure(req); err != nil {
		return err
	}

	if req.TotalPrice < 0 {
		return errors.New("Order total must be greater than 0")
	}
	return nil
}
