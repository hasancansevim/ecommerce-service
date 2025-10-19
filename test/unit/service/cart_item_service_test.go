package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/test/fixture"
	mock2 "go-ecommerce-service/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CartItemServiceTestSuite struct {
	suite.Suite
	mockCartItemRepo *mock2.MockCartItemRepository
	cartItemService  service.ICartItemService
}

func (suite *CartItemServiceTestSuite) SetupTest() {
	suite.mockCartItemRepo = new(mock2.MockCartItemRepository)
	suite.cartItemService = service.NewCartItemService(suite.mockCartItemRepo)
}

func (suite *CartItemServiceTestSuite) TestAddItemToCart_Success() {
	cartItemCreate := fixture.CreateTestCartItemCreate()

	suite.mockCartItemRepo.On("AddItemToCart", mock.MatchedBy(func(cartItem domain.CartItem) bool {
		return cartItem.CartId == cartItemCreate.CartId &&
			cartItem.ProductId == cartItemCreate.ProductId &&
			cartItem.Quantity == cartItemCreate.Quantity
	})).Return(nil)

	err := suite.cartItemService.AddItemToCart(cartItemCreate)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestUpdateItemQuantity_Success() {
	cartItemId := int64(1)
	newQuantity := 5

	suite.mockCartItemRepo.On("UpdateItemQuantity", cartItemId, newQuantity).Return(nil)

	err := suite.cartItemService.UpdateItemQuantity(cartItemId, newQuantity)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestRemoveItemFromCart_Success() {
	cartItemId := int64(1)

	suite.mockCartItemRepo.On("RemoveItemFromCart", cartItemId).Return(nil)

	err := suite.cartItemService.RemoveItemFromCart(cartItemId)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestGetItemsByCartId_Success() {
	cartId := int64(1)
	expectedItems := []domain.CartItem{
		fixture.CreateTestCartItem(),
		{Id: 2, CartId: cartId, ProductId: 2, Quantity: 3},
	}

	suite.mockCartItemRepo.On("GetItemsByCartId", cartId).Return(expectedItems)

	items := suite.cartItemService.GetItemsByCartId(cartId)

	assert.Equal(suite.T(), expectedItems, items)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestClearCartItems_Success() {
	cartId := int64(1)

	suite.mockCartItemRepo.On("ClearCartItems", cartId).Return(nil)

	err := suite.cartItemService.ClearCartItems(cartId)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestIncreaseItemQuantity_Success() {
	cartItemId := int64(1)
	amount := 2

	suite.mockCartItemRepo.On("IncreaseItemQuantity", cartItemId, amount).Return(nil)

	err := suite.cartItemService.IncreaseItemQuantity(cartItemId, amount)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestDecreaseItemQuantity_Success() {
	cartItemId := int64(1)
	amount := 1

	suite.mockCartItemRepo.On("DecreaseItemQuantity", cartItemId, amount).Return(nil)

	err := suite.cartItemService.DecreaseItemQuantity(cartItemId, amount)

	assert.NoError(suite.T(), err)
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func (suite *CartItemServiceTestSuite) TestDecreaseItemQuantity_Error() {
	cartItemId := int64(1)
	amount := 5

	suite.mockCartItemRepo.On("DecreaseItemQuantity", cartItemId, amount).
		Return(errors.New("quantity cannot be negative"))

	err := suite.cartItemService.DecreaseItemQuantity(cartItemId, amount)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "quantity cannot be negative")
	suite.mockCartItemRepo.AssertExpectations(suite.T())
}

func TestCartItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CartItemServiceTestSuite))
}
