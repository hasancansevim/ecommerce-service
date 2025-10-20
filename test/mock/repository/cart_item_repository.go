package repository

import (
	"go-ecommerce-service/domain"

	"github.com/stretchr/testify/mock"
)

type MockCartItemRepository struct {
	mock.Mock
}

func (m *MockCartItemRepository) AddItemToCart(cartItem domain.CartItem) error {
	args := m.Called(cartItem)
	return args.Error(0)
}

func (m *MockCartItemRepository) UpdateItemQuantity(cart_item_id int64, newQuantity int) error {
	args := m.Called(cart_item_id, newQuantity)
	return args.Error(0)
}

func (m *MockCartItemRepository) RemoveItemFromCart(cart_item_id int64) error {
	args := m.Called(cart_item_id)
	return args.Error(0)
}

func (m *MockCartItemRepository) GetItemsByCartId(cart_id int64) []domain.CartItem {
	args := m.Called(cart_id)
	return args.Get(0).([]domain.CartItem)
}

func (m *MockCartItemRepository) ClearCartItems(cart_id int64) error {
	args := m.Called(cart_id)
	return args.Error(0)
}

func (m *MockCartItemRepository) IncreaseItemQuantity(cart_item_id int64, amount int) error {
	args := m.Called(cart_item_id, amount)
	return args.Error(0)
}

func (m *MockCartItemRepository) DecreaseItemQuantity(cart_item_id int64, amount int) error {
	args := m.Called(cart_item_id, amount)
	return args.Error(0)
}
