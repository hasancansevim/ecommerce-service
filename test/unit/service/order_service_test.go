package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/service"
	mock_infra "go-ecommerce-service/test/mock/infrastructure"
	mock_repository "go-ecommerce-service/test/mock/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrderService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockIOrderRepository(ctrl)
	rabbitMQClient := mock_infra.NewMockIRabbitMQClient(ctrl)
	orderService := service.NewOrderService(mockRepo, rabbitMQClient)

	t.Run("GetOrderById_Success", func(t *testing.T) {

		orderId := int64(100)
		expectedOrder := domain.Order{
			Id:         int64(100),
			UserId:     int64(100),
			TotalPrice: 15000,
			Status:     "Pending",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		mockRepo.EXPECT().GetOrderById(orderId).Return(expectedOrder)

		result := orderService.GetOrderById(orderId)
		assert.Equal(t, expectedOrder.Id, result.Id)
		assert.Equal(t, expectedOrder.UserId, result.UserId)
		assert.Equal(t, expectedOrder.Status, result.Status)

	})

	t.Run("CreateOrder_Success", func(t *testing.T) {
		expectedOrder := domain.Order{
			UserId:     int64(100),
			TotalPrice: 15000,
			Status:     "Pending",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		createOrderReq := dto.CreateOrderRequest{
			UserId:     int64(100),
			TotalPrice: 15000,
			Status:     "Pending",
		}

		mockRepo.EXPECT().CreateOrder(gomock.Any()).Return(expectedOrder, nil)
		rabbitMQClient.EXPECT().Publish(
			"",
			"order_created_queue",
			false,
			false,
			gomock.Any(),
		).Return(nil)

		result, err := orderService.CreateOrder(createOrderReq)

		assert.NoError(t, err)
		assert.Equal(t, expectedOrder.Status, result.Status)
		assert.Equal(t, expectedOrder.UserId, result.UserId)

		time.Sleep(5 * time.Second)
	})

	t.Run("CreateOrder_ValidationError", func(t *testing.T) {

		createOrderReq := dto.CreateOrderRequest{
			UserId:     int64(100),
			TotalPrice: -150,
			Status:     "Pending",
		}

		mockRepo.EXPECT().CreateOrder(gomock.Any()).Times(0)
		rabbitMQClient.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
		result, err := orderService.CreateOrder(createOrderReq)
		assert.Error(t, err)
		assert.Equal(t, int64(0), result.Id)
	})

	t.Run("UpdateOrderStatus_Success", func(t *testing.T) {
		orderId := int64(1)
		newStatus := "Shipped"

		expectedOrder := domain.Order{
			Id:         int64(1),
			UserId:     int64(1),
			TotalPrice: 15000,
			Status:     "Shipped",
		}

		mockRepo.EXPECT().UpdateOrderStatus(orderId, newStatus).Return(expectedOrder, nil)
		response, err := orderService.UpdateOrderStatus(orderId, newStatus)

		assert.NoError(t, err)
		assert.Equal(t, expectedOrder.Status, response.Status)
	})
}
