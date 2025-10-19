package mock

import (
	"go-ecommerce-service/domain"

	"github.com/stretchr/testify/mock"
)

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) GetCartsByUserId(userId int64) []domain.Cart {
	args := m.Called(userId)
	return args.Get(0).([]domain.Cart)
}
func (m *MockCartRepository) GetCartById(cartId int64) domain.Cart {
	args := m.Called(cartId)
	return args.Get(0).(domain.Cart)
}
func (m *MockCartRepository) CreateCart(cart domain.Cart) error {
	args := m.Called(cart)
	return args.Error(0)
}
func (m *MockCartRepository) DeleteCartById(cartId int64) error {
	args := m.Called(cartId)
	return args.Error(0)
}
func (m *MockCartRepository) ClearUserCart(userId int64) error {
	args := m.Called(userId)
	return args.Error(0)
}
