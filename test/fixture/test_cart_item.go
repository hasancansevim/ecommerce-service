package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateTestCartItem() domain.CartItem {
	return domain.CartItem{
		Id:        1,
		CartId:    1,
		ProductId: 1,
		Quantity:  2,
	}
}

func CreateTestCartItemCreate() model.CartItemCreate {
	return model.CartItemCreate{
		CartId:    1,
		ProductId: 1,
		Quantity:  2,
	}
}
