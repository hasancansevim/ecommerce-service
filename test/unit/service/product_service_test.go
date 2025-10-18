package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/test/fixture"
	"go-ecommerce-service/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductServiceTestSuite struct {
	suite.Suite
	mockProductRepository *mock.MockProductRepository
	productService        service.IProductService
}

func (suite *ProductServiceTestSuite) SetupTest() {
	suite.mockProductRepository = new(mock.MockProductRepository)
	suite.productService = service.NewProductService(suite.mockProductRepository)
}

func (suite *ProductServiceTestSuite) TestAddProduct_Success() {
	productCreate := fixture.CreateTestProductCreate()
	suite.mockProductRepository.On("AddProduct", mock2.MatchedBy(func(product domain.Product) bool {
		return product.Name == productCreate.Name &&
			product.Price == productCreate.Price &&
			product.Discount == productCreate.Discount &&
			product.Store == productCreate.Store
	})).Return(nil)

	err := suite.productService.AddProduct(productCreate)

	assert.NoError(suite.T(), err)
	suite.mockProductRepository.AssertExpectations(suite.T())
}

func (suite *ProductServiceTestSuite) TestGetProductById_Success() {
	productId := int64(1)
	expectedProduct := fixture.CreateTestProduct()

	suite.mockProductRepository.On("GetProductById", productId).Return(expectedProduct, nil)

	product, err := suite.productService.GetProductById(productId)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedProduct, product)
	suite.mockProductRepository.AssertExpectations(suite.T())
}

func (suite *ProductServiceTestSuite) TestGetProductById_NotFound() {
	productId := int64(876)
	suite.mockProductRepository.On("GetProductById", productId).
		Return(domain.Product{}, errors.New("product not found"))

	product, err := suite.productService.GetProductById(productId)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.Product{}, product)
	suite.mockProductRepository.AssertExpectations(suite.T())
}

func (suite *ProductServiceTestSuite) TestGetAllProductsByStoreName_Success() {
	store := "electronics-store"

	expectedProducts := []domain.Product{
		{Id: 1, Name: "Laptop", Price: 999.99, Store: store},
		{Id: 2, Name: "Phone", Price: 499.99, Store: store},
	}

	suite.mockProductRepository.On("GetAllProductsByStoreName", store).
		Return(expectedProducts, nil)

	products := suite.productService.GetAllProductsByStoreName(store)

	assert.Equal(suite.T(), expectedProducts, products)
	assert.Len(suite.T(), products, 2)
	suite.mockProductRepository.AssertExpectations(suite.T())
}

func (suite *ProductServiceTestSuite) TestUpdateProductPrice_Success() {
	productId := int64(1)
	newPrice := float32(50.00)

	suite.mockProductRepository.On("UpdatePrice", productId, newPrice).
		Return(nil)

	err := suite.productService.UpdatePrice(productId, newPrice)

	assert.NoError(suite.T(), err)
	suite.mockProductRepository.AssertExpectations(suite.T())
}

func (suite *ProductServiceTestSuite) TestDeleteProduct_Success() {
	productId := int64(1)
	suite.mockProductRepository.On("DeleteProductById", productId).Return(nil)

	err := suite.productService.DeleteProductById(productId)
	assert.NoError(suite.T(), err)

	suite.mockProductRepository.AssertExpectations(suite.T())
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
