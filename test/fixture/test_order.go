package fixture

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
	"time"
)

func CreateTestOrder() domain.Order {
	return domain.Order{
		Id:         1,
		UserId:     1,
		TotalPrice: 100.0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "active",
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
		{UserId: 1, TotalPrice: 100.0, Status: "active"},
		{UserId: 1, TotalPrice: 200.0, Status: "active"},
		{UserId: 1, TotalPrice: 300.0, Status: "active"},
	}
}

func CreateTestGetAllOrders() []domain.Order {
	return []domain.Order{
		{UserId: 1, TotalPrice: 100.0, Status: "active"},
		{UserId: 1, TotalPrice: 200.0, Status: "active"},
		{UserId: 1, TotalPrice: 300.0, Status: "active"},
		{UserId: 2, TotalPrice: 400.0, Status: "active"},
		{UserId: 3, TotalPrice: 500.0, Status: "active"},
	}
}

func CreateTestOrderUpdateStatus() domain.Order {
	return domain.Order{
		Id:         1,
		UserId:     1,
		TotalPrice: 100.0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Status:     "active",
	}
}
