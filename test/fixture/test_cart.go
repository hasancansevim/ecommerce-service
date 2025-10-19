package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateGetCartsByUserIdModel() []domain.Cart {
	return []domain.Cart{
		{Id: 1, UserId: 1},
		{Id: 2, UserId: 1},
		{Id: 3, UserId: 2},
	}
}

func CreateGetCartByIdModel() domain.Cart {
	return domain.Cart{Id: 1, UserId: 1}
}

func CreateCartModel() domain.Cart {
	return domain.Cart{Id: 1, UserId: 1}
}

func CreateCartCreateModel() model.CartCreate {
	return model.CartCreate{UserId: 1}
}
