package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
)

type IOrderItemService interface {
	AddOrderItem(orderItemCreate dto.CreateOrderItemRequest) (dto.OrderItemResponse, error)
	GetOrderItemById(orderItemId int64) (dto.OrderItemResponse, error)
	GetOrderItemsByOrderId(orderId int64) ([]dto.OrderItemResponse, error)
	GetOrderItemsByProductId(productId int64) ([]dto.OrderItemResponse, error)
	UpdateOrderItem(orderItemId int64, orderItem dto.CreateOrderItemRequest) (dto.OrderItemResponse, error)
	UpdateOrderItemQuantity(orderItemId int64, quantity int) (dto.OrderItemResponse, error)
	DeleteOrderItemById(orderItemId int64) error
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

func (orderItemService *OrderItemService) AddOrderItem(orderItemCreate dto.CreateOrderItemRequest) (dto.OrderItemResponse, error) {
	addedOrderItem, repositoryErr := orderItemService.orderItemRepository.AddOrderItem(domain.OrderItem{
		OrderId:   orderItemCreate.OrderId,
		ProductId: orderItemCreate.ProductId,
		Quantity:  orderItemCreate.Quantity,
		Price:     orderItemCreate.Price,
	})
	if repositoryErr != nil {
		return dto.OrderItemResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderItemResponse(addedOrderItem), nil
}

func (orderItemService *OrderItemService) GetOrderItemById(orderItemId int64) (dto.OrderItemResponse, error) {
	orderItem, repositoryErr := orderItemService.orderItemRepository.GetOrderItemById(orderItemId)
	if repositoryErr != nil {
		return dto.OrderItemResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderItemResponse(orderItem), nil
}

func (orderItemService *OrderItemService) GetOrderItemsByOrderId(orderId int64) ([]dto.OrderItemResponse, error) {
	orderItems, repositoryErr := orderItemService.orderItemRepository.GetOrderItemsByOrderId(orderId)
	if repositoryErr != nil {
		return []dto.OrderItemResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderItemsResponse(orderItems), nil
}

func (orderItemService *OrderItemService) GetOrderItemsByProductId(productId int64) ([]dto.OrderItemResponse, error) {
	orderItems, repositoryErr := orderItemService.orderItemRepository.GetOrderItemsByProductId(productId)

	if repositoryErr != nil {
		return []dto.OrderItemResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderItemsResponse(orderItems), nil
}

func (orderItemService *OrderItemService) UpdateOrderItem(orderItemId int64, orderItem dto.CreateOrderItemRequest) (dto.OrderItemResponse, error) {
	updatedOrderItem, repositoryErr := orderItemService.orderItemRepository.UpdateOrderItem(orderItemId, domain.OrderItem{
		Id:        orderItemId,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
		OrderId:   orderItem.OrderId,
		ProductId: orderItem.ProductId,
	})
	if repositoryErr != nil {
		return dto.OrderItemResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderItemResponse(updatedOrderItem), nil
}

func (orderItemService *OrderItemService) UpdateOrderItemQuantity(orderItemId int64, quantity int) (dto.OrderItemResponse, error) {
	orderItem, repositoryErr := orderItemService.orderItemRepository.UpdateOrderItemQuantity(orderItemId, quantity)
	if repositoryErr != nil {
		return dto.OrderItemResponse{}, repositoryErr
	}
	return convertToOrderItemResponse(orderItem), nil
}

func (orderItemService *OrderItemService) DeleteOrderItemById(orderItemId int64) error {
	repositoryErr := orderItemService.orderItemRepository.DeleteOrderItemById(orderItemId)
	if repositoryErr != nil {
		return repositoryErr
	}
	return nil
}

func (orderItemService *OrderItemService) DeleteAllOrderItemsByOrderId(orderId int64) error {
	repositoryErr := orderItemService.orderItemRepository.DeleteAllOrderItemsByOrderId(orderId)
	if repositoryErr != nil {
		return repositoryErr
	}
	return nil
}

func convertToOrderItemResponse(orderItem domain.OrderItem) dto.OrderItemResponse {
	return dto.OrderItemResponse{
		Id:        orderItem.Id,
		OrderId:   orderItem.OrderId,
		ProductId: orderItem.ProductId,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
	}
}

func convertToOrderItemsResponse(orderItem []domain.OrderItem) []dto.OrderItemResponse {
	orderItemsDto := make([]dto.OrderItemResponse, 0, len(orderItem))

	for _, orderItem := range orderItem {
		orderItemsDto = append(orderItemsDto, convertToOrderItemResponse(orderItem))
	}
	return orderItemsDto
}
