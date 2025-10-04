package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
	"time"
)

type IUserService interface {
	GetAllUsers() []domain.User
	AddUser(userCreate model.UserCreate) error
}

type UserService struct {
	userRespository persistence.IUserRepository
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{
		userRespository: userRepository,
	}
}

func (userService *UserService) GetAllUsers() []domain.User {
	return userService.userRespository.GetAllUser()
}

func (userService *UserService) AddUser(userCreate model.UserCreate) error {
	validateErr := validateUserCreate(userCreate)
	if validateErr != nil {
		return validateErr
	}

	return userService.userRespository.AddUser(domain.User{
		FirstName: userCreate.FirstName,
		LastName:  userCreate.LastName,
		Email:     userCreate.Email,
		Password:  userCreate.Password,
		CreatedAt: time.Now(),
	})
}

func validateUserCreate(userCreate model.UserCreate) error {
	if userCreate.FirstName == "" {
		return errors.New("First name is required")
	}
	if userCreate.LastName == "" {
		return errors.New("Last name is required")
	}
	if userCreate.Email == "" {
		return errors.New("Email is required")
	}
	if userCreate.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}
