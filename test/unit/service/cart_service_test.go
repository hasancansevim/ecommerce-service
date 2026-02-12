package service

/*
type CartServiceTestSuite struct {
	suite.Suite
	mockCartRepo *repository.MockCartRepository
	cartService  service.ICartService
}

func (suite *CartServiceTestSuite) SetupTest() {
	suite.mockCartRepo = new(repository.MockCartRepository)
	suite.cartService = service.NewCartService(suite.mockCartRepo)
}

func (suite *CartServiceTestSuite) TestGetCartsByUserId_Success() {
	expectedCarts := fixture.CreateGetCartsByUserIdModel()
	suite.mockCartRepo.On("GetCartsByUserId", int64(1)).Return(expectedCarts)
	carts := suite.cartService.GetCartsByUserId(int64(1))

	assert.Equal(suite.T(), expectedCarts, carts)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestGetCartsByUserId_UserNotFound() {
	userId := int64(0)
	suite.mockCartRepo.On("GetCartsByUserId", userId).Return([]domain.Cart{})
	carts := suite.cartService.GetCartsByUserId(userId)

	assert.Equal(suite.T(), []domain.Cart{}, carts)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestGetCartById_Success() {
	expectedCarts := fixture.CreateGetCartByIdModel()
	suite.mockCartRepo.On("GetCartById", int64(1)).Return(expectedCarts)
	cart := suite.cartService.GetCartById(int64(1))

	assert.Equal(suite.T(), expectedCarts, cart)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestGetCartById_NotFound() {
	id := int64(0)
	suite.mockCartRepo.On("GetCartById", id).Return(domain.Cart{})
	cart := suite.cartService.GetCartById(id)
	assert.Equal(suite.T(), domain.Cart{}, cart)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestCreateCart_Success() {
	cartModel := fixture.CreateCartModel()
	suite.mockCartRepo.On("CreateCart", mock2.MatchedBy(func(cart domain.Cart) bool {
		return cart.UserId == cartModel.UserId
	})).Return(nil)

	err := suite.cartService.CreateCart(model.CartCreate{
		UserId: cartModel.UserId,
	})
	assert.NoError(suite.T(), err)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestDeleteCartById_Success() {
	id := int64(1)
	suite.mockCartRepo.On("DeleteCartById", id).Return(nil)
	err := suite.cartService.DeleteCartById(id)
	assert.NoError(suite.T(), err)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestDeleteCartById_NotFound() {
	id := int64(0)
	suite.mockCartRepo.On("DeleteCartById", id).
		Return(errors.New("cart not found"))

	err := suite.cartService.DeleteCartById(id)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "cart not found")
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestClearUserCart_Success() {
	userId := int64(1)
	suite.mockCartRepo.On("ClearUserCart", userId).Return(nil)
	err := suite.cartService.ClearUserCart(userId)

	assert.NoError(suite.T(), err)
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func (suite *CartServiceTestSuite) TestClearUserCart_UserNotFound() {
	userId := int64(0)
	suite.mockCartRepo.On("ClearUserCart", userId).Return(errors.New("user not found"))
	err := suite.cartService.ClearUserCart(userId)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "user not found")
	suite.mockCartRepo.AssertExpectations(suite.T())
}

func TestCartServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CartServiceTestSuite))
}
*/
