package controller

import (
	"bytes"
	"encoding/json"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/test/mock/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductControllerTestSuite struct {
	suite.Suite
	echo               *echo.Echo
	mockProductService *service.MockProductService
	productController  *controller.ProductController
}

func (suite *ProductControllerTestSuite) SetupTest() {
	suite.echo = echo.New()
	suite.mockProductService = new(service.MockProductService)
	suite.productController = controller.NewProductController(suite.mockProductService)
}

func (suite *ProductControllerTestSuite) TestAddProduct_Success() {
	productModel := model.ProductCreate{
		Name:     "product1",
		Price:    250.0,
		Discount: 30.0,
		Store:    "store_name",
	}

	suite.mockProductService.On("AddProduct", productModel).Return(nil)

	jsonData := []byte(`{
		"name" : "product1",
		"price" : 250.0,
		"discount" : 30.0,
		"store" : "store_name"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	err := suite.productController.AddProduct(ctx)

	assert.NoError(suite.T(), err)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestAddProduct_ValidationError() {
	jsonData := []byte(`{
        "name": "",
        "price": 0.0,
        "discount": 110.0,
        "store": ""
    }`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.productController.AddProduct(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(suite.T(), "Validation failed", response["error_description"])
	suite.mockProductService.AssertNotCalled(suite.T(), "AddProduct")
}

func (suite *ProductControllerTestSuite) TestGetAllProducts_Success() {
	expectedProducts := []domain.Product{
		{Id: 1, Name: "Product 1", Price: 100.0, Discount: 10.0, Store: "Store1"},
		{Id: 2, Name: "Product 2", Price: 200.0, Discount: 20.0, Store: "Store2"},
	}

	suite.mockProductService.On("GetAllProducts").Return(expectedProducts)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.productController.GetAllProducts(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestGetAllProductsByStoreName_Success() {
	expectedProducts := []domain.Product{
		{Id: 1, Name: "Product 1", Price: 100.0, Discount: 10.0, Store: "Store1"},
	}

	suite.mockProductService.On("GetAllProductsByStoreName", "Store1").Return(expectedProducts)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products?storeName=Store1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.productController.GetAllProducts(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestGetProductById_Success() {
	expectedProduct := domain.Product{Id: 1, Name: "Product 1", Price: 100.0, Discount: 10.0, Store: "Store1"}

	suite.mockProductService.On("GetProductById", int64(1)).Return(expectedProduct, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.productController.GetProductById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestGetProductById_InvalidID() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/invalid", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("invalid")

	err := suite.productController.GetProductById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	suite.mockProductService.AssertNotCalled(suite.T(), "GetProductById")
}

func (suite *ProductControllerTestSuite) TestUpdatePrice_Success() {
	suite.mockProductService.On("UpdatePrice", int64(1), float32(150.0)).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/products/1?newPrice=150.0", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.productController.UpdatePrice(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestUpdatePrice_InvalidPrice() {
	req := httptest.NewRequest(http.MethodPut, "/api/v1/products/1?newPrice=invalid", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.productController.UpdatePrice(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	suite.mockProductService.AssertNotCalled(suite.T(), "UpdatePrice")
}

func (suite *ProductControllerTestSuite) TestDeleteProduct_Success() {
	suite.mockProductService.On("DeleteProductById", int64(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.productController.DeleteProduct(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockProductService.AssertExpectations(suite.T())
}

func (suite *ProductControllerTestSuite) TestDeleteProduct_InvalidID() {
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/invalid", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("invalid")

	err := suite.productController.DeleteProduct(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	suite.mockProductService.AssertNotCalled(suite.T(), "DeleteProductById")
}

func TestProductControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductControllerTestSuite))
}
