package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type IOrderService interface {
	CreateOrder(order model.OrderCreate) error
	GetOrderById(orderId int64) domain.Order
	GetOrdersByUserId(userId int64) ([]domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	UpdateOrderStatus(orderId int64, status bool) error
	DeleteOrderById(orderId int64) error
	UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error
	GetOrdersByStatus(status string) ([]domain.Order, error)
	GetOrderCountByUserId(userId int64) (int, error)
}

type OrderService struct {
	orderRepository persistence.IOrderRepository
}

func NewOrderService(orderRepository persistence.IOrderRepository) IOrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (orderService *OrderService) CreateOrder(orderCreate model.OrderCreate) error {
	createOrderErr := orderService.orderRepository.CreateOrder(domain.Order{
		UserId:     orderCreate.UserId,
		TotalPrice: orderCreate.TotalPrice,
		Status:     orderCreate.Status,
		CreatedAt:  orderCreate.CreatedAt,
	})
	if createOrderErr != nil {
		return createOrderErr
	} else {
		return nil
	}
}

func (orderService *OrderService) GetOrderById(orderId int64) domain.Order {
	return orderService.orderRepository.GetOrderById(orderId)

}

func (orderService *OrderService) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	getOrdersByUserId, getOrdersByUserIdErr := orderService.orderRepository.GetOrdersByUserId(userId)
	if getOrdersByUserIdErr != nil {
		return []domain.Order{}, getOrdersByUserIdErr
	}
	return getOrdersByUserId, nil
}

func (orderService *OrderService) GetAllOrders() ([]domain.Order, error) {
	orders, getAllOrdersErr := orderService.orderRepository.GetAllOrders()
	if getAllOrdersErr != nil {
		return []domain.Order{}, getAllOrdersErr
	}
	return orders, nil
}

func (orderService *OrderService) UpdateOrderStatus(orderId int64, status bool) error {
	updateOrderStatusErr := orderService.orderRepository.UpdateOrderStatus(orderId, status)
	if updateOrderStatusErr != nil {
		return updateOrderStatusErr
	}
	return nil
}

func (orderService *OrderService) DeleteOrderById(orderId int64) error {
	deleteOrderByIdErr := orderService.orderRepository.DeleteOrderById(orderId)
	if deleteOrderByIdErr != nil {
		return deleteOrderByIdErr
	}
	return nil
}

func (orderService *OrderService) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error {
	updateOrderTotalPriceErr := orderService.orderRepository.UpdateOrderTotalPrice(orderId, newTotalPrice)
	if updateOrderTotalPriceErr != nil {
		return updateOrderTotalPriceErr
	}
	return nil
}

func (orderService *OrderService) GetOrdersByStatus(status string) ([]domain.Order, error) {
	ordersByStatus, getOrdersByStatusErr := orderService.orderRepository.GetOrdersByStatus(status)
	if getOrdersByStatusErr != nil {
		return []domain.Order{}, getOrdersByStatusErr
	}
	return ordersByStatus, nil
}

func (orderService *OrderService) GetOrderCountByUserId(userId int64) (int, error) {
	orderCountByUserId, orderCountByUserIdErr := orderService.orderRepository.GetOrderCountByUserId(userId)
	if orderCountByUserIdErr != nil {
		return 0, orderCountByUserIdErr
	}
	return orderCountByUserId, nil
}
