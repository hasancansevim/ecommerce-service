package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type IOrderItemService interface {
	AddOrderItem(orderItemCreate model.OrderItemCreate) error
	GetOrderItemById(orderItem_id int64) (domain.OrderItem, error)
	GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error)
	GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error)
	UpdateOrderItem(orderItem_id int64, orderItem domain.OrderItem) error
	UpdateOrderItemQuantity(orderItem_id int64, quantity int) error
	DeleteOrderItemById(orderItem_id int64) error
	DeleteAllOrderItemsByOrderId(orderId int64) error
}

type OrderItemService struct {
	orderItemRepository persistence.IOrderItemRepository
}

func NewOrderItemService(orderItemRepository persistence.IOrderItemRepository) IOrderItemService {
	return &OrderItemService{
		orderItemRepository: orderItemRepository,
	}
}

func (orderItemService *OrderItemService) AddOrderItem(orderItemCreate model.OrderItemCreate) error {
	addOrderItemErr := orderItemService.orderItemRepository.AddOrderItem(domain.OrderItem{
		OrderId:   orderItemCreate.OrderId,
		ProductId: orderItemCreate.ProductId,
		Quantity:  orderItemCreate.Quantity,
		Price:     orderItemCreate.Price,
	})
	if addOrderItemErr != nil {
		return addOrderItemErr
	}
	return nil
}

func (orderItemService *OrderItemService) GetOrderItemById(orderItem_id int64) (domain.OrderItem, error) {
	getOrderItemById, getOrderItemByIdErr := orderItemService.orderItemRepository.GetOrderItemById(orderItem_id)
	if getOrderItemByIdErr != nil {
		return domain.OrderItem{}, getOrderItemByIdErr
	}
	return getOrderItemById, nil
}

func (orderItemService *OrderItemService) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	getOrderItemsByOrderId, getOrderItemsByOrderIdErr := orderItemService.orderItemRepository.GetOrderItemsByOrderId(orderId)
	if getOrderItemsByOrderIdErr != nil {
		return []domain.OrderItem{}, getOrderItemsByOrderIdErr
	}
	return getOrderItemsByOrderId, nil
}

func (orderItemService *OrderItemService) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	getOrderItemsByProductId, getOrderItemsByProductIdErr := orderItemService.orderItemRepository.GetOrderItemsByProductId(productId)
	if getOrderItemsByProductIdErr != nil {
		return []domain.OrderItem{}, getOrderItemsByProductIdErr
	}
	return getOrderItemsByProductId, nil
}

func (orderItemService *OrderItemService) UpdateOrderItem(orderItem_id int64, orderItem domain.OrderItem) error {
	updateOrderItemErr := orderItemService.UpdateOrderItem(orderItem_id, orderItem)
	if updateOrderItemErr != nil {
		return updateOrderItemErr
	}
	return nil
}

func (orderItemService *OrderItemService) UpdateOrderItemQuantity(orderItem_id int64, quantity int) error {
	updateOrderItemQuantityErr := orderItemService.orderItemRepository.UpdateOrderItemQuantity(orderItem_id, quantity)
	if updateOrderItemQuantityErr != nil {
		return updateOrderItemQuantityErr
	}
	return nil
}

func (orderItemService *OrderItemService) DeleteOrderItemById(orderItem_id int64) error {
	deleteOrderItemByIdErr := orderItemService.orderItemRepository.DeleteOrderItemById(orderItem_id)
	if deleteOrderItemByIdErr != nil {
		return deleteOrderItemByIdErr
	}
	return nil
}

func (orderItemService *OrderItemService) DeleteAllOrderItemsByOrderId(orderId int64) error {
	deleteAllOrderItemsByOrderIdErr := orderItemService.orderItemRepository.DeleteAllOrderItemsByOrderId(orderId)
	if deleteAllOrderItemsByOrderIdErr != nil {
		return deleteAllOrderItemsByOrderIdErr
	}
	return nil
}
