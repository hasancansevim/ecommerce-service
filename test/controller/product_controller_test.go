package controller

import (
	"bytes"
	"encoding/json"
	"go-ecommerce-service/controller"
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

func TestProductControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductControllerTestSuite))
}
