package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/auth"
	"go-ecommerce-service/service"
	_interface "go-ecommerce-service/service/interface"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/test/fixture"
	"go-ecommerce-service/test/mock"
	"go-ecommerce-service/test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	mockUserRepo   *repository.MockUserRepository
	mockJWTManager *mock.MockJWTManager
	authService    _interface.AuthService
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.mockUserRepo = new(repository.MockUserRepository)
	suite.mockJWTManager = new(mock.MockJWTManager)
	suite.authService = service.NewAuthService(suite.mockUserRepo, suite.mockJWTManager)
}

func (suite *AuthServiceTestSuite) TestRegister_Success() {
	testUserCreate := fixture.CreateTestUserCreate()

	suite.mockUserRepo.On("GetUserByEmail", testUserCreate.Email).
		Return(domain.User{}, nil)

	suite.mockUserRepo.On("CreateUser", mock2.AnythingOfType("domain.User")).
		Return(nil)

	err := suite.authService.Register(testUserCreate)

	assert.NoError(suite.T(), err)
	suite.mockUserRepo.AssertCalled(suite.T(), "CreateUser", mock2.AnythingOfType("domain.User"))
}

func (suite *AuthServiceTestSuite) TestRegister_UserAlreadyExists() {
	existingUser := fixture.CreateTestUser()
	testUserCreate := fixture.CreateTestUserCreate()

	suite.mockUserRepo.On("GetUserByEmail", testUserCreate.Email).Return(existingUser, nil)

	err := suite.authService.Register(testUserCreate)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "User already exists")
}

func (suite *AuthServiceTestSuite) TestLogin_Success() {
	password := "test_password123"
	hash, _ := auth.HashPassword(password)

	loginRequest := model.LoginCreate{
		Email:    "test@example.com",
		Password: password,
	}
	testUser := domain.User{
		Id:           1,
		Email:        "test@example.com",
		PasswordHash: hash,
	}

	suite.mockUserRepo.On("GetUserByEmail", loginRequest.Email).Return(testUser, nil)
	suite.mockJWTManager.On("GenerateToken", testUser.Id, testUser.Email).Return("fake-jwt-token", nil)

	token, err := suite.authService.Login(loginRequest)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "fake-jwt-token", token)
}

func (suite *AuthServiceTestSuite) TestLogin_UserNotFound() {
	loginRequest := model.LoginCreate{
		Email:    "nonexistent@example.com",
		Password: "anypassword",
	}

	suite.mockUserRepo.On("GetUserByEmail", loginRequest.Email).
		Return(domain.User{}, errors.New("user not found"))

	token, err := suite.authService.Login(loginRequest)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "User not found")
	assert.Equal(suite.T(), "", token)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockJWTManager.AssertNotCalled(suite.T(), "GenerateToken")
}

func (suite *AuthServiceTestSuite) TestLogin_WrongPassword() {
	loginRequest := model.LoginCreate{
		Email:    "test@example.com",
		Password: "wrong_password",
	}

	password := "correct_password"
	hash, _ := auth.HashPassword(password)

	testUser := domain.User{
		Id:           1,
		FirstName:    "test",
		LastName:     "test",
		Email:        loginRequest.Email,
		PasswordHash: hash,
	}
	suite.mockUserRepo.On("GetUserByEmail", loginRequest.Email).Return(testUser, nil)

	token, err := suite.authService.Login(loginRequest)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "Invalid email or password")
	assert.Equal(suite.T(), "", token)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockJWTManager.AssertNotCalled(suite.T(), "GenerateToken")
}

func (suite *AuthServiceTestSuite) TestLogin_TokenGenerationError() {
	password := "correct_password"
	hash, _ := auth.HashPassword(password)
	testUser := domain.User{
		Id:           1,
		FirstName:    "test",
		LastName:     "test",
		Email:        "test@example.com",
		PasswordHash: hash,
	}
	loginRequest := model.LoginCreate{
		Email:    "test@example.com",
		Password: password,
	}
	suite.mockUserRepo.On("GetUserByEmail", testUser.Email).Return(testUser, nil)

	suite.mockJWTManager.On("GenerateToken", testUser.Id, testUser.Email).
		Return("", errors.New("token generation error"))

	token, err := suite.authService.Login(loginRequest)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "Invalid token")
	assert.Equal(suite.T(), "", token)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockJWTManager.AssertExpectations(suite.T())
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
