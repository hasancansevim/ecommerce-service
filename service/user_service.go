package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/service/validation"
	"time"
)

type IUserService interface {
	GetAllUsers() []domain.User
	AddUser(userCreate model.UserCreate) error
}

type UserService struct {
	userRepository persistence.IUserRepository
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) GetAllUsers() []domain.User {
	return userService.userRepository.GetAllUser()
}

func (userService *UserService) AddUser(userCreate model.UserCreate) error {
	if validationError := validation.ValidateUserCreate(userCreate); validationError != nil {
		return validationError
	}

	return userService.userRepository.AddUser(domain.User{
		FirstName: userCreate.FirstName,
		LastName:  userCreate.LastName,
		Email:     userCreate.Email,
		Password:  userCreate.Password,
		CreatedAt: time.Now(),
	})
}
