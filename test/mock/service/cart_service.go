package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"

	"github.com/stretchr/testify/mock"
)

type MockCartService struct {
	mock.Mock
}

func (cartService *MockCartService) GetCartsByUserId(userId int64) []domain.Cart {
	args := cartService.Called(userId)
	return args.Get(0).([]domain.Cart)
}

func (cartService *MockCartService) GetCartById(cartId int64) domain.Cart {
	args := cartService.Called(cartId)
	return args.Get(0).(domain.Cart)
}

func (cartService *MockCartService) CreateCart(cart model.CartCreate) error {
	args := cartService.Called(cart)
	return args.Error(0)
}

func (cartService *MockCartService) DeleteCartById(cartId int64) error {
	args := cartService.Called(cartId)
	return args.Error(0)
}

func (cartService *MockCartService) ClearUserCart(userId int64) error {
	args := cartService.Called(userId)
	return args.Error(0)
}
