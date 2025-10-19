package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateTestOrderItem() domain.OrderItem {
	return domain.OrderItem{
		Id:        1,
		OrderId:   1,
		ProductId: 1,
		Quantity:  2,
		Price:     150.0,
	}
}

func CreateTestOrderItemCreate() model.OrderItemCreate {
	return model.OrderItemCreate{
		OrderId:   1,
		ProductId: 1,
		Quantity:  2,
		Price:     150.0,
	}
}
