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

type AuthControllerTestSuite struct {
	suite.Suite
	echo            *echo.Echo
	mockAuthService *service.MockAuthService
	authController  *controller.AuthController
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.echo = echo.New()
	suite.mockAuthService = new(service.MockAuthService)
	suite.authController = controller.NewAuthController(suite.mockAuthService)
}

func (suite *AuthControllerTestSuite) TestRegister_Success() {
	jsonData := []byte(`{
		"first_name": "John",
		"last_name": "Doe",
		"email": "jhondoe123@gmail.com",
		"password": "123456"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	expectedModel := model.RegisterCreate{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jhondoe123@gmail.com",
		Password:  "123456",
	}

	suite.mockAuthService.On("Register", expectedModel).Return(nil)

	err := suite.authController.Register(ctx)

	// suite.T().Logf("Status: %d, Body: %s", rec.Code, rec.Body.String())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockAuthService.AssertExpectations(suite.T())
}

func (suite *AuthControllerTestSuite) TestRegister_InvalidInput() {
	jsonData := []byte(`{
		"first_name": "",
		"last_name": "",
		"email": "invalid_email",
		"password": "123"
	}`)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.authController.Register(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 400, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	// suite.T().Logf("Response : %+v", response)

	assert.Equal(suite.T(), "Validation failed", response["error_description"])

	suite.mockAuthService.AssertNotCalled(suite.T(), "Register")
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
