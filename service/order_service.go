package service

import (
	"encoding/json"
	"fmt"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/infrastructure/rabbitmq"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/internal/rules"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IOrderService interface {
	CreateOrder(order dto.CreateOrderRequest) (dto.OrderResponse, error)
	GetOrderById(orderId int64) dto.OrderResponse
	GetOrdersByUserId(userId int64) ([]dto.OrderResponse, error)
	GetAllOrders() ([]dto.OrderResponse, error)
	UpdateOrderStatus(orderId int64, status string) (dto.OrderResponse, error)
	DeleteOrderById(orderId int64) error
	UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) (dto.OrderResponse, error)
	GetOrdersByStatus(status string) ([]dto.OrderResponse, error)
}

type OrderService struct {
	orderRepository persistence.IOrderRepository
	validator       *rules.OrderRules
	rabbitMQClient  rabbitmq.IRabbitMQClient
}

func NewOrderService(orderRepository persistence.IOrderRepository, rabbit rabbitmq.IRabbitMQClient) IOrderService {
	return &OrderService{
		orderRepository: orderRepository,
		validator:       rules.NewOrderRules(),
		rabbitMQClient:  rabbit,
	}
}

func (orderService *OrderService) CreateOrder(order dto.CreateOrderRequest) (dto.OrderResponse, error) {
	if validationErr := orderService.validator.ValidateStructure(order); validationErr != nil {
		return dto.OrderResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	createdOrder, repositoryErr := orderService.orderRepository.CreateOrder(domain.Order{
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	})
	if repositoryErr != nil {
		return dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}

	orderResponse := convertToOrderResponse(createdOrder)
	// RabbitMQ event
	go func() {
		payload := map[string]interface{}{
			"order_id": orderResponse.Id,
			"user_id":  orderResponse.UserId,
			"message":  "Order received. Email will be sent",
			"total":    orderResponse.TotalPrice,
		}
		body, _ := json.Marshal(payload)
		err := orderService.rabbitMQClient.Publish(
			"",
			"order_created_queue",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
		if err != nil {
			fmt.Println("RabbitMQ error: ", err)
		} else {
			fmt.Println("Event published: Order queued!")
		}
	}()
	return orderResponse, nil
}

func (orderService *OrderService) GetOrderById(orderId int64) dto.OrderResponse {
	order := orderService.orderRepository.GetOrderById(orderId)
	return convertToOrderResponse(order)
}

func (orderService *OrderService) GetOrdersByUserId(userId int64) ([]dto.OrderResponse, error) {
	ordersById, repositoryErr := orderService.orderRepository.GetOrdersByUserId(userId)
	if repositoryErr != nil {
		return []dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrdersResponse(ordersById), nil
}

func (orderService *OrderService) GetAllOrders() ([]dto.OrderResponse, error) {
	orders, repositoryErr := orderService.orderRepository.GetAllOrders()
	if repositoryErr != nil {
		return []dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrdersResponse(orders), nil
}

func (orderService *OrderService) UpdateOrderStatus(orderId int64, status string) (dto.OrderResponse, error) {
	updatedOrder, repositoryErr := orderService.orderRepository.UpdateOrderStatus(orderId, status)
	if repositoryErr != nil {
		return dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderResponse(updatedOrder), nil
}

func (orderService *OrderService) DeleteOrderById(orderId int64) error {
	repositoryErr := orderService.orderRepository.DeleteOrderById(orderId)
	if repositoryErr != nil {
		return _errors.NewBadRequest(repositoryErr.Error())
	}
	return nil
}

func (orderService *OrderService) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) (dto.OrderResponse, error) {
	updatedOrder, repositoryErr := orderService.orderRepository.UpdateOrderTotalPrice(orderId, newTotalPrice)
	if repositoryErr != nil {
		return dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrderResponse(updatedOrder), nil
}

func (orderService *OrderService) GetOrdersByStatus(status string) ([]dto.OrderResponse, error) {
	ordersByStatus, repositoryErr := orderService.orderRepository.GetOrdersByStatus(status)
	if repositoryErr != nil {
		return []dto.OrderResponse{}, _errors.NewBadRequest(repositoryErr.Error())
	}
	return convertToOrdersResponse(ordersByStatus), nil
}

func convertToOrderResponse(order domain.Order) dto.OrderResponse {
	return dto.OrderResponse{
		Id:         order.Id,
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
	}
}

func convertToOrdersResponse(orders []domain.Order) []dto.OrderResponse {
	ordersDto := make([]dto.OrderResponse, 0, len(orders))
	for _, order := range orders {
		ordersDto = append(ordersDto, convertToOrderResponse(order))
	}
	return ordersDto
}
