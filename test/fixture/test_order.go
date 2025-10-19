package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
)

func CreateTestOrder() domain.Order {
	return domain.Order{
		Id:         1,
		UserId:     1,
		TotalPrice: 100.0,
		Status:     true,
	}
}

func CreateTestOrderCreate() model.OrderCreate {
	return model.OrderCreate{
		UserId:     1,
		TotalPrice: 100.0,
		Status:     true,
	}
}

func CreateTestGetOrdersByUserId() []domain.Order {
	return []domain.Order{
		{UserId: 1, TotalPrice: 100.0, Status: true},
		{UserId: 1, TotalPrice: 200.0, Status: true},
		{UserId: 1, TotalPrice: 300.0, Status: true},
	}
}

func CreateTestGetAllOrders() []domain.Order {
	return []domain.Order{
		{UserId: 1, TotalPrice: 100.0, Status: true},
		{UserId: 1, TotalPrice: 200.0, Status: true},
		{UserId: 1, TotalPrice: 300.0, Status: true},
		{UserId: 2, TotalPrice: 400.0, Status: true},
		{UserId: 3, TotalPrice: 500.0, Status: true},
	}
}

func CreateTestOrderUpdateStatus() domain.Order {
	return domain.Order{
		Id:         1,
		UserId:     1,
		TotalPrice: 100.0,
		Status:     false,
	}
}
