package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"

	"github.com/stretchr/testify/mock"
)

type MockCartItemService struct {
	mock.Mock
}

func (cartItemService *MockCartItemService) AddItemToCart(cartItem model.CartItemCreate) error {
	args := cartItemService.Called(cartItem)
	return args.Error(0)
}
func (cartItemService *MockCartItemService) GetItemsByCartId(cart_id int64) []domain.CartItem {
	args := cartItemService.Called(cart_id)
	return args.Get(0).([]domain.CartItem)
}
func (cartItemService *MockCartItemService) UpdateItemQuantity(cart_item_id int64, newQuantity int) error {
	args := cartItemService.Called(cart_item_id, newQuantity)
	return args.Error(0)
}
func (cartItemService *MockCartItemService) RemoveItemFromCart(cart_item_id int64) error {
	args := cartItemService.Called(cart_item_id)
	return args.Error(0)
}
func (cartItemService *MockCartItemService) ClearCartItems(cart_id int64) error {
	args := cartItemService.Called(cart_id)
	return args.Error(0)
}
func (cartItemService *MockCartItemService) IncreaseItemQuantity(cart_item_id int64, amount int) error {
	args := cartItemService.Called(cart_item_id, amount)
	return args.Error(0)
}
func (cartItemService *MockCartItemService) DecreaseItemQuantity(cart_item_id int64, amount int) error {
	args := cartItemService.Called(cart_item_id, amount)
	return args.Error(0)
}
